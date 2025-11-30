package workflowService

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	workflow "github.com/fuhanyang/CYGlowFlow/app/gateway/hertz_gen/gateway/workflow"
)

type PauseExecutionService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewPauseExecutionService(Context context.Context, RequestContext *app.RequestContext) *PauseExecutionService {
	return &PauseExecutionService{RequestContext: RequestContext, Context: Context}
}

func (h *PauseExecutionService) Run(req *workflow.ExecutionActionReq) (resp *workflow.ExecutionActionResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
