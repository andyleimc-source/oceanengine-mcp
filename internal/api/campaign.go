package api

import (
	"context"
	"fmt"
	"log"

	ad_open_sdk_go "github.com/oceanengine/ad_open_sdk_go"
	"github.com/oceanengine/ad_open_sdk_go/models"
)

// ListCampaigns queries ad groups (v2).
// API: 09-002 CampaignGetV2Api
func ListCampaigns(c *ad_open_sdk_go.Client, accessToken string, advID int64) {
	ctx := context.Background()
	fields := []string{
		"id", "name", "budget_mode", "budget",
		"landing_type", "status", "campaign_create_time",
	}
	resp, _, err := c.CampaignGetV2Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(advID).
		Fields(fields).
		Page(1).
		PageSize(100).
		Execute()
	if err != nil {
		log.Fatalf("查询广告组失败: %v", err)
	}
	checkResp("查询广告组", resp.Code, resp.Message)

	if resp.Data == nil || len(resp.Data.List) == 0 {
		fmt.Println("暂无广告组数据")
		return
	}

	fmt.Println("=== 广告组列表 ===")
	fmt.Println()
	for i, item := range resp.Data.List {
		fmt.Printf("#%d\n", i+1)
		printField("广告组ID", item.Id)
		printField("名称", item.Name)
		printField("预算模式", item.BudgetMode)
		printFieldFloat("预算(元)", item.Budget)
		printField("推广目的", item.LandingType)
		printField("状态", item.Status)
		printField("创建时间", item.CampaignCreateTime)
		fmt.Println("  ----------")
	}
	printPageInfo(resp.Data.PageInfo)
}

// UpdateCampaignStatus updates ad group status (v2).
// API: 09-004 CampaignUpdateStatusV2Api
func UpdateCampaignStatus(c *ad_open_sdk_go.Client, accessToken string, advID int64, campaignIDs []int64, optStatus string) {
	ctx := context.Background()

	var status models.CampaignUpdateStatusV2OptStatus
	switch optStatus {
	case "enable":
		status = models.ENABLE_CampaignUpdateStatusV2OptStatus
	case "disable":
		status = models.DISABLE_CampaignUpdateStatusV2OptStatus
	case "delete":
		status = models.DELETE_CampaignUpdateStatusV2OptStatus
	default:
		log.Fatalf("未知状态: %s (可选: enable, disable, delete)", optStatus)
	}

	req := models.CampaignUpdateStatusV2Request{
		AdvertiserId: advID,
		CampaignIds:  campaignIDs,
		OptStatus:    status,
	}
	resp, _, err := c.CampaignUpdateStatusV2Api().
		Post(ctx).
		AccessToken(accessToken).
		CampaignUpdateStatusV2Request(req).
		Execute()
	if err != nil {
		log.Fatalf("更新广告组状态失败: %v", err)
	}
	checkResp("更新广告组状态", resp.Code, resp.Message)
	fmt.Printf("广告组状态已更新为: %s\n", optStatus)
}
