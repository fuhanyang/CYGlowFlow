package workflowService

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	workflow "github.com/fuhanyang/CYGlowFlow/app/gateway/hertz_gen/gateway/workflow"
)

type RetryExecutionService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewRetryExecutionService(Context context.Context, RequestContext *app.RequestContext) *RetryExecutionService {
	return &RetryExecutionService{RequestContext: RequestContext, Context: Context}
}

func (h *RetryExecutionService) Run(req *workflow.ExecutionActionReq) (resp *workflow.ExecutionActionResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
