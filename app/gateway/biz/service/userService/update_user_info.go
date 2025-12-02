package userService

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	user "github.com/fuhanyang/CYGlowFlow/app/gateway/hertz_gen/gateway/user"
)

type UpdateUserInfoService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdateUserInfoService(ctx context.Context, requestContext *app.RequestContext) *UpdateUserInfoService {
	return &UpdateUserInfoService{RequestContext: requestContext, Context: ctx}
}

func (h *UpdateUserInfoService) Run(req *user.UpdateUserInfoReq) (resp *user.UpdateUserInfoResp, err error) {
	return resp, nil
}
