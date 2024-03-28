package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"moony-task-go/common/model"
	"moony-task-go/core/config"
	"moony-task-go/utils"
	"os"
	"time"
)

var (
	ContextSession = "session"
	tokenPass      = []byte("b022mc")
)

type UserHeader struct {
	AppId       int64  `header:"x-app-id" json:"appId"`
	Version     string `header:"x-version" json:"version"`
	Platform    string `header:"x-platform" json:"platform"`
	Channel     string `header:"x-channel" json:"channel"`
	DeviceId    string `header:"x-device-id" json:"deviceId"`
	RemoteIp    string `header:"x-public-ip" json:"remoteIp"`
	MobileBrand string `header:"x-mobile-brand" json:"mobileBrand"`
	MobileModel string `header:"x-mobile-model" json:"mobileModel"`
	ShareBossId string `header:"x-share-boss-id" json:"shareBossId"`
	Token       string `header:"x-token" json:"-"`
}

type UserToken struct {
	jwt.StandardClaims
	UserId int64
	HashId uint64
}

type Session struct {
	UserHeader *UserHeader
	UserToken  *UserToken
	Experiment *Experiment
	User       *model.User
	Context    *gin.Context
}

func CreateToken(user *model.User) string {
	hostname, _ := os.Hostname()
	claims := UserToken{
		UserId: user.Id,
		HashId: user.HashId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + 3650*24*3600,
			NotBefore: time.Now().Unix() - 60,
			IssuedAt:  time.Now().Unix(),
			Issuer:    hostname,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)

	tk, err := token.SignedString(tokenPass)
	if err != nil {
		panic(config.ErrInternal.New().Append(err))
	}
	return tk
}

func ParseToken(tokenStr string) *UserToken {
	tokenClaims, err := jwt.ParseWithClaims(tokenStr, &UserToken{},
		func(token *jwt.Token) (interface{}, error) {
			return tokenPass, nil
		})
	if err != nil {
		panic(config.ErrTokenFormat.New().Append(err))
	}

	token := tokenClaims.Claims.(*UserToken)
	return token
}

func (s *Session) GetHashId() uint64 {
	if s.UserToken != nil {
		return s.UserToken.HashId
	}
	if s.UserToken != nil {
		return utils.Hash64(s.UserHeader.DeviceId)
	}
	return utils.Hash64(fmt.Sprintf("%d", time.Now().UnixNano()))
}

func (s *Session) GetUserId() int64 {
	if s.UserToken != nil {
		return s.UserToken.UserId
	}
	return 0
}
