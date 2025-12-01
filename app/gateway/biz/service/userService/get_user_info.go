package userService

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	user "github.com/fuhanyang/CYGlowFlow/app/gateway/hertz_gen/gateway/user"
)

type GetUserInfoService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetUserInfoService(ctx context.Context, requestContext *app.RequestContext) *GetUserInfoService {
	return &GetUserInfoService{RequestContext: requestContext, Context: ctx}
}

func (h *GetUserInfoService) Run(req *user.GetUserInfoReq) (resp *user.GetUserInfoResp, err error) {
	// defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
