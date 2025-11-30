package userService

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	user "github.com/fuhanyang/CYGlowFlow/app/gateway/hertz_gen/gateway/user"
)

type UploadAvatarService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUploadAvatarService(Context context.Context, RequestContext *app.RequestContext) *UploadAvatarService {
	return &UploadAvatarService{RequestContext: RequestContext, Context: Context}
}

func (h *UploadAvatarService) Run(req *user.UploadAvatarReq) (resp *user.UploadAvatarResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
