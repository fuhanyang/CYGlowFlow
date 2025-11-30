package workflowService

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	workflow "github.com/fuhanyang/CYGlowFlow/app/gateway/hertz_gen/gateway/workflow"
)

type ExecuteWorkflowService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewExecuteWorkflowService(Context context.Context, RequestContext *app.RequestContext) *ExecuteWorkflowService {
	return &ExecuteWorkflowService{RequestContext: RequestContext, Context: Context}
}

func (h *ExecuteWorkflowService) Run(req *workflow.ExecuteWorkflowReq) (resp *workflow.ExecuteWorkflowResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
