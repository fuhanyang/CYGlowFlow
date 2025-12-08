package service

import (
	"context"
	"testing"
	common "github.com/fuhanyang/CYGlowFlow/rpc_gen/kitex_gen/common"
	exeution_manager "github.com/fuhanyang/CYGlowFlow/rpc_gen/kitex_gen/workflow/infra/exeution_manager"
)

func TestGetExecutionDetail_Run(t *testing.T) {
	ctx := context.Background()
	s := NewGetExecutionDetailService(ctx)
	// init req and assert value

	req := &exeution_manager.GetExecutionDetailRequest{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
