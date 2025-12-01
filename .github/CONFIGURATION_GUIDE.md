# GitHub Flow 配置步骤指南

## 第一步：在GitHub仓库中启用Actions

1. 进入你的GitHub仓库页面
2. 点击右上角的 **Settings** 标签
3. 在左侧菜单中找到 **Actions** → **General**
4. 配置Actions权限：
   - ✅ **Allow all actions and reusable workflows**
   - ✅ **Workflow permissions**: Read and write permissions

## 第二步：配置环境Secrets

1. 在仓库设置中进入 **Secrets and variables** → **Actions**
2. 点击 **New repository secret** 添加以下必需的密钥：

### 必需的环境变量（已为你配置为Docker Hub）
- **`DOCKER_USERNAME`** - 你的Docker Hub用户名 (liangjiaqi873)
- **`DOCKER_PASSWORD`** - Docker Hub密码或访问令牌
- **注意**：DOCKER_REGISTRY 已硬编码为你的仓库地址 `liangjiaqi873`

### 可选的配置
- **`SLACK_WEBHOOK_URL`** - Slack通知Webhook URL
- **`CODECOV_TOKEN`** - 覆盖率报告Token

## 第三步：设置分支保护规则

1. 进入 **Settings** → **Branches** → **Branch protection rules**
2. 为 `main` 分支设置规则：
   - ✅ Require a pull request before merging
   - ✅ Require approvals (至少1个审查)
   - ✅ Require status checks to pass
   - ✅ Require linear history
3. 为 `dev` 分支设置规则：
   - ✅ Require a pull request before merging  
   - ✅ Require status checks to pass

## 第四步：配置Webhook（可选）

1. 进入 **Settings** → **Webhooks**
2. 点击 **Add webhook**
3. 配置参数：
   - **Payload URL**: 你的Webhook接收端点
   - **Content type**: application/json
   - **Secret**: 使用 `openssl rand -base64 32` 生成的密钥
   - **Which events**: 选择触发事件（push, pull_request等）

## 第五步：验证配置

### 测试CI/CD流水线
```bash
# 提交代码到dev分支触发测试
git checkout dev
git add .
git commit -m "test: trigger CI pipeline"
git push origin dev
```

### 检查Actions运行状态
1. 进入仓库的 **Actions** 标签页
2. 查看CI Pipeline运行状态
3. 确保所有检查通过

## 常见问题解决

### Actions不触发
- 检查分支保护设置
- 确认Actions权限配置正确
- 验证工作流文件语法

### Docker构建失败
- 确认Docker Registry认证信息正确
- 检查Dockerfile配置
- 验证网络连接

### 测试失败
- 查看测试日志排查问题
- 检查Go模块依赖
- 验证测试环境配置

## 首次配置检查清单

- [ ] Actions功能已启用
- [ ] 环境Secrets已配置
- [ ] 分支保护规则已设置
- [ ] 初始测试提交成功
- [ ] CI流水线正常运行

完成以上配置后，你的GitHub Flow工作流就可以正常使用了！