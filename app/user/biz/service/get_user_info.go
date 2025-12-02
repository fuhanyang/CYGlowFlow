package service

import (
	"context"
	"errors"

	"github.com/fuhanyang/CYGlowFlow/app/user/biz/dal/mysql"
	"github.com/fuhanyang/CYGlowFlow/app/user/biz/dal/mysql/model"
	"github.com/fuhanyang/CYGlowFlow/app/user/biz/dal/mysql/query"
	user "github.com/fuhanyang/CYGlowFlow/rpc_gen/kitex_gen/user"
)

type GetUserInfoService struct {
	ctx context.Context
} // NewGetUserInfoService new GetUserInfoService
func NewGetUserInfoService(ctx context.Context) *GetUserInfoService {
	return &GetUserInfoService{ctx: ctx}
}

// Run get user info
func (s *GetUserInfoService) Run(req *user.GetUserInfoReq) (resp *user.GetUserInfoResp, err error) {
	// 1. 参数校验
	if req.UserId == 0 && req.AccountNum == 0 {
		return nil, errors.New("user_id or account_num is required")
	}

	// 2. 查询用户基础信息
	q := query.Use(mysql.DB)
	u := q.User

	var userModel *model.User

	// 优先使用user_id查询，如果为0则使用account_num查询
	if req.UserId != 0 {
		userModel, err = u.WithContext(s.ctx).Where(u.ID.Eq(uint64(req.UserId))).First()
		if err != nil {
			return nil, err
		}
	} else {
		userModel, err = u.WithContext(s.ctx).Where(u.AccountNum.Eq(req.AccountNum)).First()
		if err != nil {
			return nil, err
		}
	}

	// 3. 检查用户是否存在
	if userModel == nil {
		return nil, errors.New("user not found")
	}

	// 4. 构建响应数据
	userInfo := &user.UserBaseInfo{
		UserId:     int64(userModel.ID),
		AccountNum: userModel.AccountNum,
		Name:       getStringValue(userModel.Name),
		Email:      getStringValue(userModel.Email),
		Phone:      getStringValue(userModel.Phone),
		AvatarUrl:  getStringValue(userModel.AvatarURL),
		Desc:       getStringValue(userModel.Desc),
	}

	// 5. 返回结果
	resp = &user.GetUserInfoResp{
		UserInfo: userInfo,
	}

	return resp, nil
}

// getStringValue 安全获取字符串指针的值
func getStringValue(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
