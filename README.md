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
ocean auth                                    首次授权（只需一次）
ocean accounts                                查看已授权的广告账户
ocean report                                  最近 7 天广告主汇总数据
ocean report <开始日期> <结束日期>               指定日期范围
ocean report <开始日期> <结束日期> campaign      按项目维度（Top 20）
ocean report <开始日期> <结束日期> ad            按单元维度（Top 20）
ocean config                                  查看可用的报表指标和维度
```

日期格式：`2026-03-01`（年-月-日）

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

## 技术信息

- 基于巨量引擎官方 SDK [`oceanengine/ad_open_sdk_go`](https://github.com/oceanengine/ad_open_sdk_go) 开发
- 使用 v3 自定义报表 API（`ReportCustomGetV30Api`）
- Token 自动持久化和刷新
- 单二进制文件，无需安装运行环境

## License

MIT
