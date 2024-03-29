package model

// UserResume 查询结果的结构体
type UserResume struct {
	UserID          int64
	Name            string
	Avatar          string
	Phone           string
	Age             int
	Sex             string
	Educations      string
	WorkExperiences string
	Skills          string
	JobExpectations string
	SelfEvaluations string
}
