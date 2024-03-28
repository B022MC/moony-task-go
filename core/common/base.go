package common

import (
	"encoding/json"
)

type BaseRequest struct {
	Page int    `form:"page"`
	Size int    `form:"size"`
	Sort string `form:"sort"`
}

type BaseResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type CommonResponse struct {
	BaseResponse
	Data interface{} `json:"data"`
}

type RawResponse struct {
	BaseResponse
	Data json.RawMessage `json:"data"`
}

type ListResponse struct {
	BaseResponse
	Data struct {
		Total       int64       `json:"total"`
		TotalAmount int64       `json:"totalAmount,omitempty"`
		Items       interface{} `json:"items"`
	} `json:"data,omitempty"`
}

var (
	rspOk *BaseResponse
)

func init() {
	rspOk = new(BaseResponse)
	rspOk.Message = "OK"
	rspOk.Code = 0
}

func NewRspOk() *BaseResponse {
	return rspOk
}

func NewRawRsp(raw []byte) *RawResponse {
	rsp := &RawResponse{
		BaseResponse: BaseResponse{
			Code:    rspOk.Code,
			Message: rspOk.Message,
		},
	}
	rsp.Data = raw
	return rsp
}

func NewRsp(data interface{}) *CommonResponse {
	return &CommonResponse{
		BaseResponse: BaseResponse{
			Code:    rspOk.Code,
			Message: rspOk.Message,
		},
		Data: data,
	}
}

func NewListRsp(total int64, items interface{}) *ListResponse {
	rsp := &ListResponse{
		BaseResponse: BaseResponse{
			Code:    rspOk.Code,
			Message: rspOk.Message,
		},
	}
	rsp.Data.Total = total
	rsp.Data.Items = items
	return rsp
}
