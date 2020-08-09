package qqmeeting

// Entrance
const (
	ApiHost = "https://api.meeting.qq.com/v1"
)

// 企业用户管理
const (
	DisableRegistered = iota
	EnableRegistered
)

// 会议角色
const (
	MeetingRoleCreator = "creator" // 创建者
	MeetingRoleHoster  = "hoster"  // 主持人
	MeetingRoleInvitee = "invitee" // 被邀请者
)

// 会议类型
const (
	MeetingTypeBookingMeeting = iota // 预约会议类型
	MeetingTypeQuickMeeting          // 快速会议类型
)

// 设备类型
const (
	InstancePC = iota + 1
	InstanceMac
	InstanceAndroid
	InstanceIOS
	InstanceWeb
	InstanceIPad
	InstanceAndroidPad
	InstanceMicroProgram
)

// 错误码
const (
	ErrTinyIdOrMeetingId       = 9002
	ErrMeetingNotExists        = 9003
	ErrMeetingCreateExceed     = 9008
	ErrMeetingQueryExceed      = 9061
	ErrApiCallUnknownType      = 10000
	ErrApiCallBadParameter     = 10001
	ErrAppVersionForbidden     = 10005
	ErrUserAlreadyExists       = 20002
	ErrUserUnavailable         = 20003
	ErrInvalidPhone            = 40000
	ErrInvalidEmail            = 41001
	ErrEmailUsed               = 41002
	ErrPhoneUsed               = 41003
	ErrCorpID                  = 50000
	ErrCorpUnavailable         = 50001
	ErrXTcTimestamp            = 190300
	ErrRequestReplay           = 190301
	ErrUnauthenticatedSecret   = 190303
	ErrCallMinuteExceed        = 190310
	ErrCallDayExceed           = 190311
	ErrCallParticularDayExceed = 190312
	ErrApiRequiredInfoNotFound = 200001
	ErrApiReplay               = 200002
	ErrApiBadSignature         = 200003
	ErrApiNotSupportRequest    = 200004
	ErrJsonSchemeInvalid       = 200005
	ErrApiBadRequestParameter  = 200006
)

// 会议状态
const (
	MeetingStateInvalid   = "MEETING_STATE_INVALID"   // 非法或未知的会议状态，错误状态
	MeetingStateInit      = "MEETING_STATE_INIT"      // 会议的初始状态，表示还没有人入会
	MeetingStateCancelled = "MEETING_STATE_CANCELLED" // 会议已取消
	MeetingStateStarted   = "MEETING_STATE_STARTED"   // 会议已开始，有人入会
	MeetingStateEnded     = "MEETING_STATE_ENDED"     // 会议已结束
	MeetingStateRecycled  = "MEETING_STATE_RECYCLED"  // 会议号已被回收
	MeetingStateNull      = "MEETING_STATE_NULL"      // 未知状态
)
