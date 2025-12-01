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

func NewRefreshTextService(ctx context.Context, requestContext *app.RequestContext) *RefreshTextService {
	return &RefreshTextService{RequestContext: requestContext, Context: ctx}
}

func (h *RefreshTextService) Run(req *chat.RefreshTextReq) (resp *chat.RefreshTextResp, err error) {
	// defer func() {
	//  hlog.CtxInfof(h.Context, "req = %+v", req)
	//  hlog.CtxInfof(h.Context, "resp = %+v", resp)
	// }()
	// todo edit your code
	return
}
