package workflow

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/fuhanyang/CYGlowFlow/app/gateway/biz/service/workflowService"
	"github.com/fuhanyang/CYGlowFlow/app/gateway/biz/utils"
	workflow "github.com/fuhanyang/CYGlowFlow/app/gateway/hertz_gen/gateway/workflow"
)

// ListWorkflows .
// @router /workflows [GET]
func ListWorkflows(ctx context.Context, c *app.RequestContext) {
	var err error
	var req workflow.ListWorkflowsReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := workflowService.NewListWorkflowsService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// GetWorkflow .
// @router /workflows/:id [GET]
func GetWorkflow(ctx context.Context, c *app.RequestContext) {
	var err error
	var req workflow.GetWorkflowReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := workflowService.NewGetWorkflowService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// ExecuteWorkflow .
// @router /workflows/:id/execute [POST]
func ExecuteWorkflow(ctx context.Context, c *app.RequestContext) {
	var err error
	var req workflow.ExecuteWorkflowReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := workflowService.NewExecuteWorkflowService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// ListExecutions .
// @router /executions [GET]
func ListExecutions(ctx context.Context, c *app.RequestContext) {
	var err error
	var req workflow.ListExecutionsReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := workflowService.NewListExecutionsService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// GetExecution .
// @router /executions/:id [GET]
func GetExecution(ctx context.Context, c *app.RequestContext) {
	var err error
	var req workflow.GetExecutionReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := workflowService.NewGetExecutionService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// GetWorkflowStats .
// @router /workflow/stats [GET]
func GetWorkflowStats(ctx context.Context, c *app.RequestContext) {
	var err error
	var req workflow.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := workflowService.NewGetWorkflowStatsService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// CancelExecution .
// @router /executions/:id/cancel [POST]
func CancelExecution(ctx context.Context, c *app.RequestContext) {
	var err error
	var req workflow.ExecutionActionReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := workflowService.NewCancelExecutionService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// PauseExecution .
// @router /executions/:id/pause [POST]
func PauseExecution(ctx context.Context, c *app.RequestContext) {
	var err error
	var req workflow.ExecutionActionReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := workflowService.NewPauseExecutionService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// RetryExecution .
// @router /executions/:id/retry [POST]
func RetryExecution(ctx context.Context, c *app.RequestContext) {
	var err error
	var req workflow.ExecutionActionReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := workflowService.NewRetryExecutionService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
