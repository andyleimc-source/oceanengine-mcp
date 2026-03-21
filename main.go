package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	ad_open_sdk_go "github.com/oceanengine/ad_open_sdk_go"
	"github.com/joho/godotenv"

	"ocean/internal/api"
	"ocean/internal/auth"
	"ocean/internal/client"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	_ = godotenv.Load()

	appIDStr := os.Getenv("APP_ID")
	secret := os.Getenv("APP_SECRET")
	if appIDStr == "" || secret == "" {
		log.Fatal("请在 .env 中设置 APP_ID 和 APP_SECRET")
	}
	appID, err := strconv.ParseInt(appIDStr, 10, 64)
	if err != nil {
		log.Fatalf("APP_ID 格式错误: %v", err)
	}

	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	c := client.New()

	switch os.Args[1] {
	case "auth":
		auth.StartAuthServer(c, appID, secret)

	case "accounts":
		token := mustToken(c, appID, secret)
		auth.ListAccounts(c, token.AccessToken)

	case "report":
		advID := mustAdvertiserID()
		token := mustToken(c, appID, secret)

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

		api.FetchReport(api.ReportParams{
			Client:       c,
			AccessToken:  token.AccessToken,
			AdvertiserID: advID,
			StartDate:    startDate,
			EndDate:      endDate,
			Level:        level,
		})

	case "config":
		advID := mustAdvertiserID()
		token := mustToken(c, appID, secret)
		api.FetchReportConfig(c, token.AccessToken, advID)

	// === v2 广告组 Campaign ===
	case "campaigns":
		advID := mustAdvertiserID()
		token := mustToken(c, appID, secret)
		api.ListCampaigns(c, token.AccessToken, advID)

	case "campaign-status":
		// ocean campaign-status enable 123,456
		advID := mustAdvertiserID()
		token := mustToken(c, appID, secret)
		if len(os.Args) < 4 {
			log.Fatal("用法: ocean campaign-status <enable|disable|delete> <id1,id2,...>")
		}
		ids := mustParseIDs(os.Args[3])
		api.UpdateCampaignStatus(c, token.AccessToken, advID, ids, os.Args[2])

	// === v2 广告计划 Ad ===
	case "ads":
		advID := mustAdvertiserID()
		token := mustToken(c, appID, secret)
		api.ListAds(c, token.AccessToken, advID)

	case "ad-bid":
		// ocean ad-bid <ad_id> <bid>
		advID := mustAdvertiserID()
		token := mustToken(c, appID, secret)
		if len(os.Args) < 4 {
			log.Fatal("用法: ocean ad-bid <ad_id> <出价金额>")
		}
		adID := mustParseInt64(os.Args[2], "ad_id")
		bid := mustParseFloat64(os.Args[3], "出价")
		api.UpdateAdBid(c, token.AccessToken, advID, adID, bid)

	case "ad-budget":
		// ocean ad-budget <ad_id> <budget>
		advID := mustAdvertiserID()
		token := mustToken(c, appID, secret)
		if len(os.Args) < 4 {
			log.Fatal("用法: ocean ad-budget <ad_id> <预算金额>")
		}
		adID := mustParseInt64(os.Args[2], "ad_id")
		budget := mustParseFloat64(os.Args[3], "预算")
		api.UpdateAdBudget(c, token.AccessToken, advID, adID, budget)

	case "ad-reject":
		// ocean ad-reject <id1,id2,...>
		advID := mustAdvertiserID()
		token := mustToken(c, appID, secret)
		if len(os.Args) < 3 {
			log.Fatal("用法: ocean ad-reject <id1,id2,...>")
		}
		ids := mustParseIDs(os.Args[2])
		api.GetAdRejectReason(c, token.AccessToken, advID, ids)

	case "ad-cost-protect":
		// ocean ad-cost-protect <id1,id2,...>
		advID := mustAdvertiserID()
		token := mustToken(c, appID, secret)
		if len(os.Args) < 3 {
			log.Fatal("用法: ocean ad-cost-protect <id1,id2,...>")
		}
		ids := mustParseIDs(os.Args[2])
		api.GetAdCostProtectStatus(c, token.AccessToken, advID, ids)

	// === v3 项目 Project ===
	case "projects":
		advID := mustAdvertiserID()
		token := mustToken(c, appID, secret)
		api.ListProjects(c, token.AccessToken, advID)

	case "project-status":
		// ocean project-status <enable|disable> <id1,id2,...>
		advID := mustAdvertiserID()
		token := mustToken(c, appID, secret)
		if len(os.Args) < 4 {
			log.Fatal("用法: ocean project-status <enable|disable> <id1,id2,...>")
		}
		ids := mustParseIDs(os.Args[3])
		api.UpdateProjectStatus(c, token.AccessToken, advID, ids, os.Args[2])

	case "project-budget":
		// ocean project-budget <project_id> <budget> <day|total|infinite>
		advID := mustAdvertiserID()
		token := mustToken(c, appID, secret)
		if len(os.Args) < 5 {
			log.Fatal("用法: ocean project-budget <project_id> <预算金额> <day|total|infinite>")
		}
		projectID := mustParseInt64(os.Args[2], "project_id")
		budget := mustParseFloat64(os.Args[3], "预算")
		api.UpdateProjectBudget(c, token.AccessToken, advID, projectID, budget, os.Args[4])

	case "project-delete":
		// ocean project-delete <id1,id2,...>
		advID := mustAdvertiserID()
		token := mustToken(c, appID, secret)
		if len(os.Args) < 3 {
			log.Fatal("用法: ocean project-delete <id1,id2,...>")
		}
		ids := mustParseIDs(os.Args[2])
		api.DeleteProjects(c, token.AccessToken, advID, ids)

	case "project-cost-protect":
		// ocean project-cost-protect <id1,id2,...>
		advID := mustAdvertiserID()
		token := mustToken(c, appID, secret)
		if len(os.Args) < 3 {
			log.Fatal("用法: ocean project-cost-protect <id1,id2,...>")
		}
		ids := mustParseIDs(os.Args[2])
		api.GetProjectCostProtectStatus(c, token.AccessToken, advID, ids)

	// === v3 广告单元 Promotion ===
	case "promotions":
		advID := mustAdvertiserID()
		token := mustToken(c, appID, secret)
		api.ListPromotions(c, token.AccessToken, advID)

	case "promotion-status":
		// ocean promotion-status <enable|disable> <id1,id2,...>
		advID := mustAdvertiserID()
		token := mustToken(c, appID, secret)
		if len(os.Args) < 4 {
			log.Fatal("用法: ocean promotion-status <enable|disable> <id1,id2,...>")
		}
		ids := mustParseIDs(os.Args[3])
		api.UpdatePromotionStatus(c, token.AccessToken, advID, ids, os.Args[2])

	case "promotion-bid":
		// ocean promotion-bid <promotion_id> <bid>
		advID := mustAdvertiserID()
		token := mustToken(c, appID, secret)
		if len(os.Args) < 4 {
			log.Fatal("用法: ocean promotion-bid <promotion_id> <出价金额>")
		}
		promotionID := mustParseInt64(os.Args[2], "promotion_id")
		bid := mustParseFloat64(os.Args[3], "出价")
		api.UpdatePromotionBid(c, token.AccessToken, advID, promotionID, bid)

	case "promotion-budget":
		// ocean promotion-budget <promotion_id> <budget>
		advID := mustAdvertiserID()
		token := mustToken(c, appID, secret)
		if len(os.Args) < 4 {
			log.Fatal("用法: ocean promotion-budget <promotion_id> <预算金额>")
		}
		promotionID := mustParseInt64(os.Args[2], "promotion_id")
		budget := mustParseFloat64(os.Args[3], "预算")
		api.UpdatePromotionBudget(c, token.AccessToken, advID, promotionID, budget)

	case "promotion-delete":
		// ocean promotion-delete <id1,id2,...>
		advID := mustAdvertiserID()
		token := mustToken(c, appID, secret)
		if len(os.Args) < 3 {
			log.Fatal("用法: ocean promotion-delete <id1,id2,...>")
		}
		ids := mustParseIDs(os.Args[2])
		api.DeletePromotions(c, token.AccessToken, advID, ids)

	case "promotion-reject":
		// ocean promotion-reject <id1,id2,...>
		advID := mustAdvertiserID()
		token := mustToken(c, appID, secret)
		if len(os.Args) < 3 {
			log.Fatal("用法: ocean promotion-reject <id1,id2,...>")
		}
		ids := mustParseIDs(os.Args[2])
		api.GetPromotionRejectReason(c, token.AccessToken, advID, ids)

	case "promotion-cost-protect":
		// ocean promotion-cost-protect <id1,id2,...>
		advID := mustAdvertiserID()
		token := mustToken(c, appID, secret)
		if len(os.Args) < 3 {
			log.Fatal("用法: ocean promotion-cost-protect <id1,id2,...>")
		}
		ids := mustParseIDs(os.Args[2])
		api.GetPromotionCostProtectStatus(c, token.AccessToken, advID, ids)

	// === 创意 Creative ===
	case "creatives":
		advID := mustAdvertiserID()
		token := mustToken(c, appID, secret)
		api.ListCreatives(c, token.AccessToken, advID)

	case "creative-detail":
		// ocean creative-detail <ad_id>
		advID := mustAdvertiserID()
		token := mustToken(c, appID, secret)
		if len(os.Args) < 3 {
			log.Fatal("用法: ocean creative-detail <ad_id>")
		}
		adID := mustParseInt64(os.Args[2], "ad_id")
		api.GetCreativeDetail(c, token.AccessToken, advID, adID)

	case "creative-reject":
		// ocean creative-reject <id1,id2,...>
		advID := mustAdvertiserID()
		token := mustToken(c, appID, secret)
		if len(os.Args) < 3 {
			log.Fatal("用法: ocean creative-reject <id1,id2,...>")
		}
		ids := mustParseIDs(os.Args[2])
		api.GetCreativeRejectReason(c, token.AccessToken, advID, ids)

	default:
		printUsage()
		os.Exit(1)
	}
}

func mustToken(c *ad_open_sdk_go.Client, appID int64, secret string) *client.TokenStore {
	token, err := client.GetValidToken(c, appID, secret)
	if err != nil {
		log.Fatal(err)
	}
	return token
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
	parts := splitComma(s)
	ids := make([]int64, len(parts))
	for i, p := range parts {
		id, err := strconv.ParseInt(p, 10, 64)
		if err != nil {
			log.Fatalf("ID 格式错误: %s", p)
		}
		ids[i] = id
	}
	return ids
}

func splitComma(s string) []string {
	var result []string
	current := ""
	for _, ch := range s {
		if ch == ',' {
			if current != "" {
				result = append(result, current)
			}
			current = ""
		} else {
			current += string(ch)
		}
	}
	if current != "" {
		result = append(result, current)
	}
	return result
}

func printUsage() {
	fmt.Println(`ocean - 巨量引擎 Marketing API 工具

用法:
  ocean auth                                         OAuth2 授权
  ocean accounts                                     查看已授权账户

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

注: <ids> 支持逗号分隔多个ID, 如 123,456,789`)
}
