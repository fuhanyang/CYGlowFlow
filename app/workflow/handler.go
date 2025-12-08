package main

import (
	"context"
	common "github.com/fuhanyang/CYGlowFlow/rpc_gen/kitex_gen/common"
	exeution_manager "github.com/fuhanyang/CYGlowFlow/rpc_gen/kitex_gen/workflow/infra/exeution_manager"
	"github.com/fuhanyang/CYGlowFlow/app/workflow/biz/service"
)

// ExecutionManagerServiceImpl implements the last service interface defined in the IDL.
type ExecutionManagerServiceImpl struct{}

// ExecuteWorkflow implements the ExecutionManagerServiceImpl interface.
func (s *ExecutionManagerServiceImpl) ExecuteWorkflow(ctx context.Context, req *exeution_manager.ExecuteWorkflowRequest) (resp *common.BaseResponse, err error) {
	resp, err = service.NewExecuteWorkflowService(ctx).Run(req)

	return resp, err
}

// ListExecutions implements the ExecutionManagerServiceImpl interface.
func (s *ExecutionManagerServiceImpl) ListExecutions(ctx context.Context, req *exeution_manager.ExecutionQueryRequest) (resp *common.BaseResponse, err error) {
	resp, err = service.NewListExecutionsService(ctx).Run(req)

	return resp, err
}

// GetExecutionDetail implements the ExecutionManagerServiceImpl interface.
func (s *ExecutionManagerServiceImpl) GetExecutionDetail(ctx context.Context, req *exeution_manager.GetExecutionDetailRequest) (resp *common.BaseResponse, err error) {
	resp, err = service.NewGetExecutionDetailService(ctx).Run(req)

	return resp, err
}

// CancelExecution implements the ExecutionManagerServiceImpl interface.
func (s *ExecutionManagerServiceImpl) CancelExecution(ctx context.Context, req *exeution_manager.CancelExecutionRequest) (resp *common.BaseResponse, err error) {
	resp, err = service.NewCancelExecutionService(ctx).Run(req)

	return resp, err
}
