package config

import "github.com/gogap/errors"

var (
	ErrOK       = errors.T(0, "操作成功")
	ErrPrivate  = errors.T(196, "没有权限")
	ErrDb       = errors.T(197, "数据库错误")
	ErrRemote   = errors.T(198, "远程错误")
	ErrInternal = errors.T(199, "内部错误")

	ErrParam       = errors.T(1000, "参数错误")
	ErrTokenFormat = errors.T(1001, "TOKEN错误")
	ErrTokenExpire = errors.T(1002, "登录过期")
	ErrUserDelete  = errors.T(1003, "用户已删除")
	ErrNoLogin     = errors.T(1004, "请先登录")

	ErrPassword     = errors.T(1005, "账号或密码错误")
	ErrVerifyCode   = errors.T(1006, "验证码错误")
	ErrNoUser       = errors.T(1007, "用户不存在")
	ErrNoData       = errors.T(1012, "数据不存在")
	ErrGoodsPrice   = errors.T(1013, "价格已变更")
	ErrSmsValid     = errors.T(1014, "短信有效期内不用重复发送")
	ErrSmsExhausted = errors.T(1015, "短信发送超出限制")
	ErrUnknown      = errors.T(1016, "未知错误")
	ErrNoMember     = errors.T(1017, "非会员，无权限查看")
	ErrIdCard       = errors.T(1018, "身份证校验失败")
	ErrExceedTimes  = errors.T(1019, "获取手机号超出限制")
	ErrUserDisable  = errors.T(1020, "禁止使用")
	ErrCardFail     = errors.T(1021, "身份证识别失败")
	ErrWxErr        = errors.T(1022, "微信授权失败")
	ErrAliyunErr    = errors.T(1023, "支付宝授权失败")
	ErrOrderLimit   = errors.T(1024, "已购买成功，请联系客服")
	ErrCardFailNum  = errors.T(1025, "身份证失败次数过多")

	ErrRequestLimit = errors.T(40003, "请稍后重试")
)
