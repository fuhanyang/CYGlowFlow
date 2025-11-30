package service

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"

	"github.com/fuhanyang/CYGlowFlow/app/user/biz/dal/mysql"
	"github.com/fuhanyang/CYGlowFlow/app/user/biz/dal/mysql/query"
	user "github.com/fuhanyang/CYGlowFlow/rpc_gen/kitex_gen/user"
)

type LoginService struct {
	ctx context.Context
} // NewLoginService new LoginService
func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{ctx: ctx}
}

// Run create note info
func (s *LoginService) Run(req *user.LoginReq) (resp *user.LoginResp, err error) {
	// 1. 根据账号查询用户
	q := query.Use(mysql.DB)
	u := q.User

	userInfo, err := u.WithContext(s.ctx).Where(u.AccountNum.Eq(req.AccountNum)).First()
	if err != nil {
		return nil, errors.New("user not found")
	}

	// 2. 校验密码
	hasher := md5.New()
	hasher.Write([]byte(req.Password))
	encryptedPassword := hex.EncodeToString(hasher.Sum(nil))

	if userInfo.Password != encryptedPassword {
		return nil, errors.New("wrong password")
	}

	// 3. 返回结果
	resp = &user.LoginResp{
		UserId:  int64(userInfo.ID),
		Success: true,
	}

	return resp, nil
}
