package qqmeeting

import (
	"bytes"
	"crypto"
	"crypto/hmac"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Describe A Request
type MeetingRequestDescriptor struct {
	Url    string
	Method string
	Tag    string
}

// MeetingRequest interface
type MeetingRequest interface {
	getDescriptor() *MeetingRequestDescriptor
	fillPlaceholder(args ...interface{}) string
}

type MeetingResponse interface {
}

type MeetingErrorResponse struct {
	StatusCode int           `json:"-"`
	ErrorInfo  *MeetingError `json:"error_info"`
}

type MeetingError struct {
	Code    int    `json:"error_code"`
	Message string `json:"message"`
}

func (e MeetingError) Error() string {
	return fmt.Sprintf("server error: %s (%d)", e.Message, e.Code)
}

type Pager struct {
	TotalCount  int `json:"total_count"`  // 总数
	CurrentSize int `json:"current_size"` // 当前页实际大小
	CurrentPage int `json:"current_page"` // 当前页数
	PageSize    int `json:"page_size"`    // 分页大小
}

type Meeting struct {
	SecretKey  string
	SecretID   string
	AppID      string
	SdkID      string
	Version    string // 软件版本，用于调试
	Registered int    // 企业用户管理，最好开，否则主持人的功能用不了
}

// RequestBody Descriptor
type Request struct {
	Method        string
	URL           *url.URL
	Secret        string
	Body          string
	Key           string `json:"X-TC-Key"`
	Timestamp     int64  `json:"X-TC-Timestamp"`
	Nonce         int    `json:"X-TC-Nonce"`
	Signature     string `json:"X-TC-Signature"`
	AppID         string `json:"AppId"`
	SdkID         string `json:"SdkId"`
	Version       string `json:"X-TC-Version"`
	Registered    int    `json:"X-TC-Registered"`
	ContentType   string `json:"Content-Type"`
	ContentLength string `json:"Content-Length"`
}

//var proxyUrl, _ = url.Parse("http://127.0.0.1:18080")
var client = &http.Client{
	Transport: &http.Transport{
		//Proxy: http.ProxyURL(proxyUrl),
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		ForceAttemptHTTP2:     false,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	},
}

func GetHttpClient() *http.Client {
	return client
}

func newMeetingRequest(method, path, body string, meeting Meeting) *Request {

	req := new(Request)

	req.ContentType = "application/json"
	req.Method = method
	req.URL, _ = url.Parse(path)
	req.Secret = meeting.SecretKey
	req.Body = body

	req.Timestamp = time.Now().Unix()
	req.Key = meeting.SecretID
	req.Nonce = rand.Intn(10000) + 10000
	req.Version = meeting.Version
	req.AppID = meeting.AppID
	req.Registered = meeting.Registered
	req.SdkID = meeting.SdkID

	return req

}

func NewRequest(method, url, body string, meeting Meeting) (*http.Request, error) {
	method = strings.ToUpper(method)
	req, err := http.NewRequest(method, url, strings.NewReader(body))
	if err != nil {
		return nil, err
	}
	mReq := newMeetingRequest(method, url, body, meeting)

	fillSignature(mReq)
	fillHeader(mReq, &req.Header)

	return req, nil
}

func serializeHeader(request *Request) string {

	var buf bytes.Buffer
	buf.WriteString("X-TC-Key=" + request.Key +
		"&" + "X-TC-Nonce=" + strconv.Itoa(request.Nonce) +
		"&" + "X-TC-Timestamp=" + strconv.Itoa(int(request.Timestamp)))
	return buf.String()
}

func fillHeader(req *Request, header *http.Header) () {

	callback := func(n, v string) () {
		//fmt.Printf("%s:%s\n", n, v)
		(*header)[n] = []string{v}
	}
	fillFields(req, callback)
}

func fillSignature(req *Request) {
	ques := ""
	if len(req.URL.RawQuery) > 0 {
		ques = "?"
	}
	stringToSign := req.Method + "\n" + serializeHeader(req) + "\n" + req.URL.Path + ques + req.URL.RawQuery + "\n" + req.Body
	hm := hmac.New(crypto.SHA256.New, []byte(req.Secret))
	hm.Write([]byte(stringToSign))
	result := hm.Sum(nil)
	// debug stub
	//log.Println(stringToSign)
	req.Signature = base64.StdEncoding.EncodeToString([]byte(hex.EncodeToString(result)))
}

func fillFields(req *Request, callback func(name, value string) ()) {
	ref := reflect.ValueOf(*req)
	typ := reflect.TypeOf(*req)

	numRef := ref.NumField()
	for i := 0; i < numRef; i++ {
		field := ref.Field(i)
		fieldType := typ.Field(i)
		if fieldType.Tag.Get("json") != "" && !field.IsZero() {
			switch field.Kind() {
			case reflect.String:
				callback(fieldType.Tag.Get("json"), field.String())
			case reflect.Int64:
				callback(fieldType.Tag.Get("json"), strconv.Itoa(int(field.Int())))
			case reflect.Int:
				callback(fieldType.Tag.Get("json"), strconv.Itoa(int(field.Int())))
			}
		}
	}
}

func (meeting Meeting) Do(req MeetingRequest) (MeetingResponse, error) {

	descriptor := req.getDescriptor()
	_, method := descriptor.Url, descriptor.Method

	urlAppendix := bytes.NewBufferString("")
	urlAppendixFlag := false
	reqBody := ""

	typ := reflect.TypeOf(req)
	val := reflect.ValueOf(req)

	params := make([]interface{}, 0, 3)
	queries := NewQueryValues()

	// inject params and queries
	for i := 0; i < typ.NumField(); i++ {
		fTyp := typ.Field(i)
		fVal := val.Field(i)
		if !fVal.IsZero() && fTyp.Tag.Get("query") != "" {
			urlAppendixFlag = true
			switch fVal.Kind() {
			case reflect.String:
				queries.Add(fTyp.Tag.Get("query"), fVal.String())
			case reflect.Int, reflect.Int64, reflect.Int32:
				queries.Add(fTyp.Tag.Get("query"), strconv.Itoa(int(fVal.Int())))
			case reflect.Bool:
				queries.Add(fTyp.Tag.Get("query"), strconv.FormatBool(fVal.Bool()))
			default:
				// not formatted. print type
				queries.Add(fTyp.Tag.Get("query"), fVal.String())
			}
		} // if
		// 按照struct的顺序来填充数据
		if !fVal.IsZero() && fTyp.Tag.Get("param") != "" {
			switch fVal.Kind() {
			case reflect.String:
				params = append(params, fVal.String())
			case reflect.Int, reflect.Int64, reflect.Int32:
				params = append(params, strconv.Itoa(int(fVal.Int())))
			case reflect.Bool:
				params = append(params, strconv.FormatBool(fVal.Bool()))
			}
		} // if
	} // for

	if urlAppendixFlag {
		urlAppendix.WriteString("?" + queries.Encode())
	}

	switch method {
	case "POST", "PUT":
		body, err := json.Marshal(req)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		reqBody = string(body)
	}

	hReq, err := NewRequest(method, ApiHost+req.fillPlaceholder(params...)+urlAppendix.String(), reqBody, meeting)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	resp, err := client.Do(hReq)
	if err != nil {
		log.Println(err)
		return nil, err
	} else {
		res, _ := ioutil.ReadAll(resp.Body)
		if resp.StatusCode == 200 {
			response, err := meeting.handleResponse(res, descriptor)
			if err != nil {
				return nil, err
			} else {
				return response, nil
			}
		} else {
			var errorResponse MeetingErrorResponse
			err := json.Unmarshal(res, &errorResponse)
			if err != nil {
				return nil, err
			} else {
				return nil, *(errorResponse.ErrorInfo)
			}
		}

	}
}
