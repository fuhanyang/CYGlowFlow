# GitHub Flow 配置检查清单

## ✅ 已完成
- [x] CI/CD 工作流配置文件 (.github/workflows/ci.yml)
- [x] 多环境部署配置 (.github/workflows/deploy.yml) 
- [x] 分支保护规则文档 (.github/BRANCH_PROTECTION.md)
- [x] 服务Dockerfile配置 (app/gateway/Dockerfile, app/user/Dockerfile)

## ⚙️ 需要在GitHub仓库设置的配置

### 1. 启用Actions功能
**路径**: Settings → Actions → General
- [ ] Allow all actions and reusable workflows
- [ ] ✅ Workflow permissions: Read and write permissions

### 2. 配置环境Secrets
**路径**: Settings → Secrets and variables → Actions
- [ ] `DOCKER_REGISTRY` - 你的Docker仓库地址
- [ ] `DOCKER_USERNAME` - Docker仓库用户名
- [ ] `DOCKER_PASSWORD` - Docker仓库密码

### 3. 设置分支保护规则
**路径**: Settings → Branches → Branch protection rules
- [ ] **main分支保护规则**:
  - Require a pull request before merging
  - Require approvals (1 reviewer)
  - Require status checks to pass
  - Require linear history
- [ ] **develop分支保护规则**:
  - Require a pull request before merging
  - Require status checks to pass

### 4. 配置Webhooks (可选)
**路径**: Settings → Webhooks
- [ ] 添加Slack/Teams通知Webhook

## 🧪 测试GitHub Flow

### 第一次测试流程
1. 提交代码到 `develop` 分支
```bash
git checkout develop
git add .
git commit -m "chore: setup github flow"
git push origin develop
```

2. 观察GitHub Actions运行情况
   - 检查CI流水线是否成功
   - 验证代码质量检查
   - 确认测试覆盖率报告

3. 创建功能分支测试
```bash
git checkout -b feature/test-github-flow
# 做一些小修改
git add .
git commit -m "feat: test github flow"
git push origin feature/test-github-flow
```

4. 创建Pull Request测试代码审查流程

## 🔧 故障排除

### 常见问题
1. **Actions不触发**: 检查分支保护设置和Actions权限
2. **Docker构建失败**: 验证Secrets配置是否正确
3. **测试失败**: 检查Go模块依赖和测试环境

### 验证命令
```bash
# 本地验证Go项目
go mod tidy
go test ./...
go build ./app/gateway
go build ./app/user
```

## 📞 支持
如遇到配置问题，参考：
- [GitHub Actions文档](https://docs.github.com/en/actions)
- [Docker官方文档](https://docs.docker.com/)
- 项目中的 `.github/BRANCH_PROTECTION.md`