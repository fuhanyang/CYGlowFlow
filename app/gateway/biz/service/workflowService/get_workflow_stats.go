package workflowService

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	workflow "github.com/fuhanyang/CYGlowFlow/app/gateway/hertz_gen/gateway/workflow"
)

type GetWorkflowStatsService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetWorkflowStatsService(Context context.Context, RequestContext *app.RequestContext) *GetWorkflowStatsService {
	return &GetWorkflowStatsService{RequestContext: RequestContext, Context: Context}
}

func (h *GetWorkflowStatsService) Run(req *workflow.Empty) (resp *workflow.GetWorkflowStatsResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
