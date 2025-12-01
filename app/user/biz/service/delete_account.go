package service

import (
	"context"

	user "github.com/fuhanyang/CYGlowFlow/rpc_gen/kitex_gen/user"
)

type DeleteAccountService struct {
	ctx context.Context
} // NewDeleteAccountService new DeleteAccountService
func NewDeleteAccountService(ctx context.Context) *DeleteAccountService {
	return &DeleteAccountService{ctx: ctx}
}

// Run create note info
func (s *DeleteAccountService) Run(req *user.DeleteAccountReq) (resp *user.DeleteAccountResp, err error) {
	// Finish your business logic.

	return
}
