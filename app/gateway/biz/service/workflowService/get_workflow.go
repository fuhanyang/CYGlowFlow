package workflowService

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	workflow "github.com/fuhanyang/CYGlowFlow/app/gateway/hertz_gen/gateway/workflow"
)

type GetWorkflowService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetWorkflowService(Context context.Context, RequestContext *app.RequestContext) *GetWorkflowService {
	return &GetWorkflowService{RequestContext: RequestContext, Context: Context}
}

func (h *GetWorkflowService) Run(req *workflow.GetWorkflowReq) (resp *workflow.GetWorkflowResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
