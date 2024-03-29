package service

import (
	"encoding/json"
	"github.com/gogap/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"moony-task-go/common/dao"
	"moony-task-go/common/model"
	"moony-task-go/core/common"
	"moony-task-go/core/config"
	"moony-task-go/utils"
	"time"
)

var (
	UserNameDefault   = "微信用户"
	UserAvatarDefault = "https://thirdwx.qlogo.cn/mmopen/vi_32/POgEwh4mIHO4nibH0KlMECNjjGxQUq24ZEaGT4poC6icRiccVGKSyXwibcPq4BWmiaIGuG1icwxaQX6grC9VemZoJ8rg/132"
)

type User struct {
}

func UserInstance() *User {
	return &User{}
}

// Login 登录
func (u *User) Login(session *common.Session, request model.LoginUserReq) (map[string]interface{}, error) {
	if request.LoginType == model.LoginTypeWeixin {
		return u.loginByWeixin(session, request)
	}
	return nil, errors.New("不支持的登录方式")
}

// loginByWeixin 处理微信小程序登录
func (u *User) loginByWeixin(session *common.Session, request model.LoginUserReq) (map[string]interface{}, error) {
	if request.Weixin.Code == "" {
		panic(config.ErrParam.New().Append("微信code不能为空"))
	}
	cfg := config.GetConfig().WxLogin

	// 调用 jscode2session 获取 session_key 和 openid
	sessionResponse, err := utils.Jscode2session(cfg.AppId, cfg.AppSecret, request.Weixin.Code)
	if err != nil {
		log.Errorf("jscode2session code[%s] error: %s", request.Weixin.Code, err.Error())
		panic(config.ErrInternal.New().Append(err))
	}

	// 根据openid查找用户
	user, err := dao.UserInstance().GetByUnionid(sessionResponse.Unionid)
	if err != nil {
		log.Errorf("dao.UserInstance().GetByUnionid unionid:%s error:%s", sessionResponse.Unionid, err.Error())
		panic(config.ErrDb.New().Append(err))
	}

	if user != nil {
		// 更新现有用户信息
		user.NickName = request.Weixin.UserInfo.NickName
		user.Avatar = request.Weixin.UserInfo.AvatarUrl
		user.ActiveTime = time.Now().Unix()
	} else {
		// 如果用户不存在，则创建一个新用户
		user = &model.User{
			Unionid:    sessionResponse.Unionid,
			NickName:   request.Weixin.UserInfo.NickName,  // 微信不直接提供昵称，需要另外获取或设置默认值
			Avatar:     request.Weixin.UserInfo.AvatarUrl, // 同上
			ActiveTime: time.Now().Unix(),
		}
		// 这里应添加用户创建逻辑
		// err := dao.UserInstance().Create(user)
	}

	if err := dao.UserInstance().Update(user); err != nil {
		log.Errorf("loginByWeixin dao.UserInstance().Update error=[%s]", err.Error())
		return nil, err
	}

	token := common.CreateToken(user)
	data := make(map[string]interface{})
	data["token"] = token
	data["userInfo"] = u.Format(user)

	// 更新用户的Token信息
	user.TokenMd5 = utils.Md5(token)
	user.UpdateTime = time.Now().Unix()
	if err := dao.UserInstance().Update(user); err != nil {
		log.Errorf("loginByWeixin dao.UserInstance().Update err=%s", err.Error())
		return nil, err
	}

	return data, nil
}

// GetUserInfo 获取用户详情
func (u *User) GetUserInfo(session *common.Session) (map[string]interface{}, error) {
	user, err := dao.UserInstance().Get(session.GetUserId())
	if err != nil {
		log.Errorf("GetUserInfo userId=%d error=[%s]", session.GetUserId(), err.Error())
		return nil, err
	}

	return u.Format(user), nil
}

// SaveUserInfo 更新用户信息
func (u *User) SaveUserInfo(session *common.Session, params map[string]interface{}) error {
	user, err := dao.UserInstance().Get(session.GetUserId())
	if err != nil {
		log.Errorf("SaveUserInfo dao.UserInstance().Get error=[%s]", err.Error())
		return err
	}
	if user == nil {
		return errors.New("用户信息不存在")
	}
	if params == nil {
		return nil
	}

	userCnf := make(map[string]interface{})
	_ = json.Unmarshal([]byte(user.Config), &userCnf)

	now := time.Now().Unix()
	for k, v := range params {
		if k == "name" {
			user.Name = cast.ToString(v)
			user.UpdateTime = now
		} else if k == "avatar" {
			user.Avatar = cast.ToString(v)
			user.UpdateTime = now
		} else if k == "phone" {
			user.Phone = cast.ToString(v)
			user.UpdateTime = now
		} else if k == "age" {
			user.Age = cast.ToInt(v)
			user.UpdateTime = now
		} else if k == "sex" || k == "gender" {
			user.Sex = cast.ToInt(v)
			user.UpdateTime = now
		} else {
			userCnf[k] = v
		}
	}
	user.Config = utils.EncodeJSON(userCnf)

	if err := dao.UserInstance().Update(user); err != nil {
		log.Errorf("save user param=[%s] err=[%s]", utils.EncodeJSON(user), err.Error())
		return err
	}

	return nil
}

// UpdateUserInfo 更新用户信息
func (u *User) UpdateUserInfo(userId int64, params map[string]interface{}) error {
	user, err := dao.UserInstance().Get(userId)
	if err != nil {
		log.Errorf("SaveUserInfo dao.UserInstance().Get error=[%s]", err.Error())
		return err
	}
	if user == nil {
		return errors.New("用户信息不存在")
	}
	if params == nil {
		return nil
	}

	userCnf := make(map[string]interface{})
	_ = json.Unmarshal([]byte(user.Config), &userCnf)

	now := time.Now().Unix()
	for k, v := range params {
		if k == "name" {
			user.Name = cast.ToString(v)
			user.UpdateTime = now
		} else if k == "avatar" {
			user.Avatar = cast.ToString(v)
			user.UpdateTime = now
		} else if k == "phone" {
			user.Phone = cast.ToString(v)
			user.UpdateTime = now
		} else if k == "age" {
			user.Age = cast.ToInt(v)
			user.UpdateTime = now
		} else if k == "sex" || k == "gender" {
			user.Sex = cast.ToInt(v)
			user.UpdateTime = now
		} else {
			userCnf[k] = v
		}
	}
	user.Config = utils.EncodeJSON(userCnf)

	if err := dao.UserInstance().Update(user); err != nil {
		log.Errorf("save user param=[%s] err=[%s]", utils.EncodeJSON(user), err.Error())
		return err
	}

	return nil
}

// Format 格式化用户返回数据
func (u *User) Format(user *model.User) map[string]interface{} {
	data := make(map[string]interface{})

	data["userId"] = "8022" + cast.ToString(user.Id)
	data["hashId"] = cast.ToString(user.HashId)
	data["name"] = user.Name
	data["avatar"] = user.Avatar
	data["phone"] = user.Phone
	data["guest"] = user.IsGuest()
	data["sex"] = user.Sex
	data["age"] = user.Age

	if user.Name == "" {
		data["name"] = UserNameDefault
	}
	if user.Avatar == "" {
		data["avatar"] = UserAvatarDefault
	}

	if user.CreateTime != 0 {
		data["createTime"] = utils.TimeToDateTime(user.CreateTime)
	}

	if user.UpdateTime != 0 {
		data["updateTime"] = utils.TimeToDateTime(user.UpdateTime)
	}

	if user.ActiveTime != 0 {
		data["activeTime"] = utils.TimeToDateTime(user.ActiveTime)
	}

	cfg := make(map[string]interface{})
	_ = json.Unmarshal([]byte(user.Config), &cfg)

	for k, v := range cfg {
		data[k] = v
	}

	return data
}
