package service

import (
	"context"
	"testing"
	user "github.com/fuhanyang/CYGlowFlow/rpc_gen/kitex_gen/user"
)

func TestDeleteAccount_Run(t *testing.T) {
	ctx := context.Background()
	s := NewDeleteAccountService(ctx)
	// init req and assert value

	req := &user.DeleteAccountReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
