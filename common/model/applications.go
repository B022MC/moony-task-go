package model

type Application struct {
	ApplicationId int64 `json:"application_id" form:"applicationId"` // 申请ID
	JobId         int64 `json:"job_id" form:"jobId"`                 // 兼职信息ID
	UserId        int64 `json:"user_id" form:"userId"`               // 申请用户ID
	Status        int   `json:"status" form:"status"`                // 申请状态：1=待审核 2=接受 3=拒绝
	ApplyTime     int64 `json:"apply_time" form:"applyTime"`         // 申请时间
}
