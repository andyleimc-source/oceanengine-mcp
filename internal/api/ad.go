package api

import (
	"context"
	"fmt"
	"log"

	ad_open_sdk_go "github.com/oceanengine/ad_open_sdk_go"
	"github.com/oceanengine/ad_open_sdk_go/models"
)

// ListAds queries ad plans (v2).
// API: 10-001 AdGetV2Api
func ListAds(c *ad_open_sdk_go.Client, accessToken string, advID int64) {
	ctx := context.Background()
	fields := []string{
		"ad_id", "name", "ad_create_time", "ad_modify_time",
		"status", "campaign_id", "budget", "bid", "cpa_bid",
	}
	resp, _, err := c.AdGetV2Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(advID).
		Fields(fields).
		Page(1).
		PageSize(100).
		Execute()
	if err != nil {
		log.Fatalf("查询广告计划失败: %v", err)
	}
	checkResp("查询广告计划", resp.Code, resp.Message)

	if resp.Data == nil {
		fmt.Println("暂无广告计划数据")
		return
	}

	fmt.Println("=== 广告计划详情 ===")
	fmt.Println()
	printField("计划ID", resp.Data.AdId)
	printField("名称", resp.Data.Name)
	printField("广告组ID", resp.Data.CampaignId)
	printFieldFloat("出价(元)", resp.Data.Bid)
	printFieldFloat("CPA出价(元)", resp.Data.CpaBid)
	printField("状态", resp.Data.Status)
	fmt.Println()
}

// UpdateAdBid batch updates ad plan bids (v2).
// API: 10-002 AdUpdateBidV2Api
func UpdateAdBid(c *ad_open_sdk_go.Client, accessToken string, advID int64, adID int64, bid float64) {
	ctx := context.Background()
	req := models.AdUpdateBidV2Request{
		AdvertiserId: advID,
		Data: []*models.AdUpdateBidV2RequestDataInner{
			{AdId: adID, Bid: bid},
		},
	}
	resp, _, err := c.AdUpdateBidV2Api().
		Post(ctx).
		AccessToken(accessToken).
		AdUpdateBidV2Request(req).
		Execute()
	if err != nil {
		log.Fatalf("更新出价失败: %v", err)
	}
	checkResp("更新出价", resp.Code, resp.Message)
	fmt.Printf("广告计划 %d 出价已更新为 %.2f 元\n", adID, bid)
}

// UpdateAdBudget batch updates ad plan budgets (v2).
// API: 10-003 AdUpdateBudgetV2Api
func UpdateAdBudget(c *ad_open_sdk_go.Client, accessToken string, advID int64, adID int64, budget float64) {
	ctx := context.Background()
	req := models.AdUpdateBudgetV2Request{
		AdvertiserId: advID,
		Data: []*models.AdUpdateBudgetV2RequestDataInner{
			{AdId: adID, Budget: budget},
		},
	}
	resp, _, err := c.AdUpdateBudgetV2Api().
		Post(ctx).
		AccessToken(accessToken).
		AdUpdateBudgetV2Request(req).
		Execute()
	if err != nil {
		log.Fatalf("更新预算失败: %v", err)
	}
	checkResp("更新预算", resp.Code, resp.Message)
	fmt.Printf("广告计划 %d 预算已更新为 %.2f 元\n", adID, budget)
}

// GetAdRejectReason queries ad rejection reasons (v2).
// API: 10-005 AdRejectReasonV2Api
func GetAdRejectReason(c *ad_open_sdk_go.Client, accessToken string, advID int64, adIDs []int64) {
	ctx := context.Background()
	resp, _, err := c.AdRejectReasonV2Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(advID).
		AdIds(adIDs).
		Execute()
	if err != nil {
		log.Fatalf("查询审核拒绝原因失败: %v", err)
	}
	checkResp("查询审核拒绝原因", resp.Code, resp.Message)

	if resp.Data == nil {
		fmt.Println("暂无拒绝原因数据")
		return
	}
	fmt.Println("=== 广告计划审核拒绝原因 ===")
	fmt.Printf("%+v\n", resp.Data)
}

// GetAdCostProtectStatus queries cost protection status (v2).
// API: 10-006 AdCostProtectStatusGetV2Api
func GetAdCostProtectStatus(c *ad_open_sdk_go.Client, accessToken string, advID int64, adIDs []int64) {
	ctx := context.Background()
	resp, _, err := c.AdCostProtectStatusGetV2Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(advID).
		AdIds(adIDs).
		Execute()
	if err != nil {
		log.Fatalf("查询成本保护状态失败: %v", err)
	}
	checkResp("查询成本保护状态", resp.Code, resp.Message)

	if resp.Data == nil || len(resp.Data.List) == 0 {
		fmt.Println("暂无成本保护数据")
		return
	}
	fmt.Println("=== 成本保护状态 ===")
	for _, item := range resp.Data.List {
		fmt.Printf("%+v\n", item)
	}
}
