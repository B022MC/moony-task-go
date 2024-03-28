package model

type ComReq struct {
	Page int `form:"page" json:"page"`
	Size int `form:"size" json:"size"`
}

func (c *ComReq) Default() {
	if c.Page <= 0 {
		c.Page = 1
	}
	if c.Size <= 0 {
		c.Size = 10
	}
}

type Region struct {
	ProvinceId int64 `form:"provinceId" json:"provinceId"`
	CityId     int64 `form:"cityId" json:"cityId"`
	AreaId     int64 `form:"areaId" json:"areaId"`
}
