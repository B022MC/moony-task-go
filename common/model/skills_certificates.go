package model

import "time"

type SkillCertificate struct {
	Id                int64  `json:"id" form:"id"`                                // 技能/证书ID
	UserId            int64  `json:"user_id" form:"userId"`                       // 用户ID
	SkillName         string `json:"skill_name" form:"skillName"`                 // 技能/证书名称
	CertificateNumber string `json:"certificate_number" form:"certificateNumber"` // 证书编号
	CreateTime        int64  `json:"create_time" form:"createTime"`               // 创建时间
	UpdateTime        int64  `json:"update_time" form:"updateTime"`               // 更新时间
}

// SkillCertificateFormatted 定义了技能/证书的格式化输出结构
type SkillCertificateFormatted struct {
	Id                int64  `json:"id"`
	UserId            int64  `json:"user_id"`
	SkillName         string `json:"skill_name"`
	CertificateNumber string `json:"certificate_number"`
	CreateTime        string `json:"create_time"` // 格式化时间
	UpdateTime        string `json:"update_time"` // 格式化时间
}

func (sc *SkillCertificate) Format() SkillCertificateFormatted {
	// 格式化Unix时间戳为更易读的格式
	createTimeStr := time.Unix(sc.CreateTime, 0).Format("2006-01-02 15:04:05")
	updateTimeStr := time.Unix(sc.UpdateTime, 0).Format("2006-01-02 15:04:05")

	return SkillCertificateFormatted{
		Id:                sc.Id,
		UserId:            sc.UserId,
		SkillName:         sc.SkillName,
		CertificateNumber: sc.CertificateNumber,
		CreateTime:        createTimeStr,
		UpdateTime:        updateTimeStr,
	}
}
