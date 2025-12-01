# Docker 镜像仓库配置指南

## 推荐方案：GitHub Container Registry (ghcr.io)

### 为什么推荐使用 ghcr.io？
- 与 GitHub Actions 天然集成
- 免费额度充足（每月500MB存储，1GB传输）
- 版本管理与代码仓库同步
- 无需额外配置认证

### 配置方法
在 GitHub Secrets 中配置：
- **DOCKER_REGISTRY**: `ghcr.io/your-github-username`
- **DOCKER_USERNAME**: 你的 GitHub 用户名
- **DOCKER_PASSWORD**: 使用 Personal Access Token (PAT)

### 生成 Personal Access Token
1. 进入 GitHub Settings → Developer settings → Personal access tokens
2. 点击 "Generate new token"
3. 授予以下权限：
   - `repo` (完全控制)
   - `write:packages` (推送包)
   - `read:packages` (读取包)
4. 保存 token 作为 `DOCKER_PASSWORD`

## 备选方案：Docker Hub

### 配置方法
- **DOCKER_REGISTRY**: `docker.io/yourusername`
- **DOCKER_USERNAME**: 你的 Docker Hub 用户名
- **DOCKER_PASSWORD**: Docker Hub 密码或访问令牌

### 注意事项
- 免费版有拉取次数限制
- 私有镜像数量限制
- 企业版需要付费

## 企业方案：自建私有 Registry

### 使用 Docker 官方 Registry
```bash
# 启动私有 Registry
docker run -d -p 5000:5000 --name registry registry:2
```

### 配置方法
- **DOCKER_REGISTRY**: `your-domain.com:5000`
- **DOCKER_USERNAME**: 自定义用户名
- **DOCKER_PASSWORD**: 自定义密码

### 安全配置
```yaml
# docker-compose.yml 示例
version: '3'
services:
  registry:
    image: registry:2
    ports:
      - "5000:5000"
    environment:
      REGISTRY_AUTH: htpasswd
      REGISTRY_AUTH_HTPASSWD_REALM: Registry Realm
      REGISTRY_AUTH_HTPASSWD_PATH: /auth/htpasswd
    volumes:
      - ./auth:/auth
```

## 各方案对比

| 方案 | 免费额度 | 集成度 | 安全性 | 适用场景 |
|------|----------|--------|--------|----------|
| GitHub Container Registry | 500MB/月 | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐ | 开源项目、个人项目 |
| Docker Hub | 有限制 | ⭐⭐⭐⭐ | ⭐⭐⭐ | 个人开发、小型项目 |
| 自建 Registry | 无限 | ⭐⭐ | ⭐⭐⭐⭐⭐ | 企业级、生产环境 |

## 镜像命名规范

### GitHub Container Registry
```
ghcr.io/用户名/项目名-服务名:标签
示例：ghcr.io/fuhanyang/cyglowflow-gateway:dev-latest
```

### Docker Hub
```
用户名/项目名-服务名:标签
示例：fuhanyang/cyglowflow-gateway:dev-latest
```

### 私有 Registry
```
域名:端口/项目名-服务名:标签
示例：registry.example.com:5000/cyglowflow-gateway:dev-latest
```

## 生产环境最佳实践

1. **使用语义化版本标签**
   - `v1.0.0` - 稳定版本
   - `latest` - 最新稳定版
   - `dev-latest` - 开发版

2. **多架构支持**
   - 构建 amd64 和 arm64 镜像
   - 使用 Docker Buildx 多平台构建

3. **安全扫描**
   - 集成 Trivy 或 Snyk 进行漏洞扫描
   - 定期更新基础镜像

## 故障排除

### 认证失败
1. 检查 Secrets 配置是否正确
2. 确认 token 权限充足
3. 验证网络连接

### 推送失败
1. 检查镜像命名规范
2. 确认仓库存在且有写入权限
3. 验证磁盘空间

选择适合你项目需求的方案开始配置吧！