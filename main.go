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

func printUsage() {
	fmt.Println(`ocean - 巨量引擎 Marketing API 工具

用法:
  ocean auth                                    OAuth2 授权
  ocean accounts                                查看已授权账户
  ocean report                                  最近7天广告主报表
  ocean report 2026-03-14 2026-03-21            指定日期范围
  ocean report 2026-03-14 2026-03-21 ad         单元维度 (按消耗 Top 20)
  ocean report 2026-03-14 2026-03-21 campaign   项目维度 (按消耗 Top 20)
  ocean config                                  查看可用报表指标和维度

报表级别:
  advertiser  广告主汇总 (默认)
  campaign    项目维度
  ad          单元维度`)
}
