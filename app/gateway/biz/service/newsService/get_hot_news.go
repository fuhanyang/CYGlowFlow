package newsService

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	news "github.com/fuhanyang/CYGlowFlow/app/gateway/hertz_gen/gateway/news"
)

type GetHotNewsService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetHotNewsService(Context context.Context, RequestContext *app.RequestContext) *GetHotNewsService {
	return &GetHotNewsService{RequestContext: RequestContext, Context: Context}
}

func (h *GetHotNewsService) Run(req *news.GetHotNewsReq) (resp *news.GetHotNewsResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
