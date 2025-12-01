package userService

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	user "github.com/fuhanyang/CYGlowFlow/app/gateway/hertz_gen/gateway/user"
)

type SearchUserService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewSearchUserService(ctx context.Context, requestContext *app.RequestContext) *SearchUserService {
	return &SearchUserService{RequestContext: requestContext, Context: ctx}
}

func (h *SearchUserService) Run(req *user.SearchUserReq) (resp *user.SearchUserResp, err error) {
	// defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
