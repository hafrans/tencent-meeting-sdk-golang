package qqmeeting

// API LIST
var (
	RequestDescriptorMeetingCreate                = MeetingRequestDescriptor{"/meetings", "POST", "Create"}
	RequestDescriptorMeetingQueryByID             = MeetingRequestDescriptor{"/meetings/%s", "GET", "ID"}
	RequestDescriptorMeetingQueryByCode           = MeetingRequestDescriptor{"/meetings", "GET", "Code"}
	RequestDescriptorMeetingCancelByID            = MeetingRequestDescriptor{"/meetings/%s/cancel", "POST", "CANCEL"}
	RequestDescriptorMeetingUpdateByID            = MeetingRequestDescriptor{"/meetings/%s", "PUT", "UPDATE"}
	RequestDescriptorMeetingQueryParticipantsByID = MeetingRequestDescriptor{"/meetings/%s/participants", "GET", "QUERY"}
	RequestDescriptorMeetingQueryUserMeetingList  = MeetingRequestDescriptor{"/meetings", "GET", "QUERY_MEETING_LIST"}
	RequestDescriptorUserCreate                   = MeetingRequestDescriptor{"/users", "POST", "Create"}
	RequestDescriptorUserDetailQuery              = MeetingRequestDescriptor{"/users/%s", "GET", "QUERY"}
	RequestDescriptorUserUpdate                   = MeetingRequestDescriptor{"/users/%s", "PUT", "UPDATE"}
	RequestDescriptorUserDelete                   = MeetingRequestDescriptor{"/users/%s", "DELETE", "DELETE"}
	RequestDescriptorUserList                     = MeetingRequestDescriptor{"/users/list", "GET", "QUERY"}
)

// UserObj  用户对象 https://cloud.tencent.com/document/product/1095/42417
type UserObj struct {
	UserID      string `json:"userid"`              // 调用方用于标示用户的唯一 ID
	IsAnonymous bool   `json:"is_anonymous"`        // 匿名入会
	NickName    string `json:"nick_name,omitempty"` // 用户匿名字符串
}

// 会议设置
type Settings struct {
	MuteEnableJoin   bool `json:"mute_enable_join"`            // 入会时静音
	AllowUnmuteSelf  bool `json:"allow_unmute_self"`           // 允许参会者取消静音
	MuteAll          bool `json:"mute_all,omitempty"`          // 全体静音
	HostVideo        bool `json:"host_video,omitempty"`        // 入会时主持人视频是否开启，暂时不支持。
	ParticipantVideo bool `json:"participant_video,omitempty"` // 入会时参会者视频是否开启，暂时不支持。
	EnableRecord     bool `json:"enable_record,omitempty"`     // 开启录播，暂时不支持。
	PlayIVROnLeave   bool `json:"play_ivr_on_leave,omitempty"` // 参会者离开时播放提示音。
	PlayIVROnJoin    bool `json:"play_ivr_on_join,omitempty"`  // 有新的入会者加入时播放提示音。
	LiveUrl          bool `json:"live_url,omitempty"`          // 开启直播, 暂时不支持。
}

// 用户信息
type UserInfo struct {
	Email    string `json:"email" binding:"required"`    // 邮箱地址
	Phone    string `json:"phone" binding:"required"`    // 手机号码
	Username string `json:"username" binding:"required"` // 用户昵称
	UserID   string `json:"userid" binding:"required"`   // 调用方用于表示用户的唯一ID
}

// 用户详情
type UserDetail struct {
	Username   string `json:"username"`    // 用户昵称
	UpdateTime string `json:"update_time"` // 更新时间
	Status     string `json:"status"`      // 用户状态，1：正常；2：删除
	Email      string `json:"email"`       // 邮箱地址
	Phone      string `json:"phone"`       // 手机号码
	UserID     string `json:"userid"`      // 调用方用于标示用户的唯一 ID
	AreaCode   string `json:"area"`        // 地区编码（国内默认86）
	AvatarUrl  string `json:"avatar_url"`  // 头像
}

// 会议信息
type MeetingInfo struct {
	MeetingID    string     `json:"meeting_id"`             // 会议的唯一标示
	MeetingCode  string     `json:"meeting_code"`           // 会议 App 的呼入号码
	Subject      string     `json:"subject"`                // 会议主题
	Hosts        []*UserObj `json:"hosts,omitempty"`        // 会议主持人的用户 ID，如果没有指定，主持人被设定为参数 userid 的用户，即 API 调用者。
	Participants []*UserObj `json:"participants,omitempty"` // 会议邀请的参会者，可为空
	StartTime    string     `json:"start_time"`             // 会议开始时间戳（单位秒）。
	EndTime      string     `json:"end_time"`               // 会议结束时间戳（单位秒）。
	Password     *string    `json:"password,omitempty"`     // 会议密码，可不填
	JoinUrl      string     `json:"join_url"`               // 加入会议　URL（点击链接直接加入会议）
	Settings     *Settings  `json:"settings,omitempty"`     // 会议设置
}

// 带状态的会议信息
type MeetingQueryInfo struct {
	MeetingInfo
	Status string `json:"status"` // 当前会议状态
}

// 简要会议信息
type SimplifiedMeetingInfo struct {
	MeetingID   string `json:"meeting_id"`
	MeetingCode string `json:"meeting_code"`
}

// 参会人员信息
type MeetingParticipants struct {
	UserID                string `json:"userid"`    // 参会者用户 ID。
	Base64EncodedUsername string `json:"user_name"` // 入会用户名（base64）
	SHA256HashedPhone     string `json:"phone"`     // 参会者手机号 hash 值 SHA256（手机号/secretid）。
	JoinTime              string `json:"join_time"` // 参会者加入会议时间戳（单位秒）。
	LeftTime              string `json:"left_time"` // 参会者离开会议时间戳（单位秒）。
}

// 查询用户的会议列表中会议对象
type MeetingQueryUserListMeetingInfo struct {
	MeetingID       string     `json:"meeting_id"`        // 会议的唯一标示
	MeetingCode     string     `json:"meeting_code"`      // 会议 App 的呼入号码
	Subject         string     `json:"subject"`           // 会议主题
	Hosts           []*UserObj `json:"hosts,omitempty"`   // 会议主持人的用户 ID，如果没有指定，主持人被设定为参数 userid 的用户，即 API 调用者。
	StartTime       string     `json:"start_time"`        // 会议开始时间戳（单位秒）。
	EndTime         string     `json:"end_time"`          // 会议结束时间戳（单位秒）。
	JoinMeetingRole string     `json:"join_meeting_role"` // 查询者在会议中的角色
}
