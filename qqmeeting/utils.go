package qqmeeting

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
)

// 默认URL Parameter处理方式
func defaultPlaceholderFiller(req MeetingRequest, args ...interface{}) string {
	return fmt.Sprintf(req.getDescriptor().Url, args...)
}

func PrintResponseJsonString(response MeetingResponse) string {
	result, _ := json.Marshal(response)
	log.Println(string(result))
	return string(result)
}

type KeyValuePair struct {
	Key   string
	Value string
}
type QueryValues []*KeyValuePair

func NewQueryValues() QueryValues {
	return make(QueryValues, 0, 2)
}

func (v *QueryValues) Add(key, value string) {
	*v = append(*v, &KeyValuePair{
		key, value,
	})
}

func (v *QueryValues) Encode() string {

	if len(*v) == 0 {
		return ""
	}
	var buf bytes.Buffer
	flag := false
	for _, v := range *v {
		if flag {
			buf.WriteString("&")
		} else {
			flag = true
		}
		buf.WriteString(fmt.Sprintf("%s=%s", url.QueryEscape(v.Key), url.QueryEscape(v.Value)))
	}
	return buf.String()
}
