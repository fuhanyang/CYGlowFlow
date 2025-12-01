package service

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/fuhanyang/CYGlowFlow/app/user/biz/dal/mysql"
	"github.com/fuhanyang/CYGlowFlow/app/user/biz/dal/mysql/model"
	"github.com/fuhanyang/CYGlowFlow/app/user/biz/dal/mysql/query"
	user "github.com/fuhanyang/CYGlowFlow/rpc_gen/kitex_gen/user"
)

type RegisterService struct {
	ctx context.Context
} // NewRegisterService new RegisterService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// Run create note info
func (s *RegisterService) Run(req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	// 1. 检查参数
	if req.Username == "" || req.Password == "" {
		return nil, errors.New("username or password cannot be empty")
	}

	// 2. 检查用户是否已存在 (通过邮箱或手机号，或者后续支持用户名唯一)
	// 目前 user 表设计有 username/account_num, email, phone
	q := query.Use(mysql.DB)
	u := q.User

	// 简单的检查，实际业务可能需要更复杂的去重逻辑
	count, err := u.WithContext(s.ctx).Where(u.Email.Eq(req.Email)).Or(u.Phone.Eq(req.Phone)).Count()
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, errors.New("user already exists")
	}

	// 3. 生成账号 (数字)
	// 简单生成一个随机数作为账号，实际生产环境需要发号器
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	accountNum := int64(r.Intn(90000000) + 10000000) // 8位随机数

	// 确保账号唯一
	for {
		_, err := u.WithContext(s.ctx).Where(u.AccountNum.Eq(accountNum)).First()
		if err != nil {
			// 没找到，说明可用
			break
		}
		// 冲突了，重新生成
		accountNum = int64(r.Intn(90000000) + 10000000)
	}

	// 4. 密码加密 (MD5 简单示例，生产建议使用 bcrypt/argon2)
	hasher := md5.New()
	hasher.Write([]byte(req.Password))
	encryptedPassword := hex.EncodeToString(hasher.Sum(nil))

	// 5. 创建用户
	newUser := &model.User{
		AccountNum: accountNum,
		Name:       &req.Username,
		Password:   encryptedPassword,
		Email:      &req.Email,
		Phone:      &req.Phone,
	}

	err = u.WithContext(s.ctx).Create(newUser)
	if err != nil {
		return nil, err
	}

	// 6. 返回结果
	resp = &user.RegisterResp{
		UserId:     int64(newUser.ID),
		AccountNum: accountNum,
	}
	fmt.Printf("resp: %v\n", resp)
	return resp, nil
}
