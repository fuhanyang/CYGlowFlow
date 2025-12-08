package service

import (
	"context"
	common "github.com/fuhanyang/CYGlowFlow/rpc_gen/kitex_gen/common"
	exeution_manager "github.com/fuhanyang/CYGlowFlow/rpc_gen/kitex_gen/workflow/infra/exeution_manager"
)

type CancelExecutionService struct {
	ctx context.Context
}

// NewCancelExecutionService new CancelExecutionService
func NewCancelExecutionService(ctx context.Context) *CancelExecutionService {
	return &CancelExecutionService{ctx: ctx}
}

// Run create note info
func (s *CancelExecutionService) Run(req *exeution_manager.CancelExecutionRequest) (resp *common.BaseResponse, err error) {
	// Finish your business logic.

	return
}
