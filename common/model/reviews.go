package model

type Review struct {
	ReviewId   int64  `json:"review_id" form:"reviewId"`     // 评价ID
	JobId      int64  `json:"job_id" form:"jobId"`           // 兼职信息ID
	UserId     int64  `json:"user_id" form:"userId"`         // 评价用户ID
	Rating     int    `json:"rating" form:"rating"`          // 评分
	Comment    string `json:"comment" form:"comment"`        // 评论
	ReviewTime int64  `json:"review_time" form:"reviewTime"` // 评价时间
}
