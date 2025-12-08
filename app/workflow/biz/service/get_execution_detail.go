package service

import (
	"context"
	common "github.com/fuhanyang/CYGlowFlow/rpc_gen/kitex_gen/common"
	exeution_manager "github.com/fuhanyang/CYGlowFlow/rpc_gen/kitex_gen/workflow/infra/exeution_manager"
)

type GetExecutionDetailService struct {
	ctx context.Context
}

// NewGetExecutionDetailService new GetExecutionDetailService
func NewGetExecutionDetailService(ctx context.Context) *GetExecutionDetailService {
	return &GetExecutionDetailService{ctx: ctx}
}

// Run create note info
func (s *GetExecutionDetailService) Run(req *exeution_manager.GetExecutionDetailRequest) (resp *common.BaseResponse, err error) {
	// Finish your business logic.

	return
}
