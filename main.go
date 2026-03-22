package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"

	"github.com/andyleimc-source/oceanengine-mcp/internal/api"
	"github.com/andyleimc-source/oceanengine-mcp/internal/auth"
	"github.com/andyleimc-source/oceanengine-mcp/internal/client"
)

// 通过 -ldflags 注入: go build -ldflags "-X main.version=1.0.0"
var version = "dev"

// Debug 模式，通过 --debug 标志启用
var debug = false

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// 解析全局 --debug 标志
	args := filterGlobalFlags(os.Args[1:])
	if debug {
		log.Println("[debug] 调试模式已启用")
	}

	_ = godotenv.Load()

	if len(args) < 1 {
		printUsage()
		os.Exit(1)
	}

	cmd := args[0]
	// 将过滤后的参数放回 os.Args 以保持后续解析兼容
	os.Args = append([]string{os.Args[0]}, args...)

	// version 不需要任何配置
	if cmd == "version" {
		fmt.Printf("ocean %s\n", version)
		return
	}

	appIDStr := os.Getenv("APP_ID")
	secret := os.Getenv("APP_SECRET")
	if appIDStr == "" || secret == "" {
		log.Fatal("请在 .env 中设置 APP_ID 和 APP_SECRET")
	}
	appID, err := strconv.ParseInt(appIDStr, 10, 64)
	if err != nil {
		log.Fatalf("APP_ID 格式错误: %v", err)
	}

	c := client.New()
	if debug {
		c.SetLogEnable(true)
	}

	// auth 不需要 token 和 advID
	if cmd == "auth" {
		auth.StartAuthServer(c, appID, secret)
		return
	}

	// 其余命令都需要 token
	token, err := client.GetValidToken(c, appID, secret)
	if err != nil {
		log.Fatal(err)
	}
	accessToken := token.AccessToken

	// accounts 只需要 token，不需要 advID
	if cmd == "accounts" {
		result, err := auth.ListAccounts(c, accessToken)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(result)
		return
	}

	// 其余所有命令都需要 advID
	advID := mustAdvertiserID()

	switch cmd {
	// === 报表 ===
	case "report":
		endDate := time.Now().Format("2006-01-02")
		startDate := time.Now().AddDate(0, 0, -7).Format("2006-01-02")
		if len(os.Args) >= 4 {
			startDate = os.Args[2]
			endDate = os.Args[3]
		}
		level := api.LevelAdvertiser
		if len(os.Args) >= 5 {
			switch os.Args[4] {
			case "ad":
				level = api.LevelPromotion
			case "campaign":
				level = api.LevelProject
			case "advertiser":
				level = api.LevelAdvertiser
			default:
				log.Fatalf("未知报表级别: %s (可选: advertiser, campaign, ad)", os.Args[4])
			}
		}
		result, err := api.FetchReport(api.ReportParams{
			Client:       c,
			AccessToken:  accessToken,
			AdvertiserID: advID,
			StartDate:    startDate,
			EndDate:      endDate,
			Level:        level,
		})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(result)

	case "config":
		result, err := api.FetchReportConfig(c, accessToken, advID)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(result)

	// === v2 广告组 Campaign ===
	case "campaigns":
		result, err := api.ListCampaigns(c, accessToken, advID)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(result)

	case "campaign-status":
		mustArgs(4, "ocean campaign-status <enable|disable|delete> <id1,id2,...>")
		result, err := api.UpdateCampaignStatus(c, accessToken, advID, mustParseIDs(os.Args[3]), os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(result)

	// === v2 广告计划 Ad ===
	case "ads":
		result, err := api.ListAds(c, accessToken, advID)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(result)

	case "ad-bid":
		mustArgs(4, "ocean ad-bid <ad_id> <出价金额>")
		result, err := api.UpdateAdBid(c, accessToken, advID, mustParseInt64(os.Args[2], "ad_id"), mustParseFloat64(os.Args[3], "出价"))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(result)

	case "ad-budget":
		mustArgs(4, "ocean ad-budget <ad_id> <预算金额>")
		result, err := api.UpdateAdBudget(c, accessToken, advID, mustParseInt64(os.Args[2], "ad_id"), mustParseFloat64(os.Args[3], "预算"))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(result)

	case "ad-reject":
		mustArgs(3, "ocean ad-reject <id1,id2,...>")
		result, err := api.GetAdRejectReason(c, accessToken, advID, mustParseIDs(os.Args[2]))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(result)

	case "ad-cost-protect":
		mustArgs(3, "ocean ad-cost-protect <id1,id2,...>")
		result, err := api.GetAdCostProtectStatus(c, accessToken, advID, mustParseIDs(os.Args[2]))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(result)

	// === v3 项目 Project ===
	case "projects":
		result, err := api.ListProjects(c, accessToken, advID)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(result)

	case "project-status":
		mustArgs(4, "ocean project-status <enable|disable> <id1,id2,...>")
		result, err := api.UpdateProjectStatus(c, accessToken, advID, mustParseIDs(os.Args[3]), os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(result)

	case "project-budget":
		mustArgs(5, "ocean project-budget <project_id> <预算金额> <day|total|infinite>")
		result, err := api.UpdateProjectBudget(c, accessToken, advID, mustParseInt64(os.Args[2], "project_id"), mustParseFloat64(os.Args[3], "预算"), os.Args[4])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(result)

	case "project-delete":
		mustArgs(3, "ocean project-delete <id1,id2,...>")
		result, err := api.DeleteProjects(c, accessToken, advID, mustParseIDs(os.Args[2]))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(result)

	case "project-cost-protect":
		mustArgs(3, "ocean project-cost-protect <id1,id2,...>")
		result, err := api.GetProjectCostProtectStatus(c, accessToken, advID, mustParseIDs(os.Args[2]))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(result)

	// === v3 广告单元 Promotion ===
	case "promotions":
		result, err := api.ListPromotions(c, accessToken, advID)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(result)

	case "promotion-status":
		mustArgs(4, "ocean promotion-status <enable|disable> <id1,id2,...>")
		result, err := api.UpdatePromotionStatus(c, accessToken, advID, mustParseIDs(os.Args[3]), os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(result)

	case "promotion-bid":
		mustArgs(4, "ocean promotion-bid <promotion_id> <出价金额>")
		result, err := api.UpdatePromotionBid(c, accessToken, advID, mustParseInt64(os.Args[2], "promotion_id"), mustParseFloat64(os.Args[3], "出价"))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(result)

	case "promotion-budget":
		mustArgs(4, "ocean promotion-budget <promotion_id> <预算金额>")
		result, err := api.UpdatePromotionBudget(c, accessToken, advID, mustParseInt64(os.Args[2], "promotion_id"), mustParseFloat64(os.Args[3], "预算"))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(result)

	case "promotion-delete":
		mustArgs(3, "ocean promotion-delete <id1,id2,...>")
		result, err := api.DeletePromotions(c, accessToken, advID, mustParseIDs(os.Args[2]))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(result)

	case "promotion-reject":
		mustArgs(3, "ocean promotion-reject <id1,id2,...>")
		result, err := api.GetPromotionRejectReason(c, accessToken, advID, mustParseIDs(os.Args[2]))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(result)

	case "promotion-cost-protect":
		mustArgs(3, "ocean promotion-cost-protect <id1,id2,...>")
		result, err := api.GetPromotionCostProtectStatus(c, accessToken, advID, mustParseIDs(os.Args[2]))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(result)

	// === 创意 Creative ===
	case "creatives":
		result, err := api.ListCreatives(c, accessToken, advID)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(result)

	case "creative-detail":
		mustArgs(3, "ocean creative-detail <ad_id>")
		result, err := api.GetCreativeDetail(c, accessToken, advID, mustParseInt64(os.Args[2], "ad_id"))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(result)

	case "creative-reject":
		mustArgs(3, "ocean creative-reject <id1,id2,...>")
		result, err := api.GetCreativeRejectReason(c, accessToken, advID, mustParseIDs(os.Args[2]))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(result)

	default:
		printUsage()
		os.Exit(1)
	}
}

// filterGlobalFlags 从参数中提取 --debug 等全局标志，返回过滤后的参数。
func filterGlobalFlags(args []string) []string {
	var filtered []string
	for _, arg := range args {
		switch arg {
		case "--debug":
			debug = true
		default:
			filtered = append(filtered, arg)
		}
	}
	return filtered
}

func mustArgs(n int, usage string) {
	if len(os.Args) < n {
		log.Fatalf("用法: %s", usage)
	}
}

func mustAdvertiserID() int64 {
	s := os.Getenv("ADVERTISER_ID")
	if s == "" {
		log.Fatal("请在 .env 中设置 ADVERTISER_ID")
	}
	id, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Fatalf("ADVERTISER_ID 格式错误: %v", err)
	}
	return id
}

func mustParseInt64(s string, name string) int64 {
	id, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Fatalf("%s 格式错误: %v", name, err)
	}
	return id
}

func mustParseFloat64(s string, name string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Fatalf("%s 格式错误: %v", name, err)
	}
	return f
}

func mustParseIDs(s string) []int64 {
	parts := strings.Split(s, ",")
	var ids []int64
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		id, err := strconv.ParseInt(p, 10, 64)
		if err != nil {
			log.Fatalf("ID 格式错误: %s", p)
		}
		ids = append(ids, id)
	}
	if len(ids) == 0 {
		log.Fatal("请提供至少一个 ID")
	}
	return ids
}

func printUsage() {
	fmt.Println(`ocean - 巨量引擎 Marketing API 工具

用法:
  ocean auth                                         OAuth2 授权
  ocean accounts                                     查看已授权账户
  ocean version                                      查看版本号

报表:
  ocean report                                       最近7天广告主报表
  ocean report 2026-03-14 2026-03-21                 指定日期范围
  ocean report 2026-03-14 2026-03-21 campaign        项目维度 (按消耗 Top 20)
  ocean report 2026-03-14 2026-03-21 ad              单元维度 (按消耗 Top 20)
  ocean config                                       查看可用报表指标和维度

v2 广告组 (Campaign):
  ocean campaigns                                    查询广告组列表
  ocean campaign-status <enable|disable|delete> <ids> 更新广告组状态

v2 广告计划 (Ad):
  ocean ads                                          查询广告计划列表
  ocean ad-bid <ad_id> <出价>                         更新出价
  ocean ad-budget <ad_id> <预算>                      更新预算
  ocean ad-reject <ids>                              查看审核拒绝原因
  ocean ad-cost-protect <ids>                        查看成本保护状态

v3 项目 (Project):
  ocean projects                                     查询项目列表
  ocean project-status <enable|disable> <ids>        更新项目状态
  ocean project-budget <id> <预算> <day|total|infinite> 更新项目预算
  ocean project-delete <ids>                         删除项目
  ocean project-cost-protect <id>                    查看成本保护状态

v3 广告单元 (Promotion):
  ocean promotions                                   查询单元列表
  ocean promotion-status <enable|disable> <ids>      更新单元状态
  ocean promotion-bid <id> <出价>                     更新出价
  ocean promotion-budget <id> <预算>                  更新预算
  ocean promotion-delete <ids>                       删除单元
  ocean promotion-reject <ids>                       查看审核拒绝原因
  ocean promotion-cost-protect <id>                  查看成本保护状态

创意 (Creative):
  ocean creatives                                    查询创意列表
  ocean creative-detail <ad_id>                      查看创意详情
  ocean creative-reject <ids>                        查看审核拒绝原因

全局选项:
  --debug                                            打印 SDK 请求/响应日志

注: <ids> 支持逗号分隔多个ID, 如 123,456,789`)
}
