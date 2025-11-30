package friendService

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	friend "github.com/fuhanyang/CYGlowFlow/app/gateway/hertz_gen/gateway/friend"
)

type GetFriendListService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetFriendListService(Context context.Context, RequestContext *app.RequestContext) *GetFriendListService {
	return &GetFriendListService{RequestContext: RequestContext, Context: Context}
}

func (h *GetFriendListService) Run(req *friend.GetFriendListReq) (resp *friend.GetFriendListResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
