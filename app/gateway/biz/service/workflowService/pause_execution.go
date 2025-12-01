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

func NewPauseExecutionService(ctx context.Context, requestContext *app.RequestContext) *PauseExecutionService {
	return &PauseExecutionService{RequestContext: requestContext, Context: ctx}
}

func (h *PauseExecutionService) Run(req *workflow.ExecutionActionReq) (resp *workflow.ExecutionActionResp, err error) {
	// defer func() {
	//  hlog.CtxInfof(h.Context, "req = %+v", req)
	//  hlog.CtxInfof(h.Context, "resp = %+v", resp)
	// }()
	// todo edit your code
	return
}
