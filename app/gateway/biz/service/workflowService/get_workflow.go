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

func NewGetWorkflowService(ctx context.Context, requestContext *app.RequestContext) *GetWorkflowService {
	return &GetWorkflowService{RequestContext: requestContext, Context: ctx}
}

func (h *GetWorkflowService) Run(req *workflow.GetWorkflowReq) (resp *workflow.GetWorkflowResp, err error) {
	// defer func() {
	//  hlog.CtxInfof(h.Context, "req = %+v", req)
	//  hlog.CtxInfof(h.Context, "resp = %+v", resp)
	// }()
	// todo edit your code
	return
}
