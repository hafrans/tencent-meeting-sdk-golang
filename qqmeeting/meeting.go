package qqmeeting

// 创建会议
type MeetingCreateRequest struct {
	UserID     string     `json:"userid"`             // 调用方用于标示用户的唯一 ID
	InstanceID int        `json:"instanceid"`         // 用户的终端设备类型
	Subject    string     `json:"subject"`            // 会议主题
	Hosts      []*UserObj `json:"hosts,omitempty"`    // 会议主持人的用户 ID，如果没有指定，主持人被设定为参数 userid 的用户，即 API 调用者。
	Type       int        `json:"type"`               // 会议类型
	Invitees   []*UserObj `json:"invitees,omitempty"` // 会议邀请的参会者，可为空
	StartTime  string     `json:"start_time"`         // 会议开始时间戳（单位秒）。
	EndTime    string     `json:"end_time"`           // 会议结束时间戳（单位秒）。
	Password   string     `json:"password,omitempty"` // 会议密码，可不填
	Settings   *Settings  `json:"settings,omitempty"` // 会议设置
}

func (req MeetingCreateRequest) getDescriptor() *MeetingRequestDescriptor {
	return &RequestDescriptorMeetingCreate
}

func (req MeetingCreateRequest) fillPlaceholder(args ...interface{}) string {
	return req.getDescriptor().Url
}

type MeetingCreateResponse struct {
	NextPos             int            `json:"next_pos"`
	Remaining           int            `json:"remaining"`
	MeetingNumber       int            `json:"meeting_number"`    // 会议数量
	MeetingCreationInfo []*MeetingInfo `json:"meeting_info_list"` // 预约会议列表
}

// 通过会议ID获取会议信息
type MeetingQueryByIDRequest struct {
	MeetingID  string `param:"meeting_id"`
	UserID     string `query:"userid"`
	InstanceID int    `query:"instanceid"`
}

func (req MeetingQueryByIDRequest) getDescriptor() *MeetingRequestDescriptor {
	return &RequestDescriptorMeetingQueryByID
}

func (req MeetingQueryByIDRequest) fillPlaceholder(args ...interface{}) string {
	return defaultPlaceholderFiller(req, args...)
}

// 通过Code获取会议信息
//
// BUG(hafrans) 腾讯会议API接口报错，不知道什么问题，本地HTTP抓包未发现异常。
type MeetingQueryByCodeRequest struct {
	MeetingCode string `json:"-" query:"meeting_code"`
	UserID      string `json:"-" query:"userid"`
	InstanceID  int    `json:"-" query:"instanceid"`
}

func (req MeetingQueryByCodeRequest) getDescriptor() *MeetingRequestDescriptor {
	return &RequestDescriptorMeetingQueryByCode
}

func (req MeetingQueryByCodeRequest) fillPlaceholder(args ...interface{}) string {
	return req.getDescriptor().Url
}

type MeetingQueryResponse struct {
	NextPos         int                 `json:"next_pos"`
	Remaining       int                 `json:"remaining"`
	MeetingNumber   int                 `json:"meeting_number"`    // 会议数量
	MeetingInfoList []*MeetingQueryInfo `json:"meeting_info_list"` // 会议列表
}

// 取消会议
type MeetingCancelRequest struct {
	MeetingID    string `json:"-" param:"meeting_id"`    // 会议的唯一 ID。
	UserID       string `json:"userid"`                  // 调用方用于标示用户的唯一 ID
	InstanceID   int    `json:"instanceid"`              // 用户的终端设备类型
	ReasonCode   int    `json:"reason_code"`             // 原因代码，可为用户自定义
	ReasonDetail string `json:"reason_detail,omitempty"` // 详细取消原因描述。
}

func (req MeetingCancelRequest) getDescriptor() *MeetingRequestDescriptor {
	return &RequestDescriptorMeetingCancelByID
}

func (req MeetingCancelRequest) fillPlaceholder(args ...interface{}) string {
	return defaultPlaceholderFiller(req, args...)
}

// MeetingUpdateRequest 修改某指定会议的会议信息
type MeetingUpdateRequest struct {
	MeetingID  string     `json:"-" param:"meeting_id"`
	UserID     string     `json:"userid"`               // 调用方用于标示用户的唯一 ID
	InstanceID int        `json:"instanceid"`           // 用户的终端设备类型
	Subject    string     `json:"subject"`              // 会议主题
	Hosts      []*UserObj `json:"hosts,omitempty"`      // 会议主持人的用户 ID，如果没有指定，主持人被设定为参数 userid 的用户，即 API 调用者。
	Invitees   []*UserObj `json:"invitees,omitempty"`   // 会议邀请的参会者，可为空
	StartTime  string     `json:"start_time,omitempty"` // 会议开始时间戳（单位秒）。
	EndTime    string     `json:"end_time,omitempty"`   // 会议结束时间戳（单位秒）。
	Password   string     `json:"password,omitempty"`   // 会议密码，可不填,修改会议接口只支持对有密码会议的密码修改，暂不支持取消会议密码。
	Settings   *Settings  `json:"settings,omitempty"`   // 会议设置,不修改就不填写
}

func (req MeetingUpdateRequest) getDescriptor() *MeetingRequestDescriptor {
	return &RequestDescriptorMeetingUpdateByID
}

func (req MeetingUpdateRequest) fillPlaceholder(args ...interface{}) string {
	return defaultPlaceholderFiller(req, args...)
}

type MeetingUpdateResponse struct {
	NextPos         int                      `json:"next_pos"`
	Remaining       int                      `json:"remaining"`
	MeetingNumber   int                      `json:"meeting_number"`    // 会议数量
	MeetingInfoList []*SimplifiedMeetingInfo `json:"meeting_info_list"` // 会议列表
}

// MeetingQueryParticipantsRequest
// 会议拥有者获取某指定会议的参与人员列表。只有会议的拥有者即创建者可以查询参会成员的列表。其他用
// 户的调用会被拒绝。如果会议还未开始，调用此接口查询会返回空列表。
type MeetingQueryParticipantsRequest struct {
	MeetingID string `param:"meeting_id"`
	UserID    string `query:"userid"`
}

func (req MeetingQueryParticipantsRequest) getDescriptor() *MeetingRequestDescriptor {
	return &RequestDescriptorMeetingQueryParticipantsByID
}

func (req MeetingQueryParticipantsRequest) fillPlaceholder(args ...interface{}) string {
	return defaultPlaceholderFiller(req, args...)
}

type MeetingQueryParticipantsResponse struct {
	MeetingID         string                 `json:"meeting_id"`          // 会议的唯一 ID
	MeetingCode       string                 `json:"meeting_code"`        // 会议号码
	Subject           string                 `json:"subject"`             // 会议主题
	ScheduleStartTime string                 `json:"schedule_start_time"` // 预定会议开始时间戳（单位秒）
	ScheduleEndTime   string                 `json:"schedule_end_time"`   // 预定会议结束时间戳（单位秒）
	Participants      []*MeetingParticipants `json:"participants"`        // 参会人员对象数组。
	HasRemaining      bool                   `json:"has_remaining"`
}

// 获取某指定用户的会议列表
type MeetingQueryUserMeetingListRequest struct {
	UserID     string `query:"userid"`
	InstanceID int    `query:"instanceid"`
}

func (req MeetingQueryUserMeetingListRequest) getDescriptor() *MeetingRequestDescriptor {
	return &RequestDescriptorMeetingQueryUserMeetingList
}

func (req MeetingQueryUserMeetingListRequest) fillPlaceholder(args ...interface{}) string {
	return req.getDescriptor().Url
}

type MeetingQueryUserMeetingListResponse struct {
	NextPos         int                                `json:"next_pos"`
	Remaining       int                                `json:"remaining"`
	MeetingNumber   int                                `json:"meeting_number"`
	MeetingInfoList []*MeetingQueryUserListMeetingInfo `json:"meeting_info_list"`
}
