# 分支保护规则和GitHub Flow指南

## 分支策略

### 主要分支
- **main**: 生产环境代码，只接受来自release分支的合并
- **dev**: 开发主干分支，用于集成功能开发

### 辅助分支
- **feature/***: 功能开发分支
- **release/***: 发布准备分支
- **hotfix/***: 紧急修复分支

## 分支保护规则

### main分支保护
- ✅ 要求Pull Request才能合并
- ✅ 要求至少1个代码审查通过
- ✅ 要求CI状态检查通过
- ✅ 要求线性提交历史
- ✅ 限制推送权限给管理员

### dev分支保护
- ✅ 要求Pull Request才能合并
- ✅ 要求CI状态检查通过
- ✅ 限制直接推送权限

## GitHub Flow工作流程

### 1. 功能开发流程
```bash
# 从dev分支创建功能分支
git checkout dev
git pull origin dev
git checkout -b feature/your-feature-name

# 开发完成后提交
git add .
git commit -m "feat: 描述功能"

# 推送到远程并创建Pull Request
git push origin feature/your-feature-name
```

### 2. 代码审查
- 至少需要1名团队成员审查
- 所有CI检查必须通过
- 解决所有审查评论后才能合并

### 3. 发布流程
```bash
# 从dev创建release分支
git checkout dev
git checkout -b release/v1.0.0

# 进行最终测试和文档更新
# 合并到main分支
git checkout main
git merge release/v1.0.0
git tag v1.0.0
git push origin main --tags

# 同步到dev分支
git checkout dev
git merge release/v1.0.0
git branch -d release/v1.0.0
```

### 4. 紧急修复流程
```bash
# 从main分支创建hotfix分支
git checkout main
git checkout -b hotfix/issue-description

# 修复问题后合并
git checkout main
git merge hotfix/issue-description
git tag v1.0.1
git checkout dev
git merge hotfix/issue-description
```

## 提交信息规范

遵循Conventional Commits格式：
- `feat:` 新功能
- `fix:` 修复bug
- `docs:` 文档更新
- `style:` 代码格式调整
- `refactor:` 代码重构
- `test:` 测试相关
- `chore:` 构建过程或辅助工具变动

## CI/CD触发规则

- **push到feature分支**: 触发代码质量检查
- **push到develop分支**: 触发完整CI流程和开发环境部署
- **创建release标签**: 触发预发布环境部署
- **push到main分支**: 触发生产环境部署