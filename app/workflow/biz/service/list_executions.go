package service

import (
	"context"
	common "github.com/fuhanyang/CYGlowFlow/rpc_gen/kitex_gen/common"
	exeution_manager "github.com/fuhanyang/CYGlowFlow/rpc_gen/kitex_gen/workflow/infra/exeution_manager"
)

type ListExecutionsService struct {
	ctx context.Context
}

// NewListExecutionsService new ListExecutionsService
func NewListExecutionsService(ctx context.Context) *ListExecutionsService {
	return &ListExecutionsService{ctx: ctx}
}

// Run create note info
func (s *ListExecutionsService) Run(req *exeution_manager.ExecutionQueryRequest) (resp *common.BaseResponse, err error) {
	// Finish your business logic.

	return
}
