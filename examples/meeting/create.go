package main

import (
	"fmt"
	"github.com/hafrans/tencent-meeting-sdk-golang/qqmeeting"
	"strconv"
	"time"
)

var	meeting = qqmeeting.Meeting{
	Registered: qqmeeting.EnableRegistered,
	Version:    "1.0.0",
	SecretKey:  "",
	AppID:      "",
	SdkID:      "",
	SecretID:   "",
}


func MeetingCreate() {

	response, err := meeting.Do(qqmeeting.MeetingCreateRequest{
		InstanceID: qqmeeting.InstancePC,
		UserID:     "13800138000",
		Hosts: []*qqmeeting.UserObj{
			{
				UserID: "13800138000",
			},
		},
		Subject:   "测试会议",
		StartTime: strconv.Itoa(int(time.Now().Unix() + 60)),
		EndTime:   strconv.Itoa(int(time.Now().Unix() + 360)),
		Settings: &qqmeeting.Settings{
			MuteEnableJoin:  true,
			AllowUnmuteSelf: true,
		},
	})

	if err != nil {
		fmt.Println(err)
		if e, ok := err.(qqmeeting.MeetingError); ok {
			fmt.Println("CODE:", e.Code)
			fmt.Println("MSG:", e.Message)
		}
	} else {
		result := response.(qqmeeting.MeetingCreateResponse)
		fmt.Println("会议主题", result.MeetingCreationInfo[0].Subject)
		fmt.Println("会议ID", result.MeetingCreationInfo[0].MeetingID)
		fmt.Println("会议号", result.MeetingCreationInfo[0].MeetingCode)
		fmt.Println("开始时间", result.MeetingCreationInfo[0].StartTime)
		fmt.Println("结束时间", result.MeetingCreationInfo[0].EndTime)
		fmt.Println("密码", result.MeetingCreationInfo[0].Password)
		fmt.Println("入会连接", result.MeetingCreationInfo[0].JoinUrl)

		/*
		会议主题 测试会议
		会议ID 9643171792387579848
		会议号 725214060
		开始时间 1596809701
		结束时间 1596810001
		密码 <nil>
		入会连接 https://meeting.tencent.com/s/qNl8n01a89f1
		*/

	}

}
