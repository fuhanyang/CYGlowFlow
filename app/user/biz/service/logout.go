package service

import (
	"context"

	user "github.com/fuhanyang/CYGlowFlow/rpc_gen/kitex_gen/user"
)

type LogoutService struct {
	ctx context.Context
} // NewLogoutService new LogoutService
func NewLogoutService(ctx context.Context) *LogoutService {
	return &LogoutService{ctx: ctx}
}

// Run create note info
func (s *LogoutService) Run(req *user.LogoutReq) (resp *user.LogoutResp, err error) {
	// 实际上 JWT 无状态，服务端不需要做特殊处理
	// 如果需要做登出，通常是将 Token 加入黑名单 (Redis)，这里暂时只返回成功
	resp = &user.LogoutResp{
		Success: true,
	}
	return resp, nil
}
