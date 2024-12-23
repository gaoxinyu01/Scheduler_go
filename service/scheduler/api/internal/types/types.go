// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.2

package types

type AttendanceAddRequest struct {
	CheckInTime  int64  `json:"check_in_time"`  // 签到时间
	CheckInPhoto string `json:"check_in_photo"` // 签到图片
	UserId       string `json:"user_id"`        // 用户ID
}

type AttendanceDelRequest struct {
	Id string `path:"id"` // 考勤ID
}

type AttendanceInfoRequest struct {
	Id string `form:"id"` // 考勤ID
}

type AttendanceListRequest struct {
	Current     int64  `form:"current,default=1,optional"`        // 页码
	PageSize    int64  `form:"page_size,default=10,optional"`     // 页数
	Name        string `form:"name,optional"`                     // 考勤人
	UserId      string `form:"user_id,optional"`                  // 用户ID
	Date        string `form:"date,optional"`                     // 考勤日期
	CheckInTime int64  `form:"check_in_time,default=99,optional"` // 签到时间
	SignOffTime int64  `form:"sign_off_time,default=99,optional"` // 签退时间
	State       int64  `form:"state,default=99,optional"`         // 考勤状态 上班打卡:1,下班打卡:2,打卡正常:3,打卡异常:4
}

type AttendancePatchRequest struct {
	SignOffTime  int64  `json:"sign_off_time"`  // 签退时间
	SignOffPhoto string `json:"sign_off_photo"` // 签退图片
	UserId       string `json:"user_id"`        // 用户ID
}

type AttendanceUpRequest struct {
	Id           string `json:"id"`                      // 考勤ID
	Name         string `json:"name,optional"`           // 考勤人
	UserId       string `json:"user_id,optional"`        // 用户ID
	Date         string `json:"date,optional"`           // 考勤日期
	CheckInTime  int64  `json:"check_in_time,optional"`  // 签到时间
	CheckInPhoto string `json:"check_in_photo,optional"` // 签到图片
	SignOffTime  int64  `json:"sign_off_time,optional"`  // 签退时间
	SignOffPhoto string `json:"sign_off_photo,optional"` // 签退图片
	State        int64  `json:"state,optional"`          // 考勤状态 上班打卡:1,下班打卡:2,打卡正常:3,打卡异常:4
}

type DayAttendanceByTimeInfoRequest struct {
	StartTime int64  `form:"start_time"` // 开始时间
	EndTime   int64  `form:"end_time"`   // 结束时间
	UserId    string `form:"user_id"`    // 用户ID
}

type DayAttendanceInfoRequest struct {
	Date   string `form:"date"`    // 考勤日期
	UserId string `form:"user_id"` // 用户ID
}

type NotTeamUserRequest struct {
	ProductId int64 `json:"product_id"`
}

type Response struct {
	Code int64       `json:"code"` // 状态码
	Msg  string      `json:"msg"`  // 消息
	Data interface{} `json:"data"` // 数据
}

type SchedulingAddList struct {
	Time      int64    `json:"time"`
	Name      string   `json:"name"`
	StartTime int64    `json:"start_time"`
	EndTime   int64    `json:"end_time"`
	Colour    string   `json:"colour"`
	TeamId    string   `json:"team_id"`
	UserIds   []string `json:"user_ids"`
}

type SchedulingAddRequest struct {
	Data []SchedulingAddList `json:"data"`
}

type SchedulingDelRequest struct {
	Id string `path:"id"`
}

type SchedulingListRequest struct {
	Current   int64  `form:"current,default=1,optional"`
	PageSize  int64  `form:"page_size,default=10,optional"`
	Time      int64  `form:"time,optional"`
	Name      string `form:"name,optional"`
	StartTime int64  `form:"start_time,optional"`
	EndTime   int64  `form:"end_time,optional"`
	TeamName  string `form:"team_name,optional"`
	UserName  string `form:"user_name,optional"`
}

type SchedulingTypeAddRequest struct {
	Name      string `json:"name"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Colour    string `json:"colour,optional"`
	Remark    string `json:"remark"`
}

type SchedulingTypeDelRequest struct {
	Id string `path:"id"`
}

type SchedulingTypeListRequest struct {
	Current  int64 `form:"current,default=1,optional"`
	PageSize int64 `form:"page_size,default=10,optional"`
}

type SchedulingTypeUpRequest struct {
	Id        string `json:"id"`
	Name      string `json:"name,optional"`
	StartTime string `json:"start_time,optional"`
	EndTime   string `json:"end_time,optional"`
	Colour    string `json:"colour,optional"`
	Remark    string `json:"remark,optional"`
}

type SchedulingUpRequest struct {
	Id        string `json:"id"`
	Time      int64  `json:"time"`
	Name      string `json:"name"`
	StartTime int64  `json:"start_time"`
	EndTime   int64  `json:"end_time"`
	Colour    string `json:"colour"`
	TeamId    string `json:"team_id"`
	UserId    string `json:"user_id"`
}

type TeamAddRequest struct {
	UserIds    []string `json:"user_ids"`
	TeamTypeId string   `json:"team_type_id"`
}

type TeamDelRequest struct {
	Id string `path:"id"`
}

type TeamListRequest struct {
	Current    int64  `form:"current,default=1,optional"`
	PageSize   int64  `form:"page_size,default=10,optional"`
	NickName   string `form:"nick_name,optional"`
	Major      string `form:"major,optional"`
	Position   string `form:"position,optional"`
	Telephone  string `form:"telephone,optional"`
	TeamTypeId string `form:"team_type_id"`
}

type TeamTypeAddRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type TeamTypeDelRequest struct {
	Id string `path:"id"`
}

type TeamTypeListRequest struct {
	Current     int64  `form:"current,default=1,optional"`
	PageSize    int64  `form:"page_size,default=10,optional"`
	Name        string `form:"name,optional"`
	Description string `form:"description,optional"`
}

type TeamTypeUpRequest struct {
	Id          string `json:"id"`
	Name        string `json:"name,optional"`
	Description string `json:"description,optional"`
}
