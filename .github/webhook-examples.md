# GitHub Webhook 配置示例

## 常用Webhook事件触发器

### 1. 构建状态通知
```yaml
事件类型: 
- push
- pull_request
- workflow_run

通知目标:
- Slack频道
- 飞书机器人
- 邮件列表
```

### 2. 部署状态跟踪
```yaml
事件类型:
- deployment_status
- release

用途:
- 通知部署成功/失败
- 更新部署面板
- 记录部署历史
```

### 3. 代码质量监控
```yaml
事件类型:
- pull_request_review
- status

用途:
- 代码审查通知
- 测试覆盖率报告
- 安全扫描结果
```

## 配置示例

### Slack集成示例
```bash
Webhook URL: https://hooks.slack.com/services/...
事件: push, pull_request, workflow_run
```

### 飞书机器人示例
```bash
Webhook URL: https://open.feishu.cn/open-apis/bot/v2/hook/...
签名算法: hmac-sha256
```

### 钉钉机器人示例
```bash
Webhook URL: https://oapi.dingtalk.com/robot/send?access_token=...
签名算法: 自定义签名
```

## 安全配置

### Webhook Secret配置
```bash
# 生成随机密钥
openssl rand -base64 32
```

### 验证签名示例
```go
// Go语言验证Webhook签名示例
func verifySignature(secret, signature string, body []byte) bool {
    mac := hmac.New(sha256.New, []byte(secret))
    mac.Write(body)
    expectedMAC := hex.EncodeToString(mac.Sum(nil))
    return hmac.Equal([]byte(signature), []byte(expectedMAC))
}
```

## 最佳实践

1. **使用HTTPS** - 确保Webhook URL使用加密传输
2. **设置Secret** - 验证请求来源，防止伪造
3. **处理重试** - 实现幂等性，处理重复请求
4. **超时设置** - 设置合理的超时时间（建议5-10秒）
5. **错误处理** - 记录失败原因，提供重试机制
6. **Payload验证** - 验证收到的数据格式和内容

## 在GitHub中配置Webhook的步骤

1. 进入仓库设置 → Webhooks
2. 点击"Add webhook"
3. 填写Payload URL（你的接收端点）
4. 设置Content type为application/json
5. 设置Secret（安全验证）
6. 选择触发的事件类型
7. 保存配置并测试

## 实际应用场景

### 团队协作通知
- PR创建和合并通知
- 代码审查提醒
- 构建失败告警

### 运维监控
- 部署状态实时通知
- 服务健康状态监控
- 异常告警推送

### 业务流程集成
- 自动创建工单
- 数据同步更新
- 报表生成触发
```

是否需要我为你创建具体的Webhook配置示例，或者你希望集成特定的第三方服务（如Slack、飞书等）？