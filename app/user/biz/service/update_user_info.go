package service

import (
	"context"
	"errors"
	"strings"

	"github.com/fuhanyang/CYGlowFlow/app/user/biz/dal/mysql"
	"github.com/fuhanyang/CYGlowFlow/app/user/biz/dal/mysql/query"
	user "github.com/fuhanyang/CYGlowFlow/rpc_gen/kitex_gen/user"
)

type UpdateUserInfoService struct {
	ctx context.Context
} // NewUpdateUserInfoService new UpdateUserInfoService
func NewUpdateUserInfoService(ctx context.Context) *UpdateUserInfoService {
	return &UpdateUserInfoService{ctx: ctx}
}

// Run update user info
func (s *UpdateUserInfoService) Run(req *user.UpdateUserInfoReq) (resp *user.UpdateUserInfoResp, err error) {
	// 1. 参数校验
	if req.UserId == 0 {
		return nil, errors.New("user_id is required")
	}

	// 2. 查询用户是否存在
	q := query.Use(mysql.DB)
	u := q.User

	userModel, err := u.WithContext(s.ctx).Where(u.ID.Eq(uint64(req.UserId))).First()
	if err != nil {
		return nil, errors.New("user not found")
	}

	if userModel == nil {
		return nil, errors.New("user not found")
	}

	// 3. 构建更新数据
	updates := make(map[string]interface{})

	// 只更新非空字段
	if req.Name != "" {
		updates["name"] = req.Name
	}
	if req.Email != "" {
		// 邮箱格式简单校验
		if !strings.Contains(req.Email, "@") {
			return nil, errors.New("invalid email format")
		}
		updates["email"] = req.Email
	}
	if req.Phone != "" {
		updates["phone"] = req.Phone
	}
	if req.AvatarUrl != "" {
		updates["avatar_url"] = req.AvatarUrl
	}
	if req.Desc != "" {
		updates["desc"] = req.Desc
	}

	// 4. 检查是否有需要更新的字段
	if len(updates) == 0 {
		return &user.UpdateUserInfoResp{Success: true}, nil
	}

	// 5. 执行更新操作
	result, err := u.WithContext(s.ctx).Where(u.ID.Eq(uint64(req.UserId))).Updates(updates)
	if err != nil {
		return nil, err
	}

	// 6. 检查更新是否成功
	if result.Error != nil {
		return nil, result.Error
	}

	// 7. 返回结果
	resp = &user.UpdateUserInfoResp{
		Success: true,
	}

	return resp, nil
}
