package qqmeeting_test

import (
	"github.com/hafrans/tencent-meeting-sdk-golang/qqmeeting"
	"io/ioutil"
	"log"
	"strconv"
	"testing"
	"time"
)

var meeting = qqmeeting.Meeting{
	
	Registered: qqmeeting.EnableRegistered,
}

func TestNewRequest(t *testing.T) {

	req, err := qqmeeting.NewRequest("GET", "https://api.meeting.qq.com/v1/meetings", "", meeting)
	if err != nil {
		t.Error(err)
	} else {
		client := qqmeeting.GetHttpClient()
		resp, err := client.Do(req)
		if err != nil {
			t.Error(err)
		} else {
			content, _ := ioutil.ReadAll(resp.Body)
			t.Log(string(content))
		}
	}

}

func TestMeeting_Do_CreateMeeting(t *testing.T) {

	resp, err := meeting.Do(qqmeeting.MeetingCreateRequest{
		Settings: &qqmeeting.Settings{
			MuteEnableJoin:  true,
			MuteAll:         true,
			AllowUnmuteSelf: true,
		},
		InstanceID: qqmeeting.InstancePC,
		Type:       qqmeeting.MeetingTypeBookingMeeting,
		UserID:     "13800138000",
		Password:   "123456",
		StartTime:  strconv.Itoa(int(time.Now().Unix() + 30)),
		EndTime:    strconv.Itoa(int(time.Now().Unix() + 360)),
		Subject:    "测试会议",
	})

	if err != nil {
		t.Error(err)
	} else {

		r := resp.(qqmeeting.MeetingCreateResponse)
		t.Log(r.MeetingCreationInfo[0].Subject)
		t.Log(r.MeetingCreationInfo[0].StartTime)
		t.Log(r.MeetingCreationInfo[0].EndTime)
		t.Log("meeting code:", r.MeetingCreationInfo[0].MeetingCode)
		t.Log("meeting id:", r.MeetingCreationInfo[0].MeetingID)
		t.Log("meeting join url:", r.MeetingCreationInfo[0].JoinUrl)
		qqmeeting.PrintResponseJsonString(resp)

	}
}

func TestMeeting_Do_CreateUser(t *testing.T) {

	_, err := meeting.Do(qqmeeting.UserCreateRequest{
		UserInfo: qqmeeting.UserInfo{
			UserID:   "13800138000",
			Username: "刘大胆",
			Phone:    "13800138000",
			Email:    "admin@163.com",
		},
	})
	if err != nil {
		t.Error(err)
	} else {
		t.Log("用户创建成功")
	}
}

func TestMeeting_Do_QueryUserDetail(t *testing.T) {

	detail, err := meeting.Do(qqmeeting.UserDetailQueryRequest{
		UserID: "13800138000",
	})
	if err != nil {
		t.Error(err)
	} else {
		t.Log("用户信息获取成功")
		d, ok := detail.(qqmeeting.UserDetailQueryResponse)
		if ok {
			t.Log(d.UserID)
			t.Log(d.Email)
			t.Log(d.Status)
			t.Log(d.Username)
			t.Log(d.AreaCode)
			t.Log(d.UpdateTime)
		} else {
			t.Error("判断错误")
		}
	}
}

func TestMeeting_Do_UpdateUser(t *testing.T) {

	_, err := meeting.Do(qqmeeting.UserDetailUpdateRequest{
		Username: "刘华强",
		Email:    "huaqiang@test.com",
		UserID:   "13800138000",
	})
	if err != nil {
		t.Error(err)
	} else {
		t.Log("用户修改成功")
	}
}

func TestMeeting_Do_DeleteUser(t *testing.T) {

	_, err := meeting.Do(qqmeeting.UserDeleteRequest{
		UserID: "13800138000",
	})
	if err != nil {
		t.Error(err)
	} else {
		t.Log("用户修改成功")
	}
}

func TestMeeting_Do_ListUser(t *testing.T) {

	list, err := meeting.Do(qqmeeting.UserListRequest{
		Page:     1,
		PageSize: 20,
	})
	if err != nil {
		t.Error(err)
	} else {
		t.Log("数据获取成功")

		for _, d := range list.(qqmeeting.UserListResponse).Users {
			t.Log("-------------------------")
			t.Log(d.UserID)
			t.Log(d.Email)
			t.Log(d.Status)
			t.Log(d.Username)
			t.Log(d.AreaCode)
			t.Log(d.UpdateTime)
			t.Log("-------------------------")
		}
	}
}

func TestMeeting_Do_QueryByID(t *testing.T) {

	resp, err := meeting.Do(qqmeeting.MeetingQueryByIDRequest{
		UserID:     "13800138000",
		MeetingID:  "3618694739060173625",
		InstanceID: qqmeeting.InstancePC,
	})
	if err != nil {
		t.Error(err)
	} else {
		r := resp.(qqmeeting.MeetingQueryResponse)
		t.Log(r.MeetingInfoList[0].Subject)
		t.Log(r.MeetingInfoList[0].StartTime)
		t.Log(r.MeetingInfoList[0].EndTime)
		t.Log(r.MeetingInfoList[0].Hosts[0].UserID)
		t.Log("meeting code:", r.MeetingInfoList[0].MeetingCode)
		t.Log("meeting id:", r.MeetingInfoList[0].MeetingID)
		t.Log("meeting join url:", r.MeetingInfoList[0].JoinUrl)
		t.Log("status:", r.MeetingInfoList[0].Status)
	}
}

func TestMeeting_Do_QueryByCode(t *testing.T) {

	resp, err := meeting.Do(qqmeeting.MeetingQueryByCodeRequest{
		UserID:      "13800138000",
		MeetingCode: "290113240",
		InstanceID:  qqmeeting.InstancePC,
	})
	if err != nil {
		t.Error(err)
	} else {
		r := resp.(qqmeeting.MeetingQueryResponse)
		t.Log(r.MeetingInfoList[0].Subject)
		t.Log(r.MeetingInfoList[0].StartTime)
		t.Log(r.MeetingInfoList[0].EndTime)
		t.Log("meeting code:", r.MeetingInfoList[0].MeetingCode)
		t.Log("meeting id:", r.MeetingInfoList[0].MeetingID)
		t.Log("meeting join url:", r.MeetingInfoList[0].JoinUrl)
		t.Log("status:", r.MeetingInfoList[0].Status)
	}
}

func TestMeeting_Do_CancelMeeting(t *testing.T) {

	_, err := meeting.Do(qqmeeting.MeetingCancelRequest{
		UserID:     "13800138000",
		MeetingID:  "3618694739060173625",
		InstanceID: qqmeeting.InstancePC,
		ReasonCode: 1080,
	})
	if err != nil {
		t.Error(err)
	} else {

	}
}

func TestMeeting_Do_UpdateMeeting(t *testing.T) {

	result, err := meeting.Do(qqmeeting.MeetingUpdateRequest{
		UserID:     "13800138000",
		MeetingID:  "3710955756145476773",
		InstanceID: qqmeeting.InstancePC,
		Subject:    "我要修改会议信息",
	})
	if err != nil {
		t.Error(err)
	} else {
		log.Println(result.(qqmeeting.MeetingUpdateResponse).MeetingInfoList[0].MeetingID)
		log.Println(result.(qqmeeting.MeetingUpdateResponse).MeetingInfoList[0].MeetingCode)
	}
}

func TestMeeting_Do_QueryParticipants(t *testing.T) {

	r, err := meeting.Do(qqmeeting.MeetingQueryParticipantsRequest{
		UserID:    "13800138000",
		MeetingID: "7814232111865567045",
	})
	if err != nil {
		t.Error(err)
	} else {
		qqmeeting.PrintResponseJsonString(r)
	}
}


func TestMeeting_Do_QueryUserMeetingList(t *testing.T) {

	r, err := meeting.Do(qqmeeting.MeetingQueryUserMeetingListRequest{
		UserID:    "13800138000",
		InstanceID: 1,
	})
	if err != nil {
		t.Error(err)
	} else {
		qqmeeting.PrintResponseJsonString(r)
	}
}

