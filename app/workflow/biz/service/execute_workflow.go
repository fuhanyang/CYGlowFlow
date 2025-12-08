package service

import (
	"context"
	common "github.com/fuhanyang/CYGlowFlow/rpc_gen/kitex_gen/common"
	exeution_manager "github.com/fuhanyang/CYGlowFlow/rpc_gen/kitex_gen/workflow/infra/exeution_manager"
)

type ExecuteWorkflowService struct {
	ctx context.Context
}

// NewExecuteWorkflowService new ExecuteWorkflowService
func NewExecuteWorkflowService(ctx context.Context) *ExecuteWorkflowService {
	return &ExecuteWorkflowService{ctx: ctx}
}

// Run create note info
func (s *ExecuteWorkflowService) Run(req *exeution_manager.ExecuteWorkflowRequest) (resp *common.BaseResponse, err error) {
	// Finish your business logic.

	return
}
