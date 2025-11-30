package chatService

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	chat "github.com/fuhanyang/CYGlowFlow/app/gateway/hertz_gen/gateway/chat"
)

type RefreshTextService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewRefreshTextService(Context context.Context, RequestContext *app.RequestContext) *RefreshTextService {
	return &RefreshTextService{RequestContext: RequestContext, Context: Context}
}

func (h *RefreshTextService) Run(req *chat.RefreshTextReq) (resp *chat.RefreshTextResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
