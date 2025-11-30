package workflowService

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	workflow "github.com/fuhanyang/CYGlowFlow/app/gateway/hertz_gen/gateway/workflow"
)

type ListExecutionsService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewListExecutionsService(Context context.Context, RequestContext *app.RequestContext) *ListExecutionsService {
	return &ListExecutionsService{RequestContext: RequestContext, Context: Context}
}

func (h *ListExecutionsService) Run(req *workflow.ListExecutionsReq) (resp *workflow.ListExecutionsResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
