package loginService

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	login "github.com/fuhanyang/CYGlowFlow/app/gateway/hertz_gen/gateway/login"
	"github.com/fuhanyang/CYGlowFlow/app/gateway/infra/rpc"
	"github.com/fuhanyang/CYGlowFlow/rpc_gen/kitex_gen/user"
)

type LogoffService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLogoffService(Context context.Context, RequestContext *app.RequestContext) *LogoffService {
	return &LogoffService{RequestContext: RequestContext, Context: Context}
}

func (h *LogoffService) Run(req *login.LogoffReq) (resp *login.LogoffResp, err error) {
	// 1. 调用 RPC
	_, err = rpc.UserClient.Logout(h.Context, &user.LogoutReq{
		// Token:  "", // 如果需要传递 Token
		UserId: req.AccountNum, // 注意：这里 IDL 定义可能是 AccountNum 但微服务可能需要 UserId，或者微服务通过 AccountNum 查找
	})
	if err != nil {
		resp = &login.LogoffResp{
			Code: consts.StatusInternalServerError,
			Msg:  err.Error(),
		}
		return resp, nil
	}

	// 2. 返回结果
	resp = &login.LogoffResp{
		Code: consts.StatusOK,
		Msg:  "success",
	}
	return resp, nil
}
