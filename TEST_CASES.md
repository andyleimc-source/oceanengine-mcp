# TEST_CASES.md — 手工验收检查清单

本项目暂无自动化测试。以下是每次修改后应检查的手工验收步骤。

## 前置条件

- `.env` 已配置 `APP_ID`、`APP_SECRET`、`ADVERTISER_ID`
- `token.json` 存在且未过期（或可自动刷新）

## 编译检查

- [ ] `go build -o ocean .` 无报错
- [ ] `go vet ./...` 无警告

## 基础命令

| 命令 | 预期 | 检查 |
|------|------|------|
| `ocean` (无参数) | 打印帮助信息 | [ ] |
| `ocean version` | 输出 `ocean dev` 或注入的版本号 | [ ] |
| `ocean auth` | 启动本地服务器，打印授权链接 | [ ] |
| `ocean accounts` | 列出已授权账户 | [ ] |

## 报表

| 命令 | 预期 | 检查 |
|------|------|------|
| `ocean report` | 显示最近 7 天广告主汇总数据 | [ ] |
| `ocean report 2026-03-01 2026-03-21` | 显示指定日期范围数据 | [ ] |
| `ocean report 2026-03-01 2026-03-21 campaign` | 按项目维度 Top 20 | [ ] |
| `ocean report 2026-03-01 2026-03-21 ad` | 按单元维度 Top 20 | [ ] |
| `ocean config` | 显示可用指标和维度 | [ ] |

## 管理命令（v2 广告组）

| 命令 | 预期 | 检查 |
|------|------|------|
| `ocean campaigns` | 列出广告组，含 ID、名称、状态 | [ ] |
| `ocean campaign-status enable <id>` | 输出"状态已更新" | [ ] |

## 管理命令（v2 广告计划）

| 命令 | 预期 | 检查 |
|------|------|------|
| `ocean ads` | 显示广告计划信息（含 SDK 限制提示） | [ ] |
| `ocean ad-bid <id> <金额>` | 输出"出价已更新" | [ ] |
| `ocean ad-budget <id> <金额>` | 输出"预算已更新" | [ ] |
| `ocean ad-reject <id>` | 结构化显示拒绝原因（非 %+v） | [ ] |
| `ocean ad-cost-protect <id>` | 结构化显示保护状态 | [ ] |

## 管理命令（v3 项目）

| 命令 | 预期 | 检查 |
|------|------|------|
| `ocean projects` | 列出项目，含 ID、名称、状态 | [ ] |
| `ocean project-status enable <id>` | 输出"状态已更新" | [ ] |
| `ocean project-budget <id> <金额> day` | 输出"预算已更新" | [ ] |
| `ocean project-cost-protect <id>` | 结构化显示保护状态 | [ ] |

## 管理命令（v3 广告单元）

| 命令 | 预期 | 检查 |
|------|------|------|
| `ocean promotions` | 列出单元，含出价、预算等 | [ ] |
| `ocean promotion-status enable <id>` | 输出"状态已更新" | [ ] |
| `ocean promotion-bid <id> <金额>` | 输出"出价已更新" | [ ] |
| `ocean promotion-budget <id> <金额>` | 输出"预算已更新" | [ ] |
| `ocean promotion-reject <id>` | 结构化显示拒绝原因 | [ ] |
| `ocean promotion-cost-protect <id>` | 结构化显示保护状态 | [ ] |

## 创意

| 命令 | 预期 | 检查 |
|------|------|------|
| `ocean creatives` | 列出创意列表 | [ ] |
| `ocean creative-detail <ad_id>` | 显示创意详情 | [ ] |
| `ocean creative-reject <id>` | 显示拒绝原因 | [ ] |

## --debug 模式

| 命令 | 预期 | 检查 |
|------|------|------|
| `ocean --debug report` | 输出 SDK 请求/响应日志 + 正常报表 | [ ] |
| `ocean report --debug` | 同上（--debug 位置无关） | [ ] |

## 异常路径

| 场景 | 预期 | 检查 |
|------|------|------|
| 未配置 .env | 提示"请在 .env 中设置" | [ ] |
| APP_ID 非数字 | 提示"格式错误" | [ ] |
| 未授权（无 token.json） | 提示"run 'ocean auth' first" | [ ] |
| 无效 ID 格式 `ocean ad-bid abc 10` | 提示"格式错误" | [ ] |
| 逗号分隔含空值 `ocean ad-reject ,123,,` | 正常解析出 123 | [ ] |
| 未知命令 `ocean foo` | 打印帮助信息 | [ ] |
