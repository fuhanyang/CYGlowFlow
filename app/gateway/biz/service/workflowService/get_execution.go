package workflowService

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	workflow "github.com/fuhanyang/CYGlowFlow/app/gateway/hertz_gen/gateway/workflow"
)

type GetExecutionService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetExecutionService(Context context.Context, RequestContext *app.RequestContext) *GetExecutionService {
	return &GetExecutionService{RequestContext: RequestContext, Context: Context}
}

func (h *GetExecutionService) Run(req *workflow.GetExecutionReq) (resp *workflow.GetExecutionResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
