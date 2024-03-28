package model

type SkillCertificate struct {
	Id                int64  `json:"id" form:"id"`                                // 技能/证书ID
	UserId            int64  `json:"user_id" form:"userId"`                       // 用户ID
	SkillName         string `json:"skill_name" form:"skillName"`                 // 技能/证书名称
	CertificateNumber string `json:"certificate_number" form:"certificateNumber"` // 证书编号
	CreateTime        int64  `json:"create_time" form:"createTime"`               // 创建时间
	UpdateTime        int64  `json:"update_time" form:"updateTime"`               // 更新时间
}
