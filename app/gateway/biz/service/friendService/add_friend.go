package friendService

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	friend "github.com/fuhanyang/CYGlowFlow/app/gateway/hertz_gen/gateway/friend"
)

type AddFriendService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAddFriendService(Context context.Context, RequestContext *app.RequestContext) *AddFriendService {
	return &AddFriendService{RequestContext: RequestContext, Context: Context}
}

func (h *AddFriendService) Run(req *friend.AddFriendReq) (resp *friend.AddFriendResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
