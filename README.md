# Ocean CLI

**巨量引擎命令行工具** —— 一行命令查广告数据，不用写代码。

专为广告投放人员、运营、优化师设计。告别反复登录后台导报表，在终端里直接查看消耗、点击、转化等核心数据。

## 它能做什么

```bash
# 查看最近 7 天账户整体数据
ocean report

# 查看指定日期范围的项目数据（按消耗排序 Top 20）
ocean report 2026-03-01 2026-03-21 campaign

# 查看单元维度数据
ocean report 2026-03-01 2026-03-21 ad

# 查看已授权的广告账户
ocean accounts
```

输出示例：

```
=== 项目报表 (按消耗降序 Top 20) ===
日期范围: 2026-03-01 ~ 2026-03-21

#1 示例项目-品牌推广-通投
  消耗(元)                    2000.00
  展示数                      20000
  点击数                      200
  点击率                      1.00%
  转化数                      10
  转化成本                     200.00
------------------------------------------------------------
```

## 快速开始

### 第一步：下载

前往 [Releases](../../releases) 页面，根据你的系统下载对应文件：

| 系统 | 文件 |
|------|------|
| Mac (Apple 芯片) | `ocean-darwin-arm64` |
| Mac (Intel 芯片) | `ocean-darwin-amd64` |
| Windows | `ocean-windows-amd64.exe` |
| Linux | `ocean-linux-amd64` |

下载后重命名为 `ocean`（Windows 为 `ocean.exe`），放到你方便找到的目录。

<details>
<summary>Mac 用户提示：首次运行可能提示"无法验证开发者"</summary>

在终端执行：
```bash
chmod +x ocean
xattr -d com.apple.quarantine ocean
```
</details>

### 第二步：配置

在 `ocean` 同目录下创建 `.env` 文件，填入你的应用信息：

```
APP_ID=你的应用ID
APP_SECRET=你的应用密钥
ADVERTISER_ID=你的广告主ID
```

> **如何获取这些信息？**
> 1. 登录 [巨量引擎开放平台](https://open.oceanengine.com)，创建应用，获取 APP_ID 和 APP_SECRET
> 2. ADVERTISER_ID 是你要查询的广告账户 ID，可以在巨量引擎后台的账户管理中找到

### 第三步：授权

首次使用需要完成一次 OAuth 授权：

```bash
./ocean auth
```

会弹出一个链接，在浏览器中打开并完成授权。授权成功后 token 会自动保存，后续无需重复操作（token 过期会自动刷新）。

### 第四步：开始使用

```bash
# 查看授权了哪些账户
./ocean accounts

# 查看最近 7 天的广告数据
./ocean report

# 查看本月数据，按项目维度
./ocean report 2026-03-01 2026-03-21 campaign
```

## 命令参考

```
ocean version                                 查看版本号
ocean auth                                    首次授权（只需一次）
ocean accounts                                查看已授权的广告账户
ocean report                                  最近 7 天广告主汇总数据
ocean report <开始日期> <结束日期>               指定日期范围
ocean report <开始日期> <结束日期> campaign      按项目维度（Top 20）
ocean report <开始日期> <结束日期> ad            按单元维度（Top 20）
ocean config                                  查看可用的报表指标和维度
```

日期格式：`2026-03-01`（年-月-日）

### 调试模式

任何命令加 `--debug` 可打印 SDK 请求/响应日志，用于排查 API 问题：

```bash
ocean --debug report
ocean report --debug 2026-03-01 2026-03-21
```

## 报表层级说明

巨量引擎的广告结构分为多层，本工具支持三个维度：

```
广告主 (advertiser)  ← 账户级汇总，默认
  └── 项目 (campaign)   ← 每个推广项目的数据
        └── 单元 (ad)       ← 最细粒度，每个投放策略的数据
```

## 常见问题

**Q: token 过期了怎么办？**
A: 工具会自动刷新 token，无需手动操作。如果刷新也失败，重新执行 `ocean auth` 即可。

**Q: 可以查多个广告账户吗？**
A: 修改 `.env` 中的 `ADVERTISER_ID` 切换账户。后续版本会支持多账户管理。

**Q: 数据和后台对不上？**
A: 本工具使用巨量引擎官方 API，数据与后台一致。注意检查日期范围是否相同。

## 作为 MCP Server 使用

除了命令行工具，本项目也可以作为 **MCP Server** 接入各类 AI 工具，让 AI 直接调用巨量引擎广告数据。

### 准备工作

先完成授权（与 CLI 工具共用同一份 token）：

```bash
./ocean auth
```

确保 `.env` 文件已配置好 `APP_ID`、`APP_SECRET`、`ADVERTISER_ID`。

---

### Claude Code

编辑 `~/.claude.json`，在 `mcpServers` 中添加：

```json
"oceanengine": {
  "type": "stdio",
  "command": "/bin/sh",
  "args": ["-c", "cd /absolute/path/to/oceanengine-mcp && ./oceanengine-mcp-server"],
  "env": {
    "APP_ID": "你的应用ID",
    "APP_SECRET": "你的应用密钥",
    "ADVERTISER_ID": "你的广告主ID"
  }
}
```

> 将 `/absolute/path/to/oceanengine-mcp` 替换为你的实际目录（用 `pwd` 查看），`oceanengine-mcp-server` 是从 [Releases](../../releases) 下载后重命名的可执行文件。

---

### OpenAI Codex

编辑 `~/.codex/config.toml`（全局）或项目目录下的 `.codex/config.toml`：

```toml
[mcp_servers.oceanengine]
command = "/bin/sh"
args = ["-c", "cd /absolute/path/to/oceanengine-mcp && ./oceanengine-mcp-server"]
```

---

### Gemini CLI

编辑 `~/.gemini/settings.json`（全局）或项目目录下的 `.gemini/settings.json`：

```json
{
  "mcpServers": {
    "oceanengine": {
      "command": "/bin/sh",
      "args": ["-c", "cd /absolute/path/to/oceanengine-mcp && ./oceanengine-mcp-server"]
    }
  }
}
```

## 构建

```bash
# 开发构建
go build -o ocean .

# 发布构建（注入版本号）
go build -ldflags "-X main.version=1.0.0" -o ocean .
```

## 技术信息

- 基于巨量引擎官方 SDK [`oceanengine/ad_open_sdk_go`](https://github.com/oceanengine/ad_open_sdk_go) 开发
- 使用 v3 自定义报表 API（`ReportCustomGetV30Api`）
- Token 自动持久化和刷新
- 单二进制文件，无需安装运行环境
- `--debug` 模式启用 SDK 原始日志

## License

MIT


---

## 关注我

<img src="./雷码工坊微信公众号.jpg" alt="雷码工坊笔记微信公众号" width="200" />

**雷码工坊笔记** — 微信扫码关注