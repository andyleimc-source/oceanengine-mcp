package api

import (
	"context"
	"fmt"
	"log"

	ad_open_sdk_go "github.com/oceanengine/ad_open_sdk_go"
	"github.com/oceanengine/ad_open_sdk_go/models"
)

// ListPromotions queries promotions/ad units (v3).
// API: 12-002 PromotionListV30Api
func ListPromotions(c *ad_open_sdk_go.Client, accessToken string, advID int64) {
	ctx := context.Background()
	fields := []string{
		"promotion_id", "promotion_name", "status", "status_first",
		"project_id", "budget_mode", "budget", "bid", "cpa_bid",
		"deep_cpabid", "roi_goal", "learning_phase",
		"promotion_create_time", "promotion_modify_time",
	}
	resp, _, err := c.PromotionListV30Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(advID).
		Fields(fields).
		Page(1).
		PageSize(20).
		Execute()
	if err != nil {
		log.Fatalf("查询广告单元列表失败: %v", err)
	}
	checkResp("查询广告单元列表", resp.Code, resp.Message)

	if resp.Data == nil || len(resp.Data.List) == 0 {
		fmt.Println("暂无广告单元数据")
		return
	}

	fmt.Println("=== 广告单元列表 (v3) ===")
	fmt.Println()
	for i, item := range resp.Data.List {
		fmt.Printf("#%d\n", i+1)
		printField("单元ID", item.PromotionId)
		printField("名称", item.PromotionName)
		printField("所属项目ID", item.ProjectId)
		printField("状态", item.Status)
		printField("一级状态", item.StatusFirst)
		printField("预算模式", item.BudgetMode)
		printFieldFloat("预算(元)", item.Budget)
		printFieldFloat("出价(元)", item.Bid)
		printFieldFloat("CPA出价(元)", item.CpaBid)
		printFieldFloat("深度出价(元)", item.DeepCpabid)
		printFieldFloat("ROI目标", item.RoiGoal)
		printField("学习期状态", item.LearningPhase)
		printField("创建时间", item.PromotionCreateTime)
		fmt.Println("  ----------")
	}
	printPageInfo(resp.Data.PageInfo)
}

// UpdatePromotionStatus updates promotion status (v3).
// API: 12-005 PromotionStatusUpdateV30Api
func UpdatePromotionStatus(c *ad_open_sdk_go.Client, accessToken string, advID int64, promotionIDs []int64, optStatus string) {
	ctx := context.Background()

	var status models.PromotionStatusUpdateV30DataOptStatus
	switch optStatus {
	case "enable":
		status = models.ENABLE_PromotionStatusUpdateV30DataOptStatus
	case "disable":
		status = models.DISABLE_PromotionStatusUpdateV30DataOptStatus
	default:
		log.Fatalf("未知状态: %s (可选: enable, disable)", optStatus)
	}

	data := make([]*models.PromotionStatusUpdateV30RequestDataInner, len(promotionIDs))
	for i, id := range promotionIDs {
		data[i] = &models.PromotionStatusUpdateV30RequestDataInner{
			PromotionId: id,
			OptStatus:   status,
		}
	}

	req := models.PromotionStatusUpdateV30Request{
		AdvertiserId: advID,
		Data:         data,
	}
	resp, _, err := c.PromotionStatusUpdateV30Api().
		Post(ctx).
		AccessToken(accessToken).
		PromotionStatusUpdateV30Request(req).
		Execute()
	if err != nil {
		log.Fatalf("更新单元状态失败: %v", err)
	}
	checkResp("更新单元状态", resp.Code, resp.Message)
	fmt.Printf("广告单元状态已更新为: %s\n", optStatus)
}

// UpdatePromotionBid batch updates promotion bids (v3).
// API: 12-006 PromotionBidUpdateV30Api
func UpdatePromotionBid(c *ad_open_sdk_go.Client, accessToken string, advID int64, promotionID int64, bid float64) {
	ctx := context.Background()
	req := models.PromotionBidUpdateV30Request{
		AdvertiserId: advID,
		Data: []*models.PromotionBidUpdateV30RequestDataInner{
			{PromotionId: promotionID, Bid: bid},
		},
	}
	resp, _, err := c.PromotionBidUpdateV30Api().
		Post(ctx).
		AccessToken(accessToken).
		PromotionBidUpdateV30Request(req).
		Execute()
	if err != nil {
		log.Fatalf("更新单元出价失败: %v", err)
	}
	checkResp("更新单元出价", resp.Code, resp.Message)
	fmt.Printf("广告单元 %d 出价已更新为 %.2f 元\n", promotionID, bid)
}

// UpdatePromotionBudget updates promotion budget (v3).
// API: 12-007 PromotionBudgetUpdateV30Api
func UpdatePromotionBudget(c *ad_open_sdk_go.Client, accessToken string, advID int64, promotionID int64, budget float64) {
	ctx := context.Background()
	req := models.PromotionBudgetUpdateV30Request{
		AdvertiserId: advID,
		Data: []*models.PromotionBudgetUpdateV30RequestDataInner{
			{PromotionId: promotionID, Budget: budget},
		},
	}
	resp, _, err := c.PromotionBudgetUpdateV30Api().
		Post(ctx).
		AccessToken(accessToken).
		PromotionBudgetUpdateV30Request(req).
		Execute()
	if err != nil {
		log.Fatalf("更新单元预算失败: %v", err)
	}
	checkResp("更新单元预算", resp.Code, resp.Message)
	fmt.Printf("广告单元 %d 预算已更新为 %.2f 元\n", promotionID, budget)
}

// DeletePromotions deletes promotions (v3).
// API: 12-004 PromotionDeleteV30Api
func DeletePromotions(c *ad_open_sdk_go.Client, accessToken string, advID int64, promotionIDs []int64) {
	ctx := context.Background()
	req := models.PromotionDeleteV30Request{
		AdvertiserId: advID,
		PromotionIds: promotionIDs,
	}
	resp, _, err := c.PromotionDeleteV30Api().
		Post(ctx).
		AccessToken(accessToken).
		PromotionDeleteV30Request(req).
		Execute()
	if err != nil {
		log.Fatalf("删除单元失败: %v", err)
	}
	checkResp("删除单元", resp.Code, resp.Message)
	fmt.Println("广告单元已删除")
}

// GetPromotionRejectReason queries promotion rejection reasons (v3).
// API: 12-013 PromotionRejectReasonGetV30Api
func GetPromotionRejectReason(c *ad_open_sdk_go.Client, accessToken string, advID int64, promotionIDs []int64) {
	ctx := context.Background()
	resp, _, err := c.PromotionRejectReasonGetV30Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(advID).
		PromotionIds(promotionIDs).
		Execute()
	if err != nil {
		log.Fatalf("查询单元审核拒绝原因失败: %v", err)
	}
	checkResp("查询单元审核拒绝原因", resp.Code, resp.Message)

	if resp.Data == nil {
		fmt.Println("暂无拒绝原因数据")
		return
	}
	fmt.Println("=== 广告单元审核拒绝原因 ===")
	fmt.Printf("%+v\n", resp.Data)
}

// GetPromotionCostProtectStatus queries promotion cost protection (v3).
// API: 12-014 PromotionCostProtectStatusGetV30Api
func GetPromotionCostProtectStatus(c *ad_open_sdk_go.Client, accessToken string, advID int64, promotionIDs []int64) {
	ctx := context.Background()
	resp, _, err := c.PromotionCostProtectStatusGetV30Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(advID).
		PromotionIds(promotionIDs).
		Execute()
	if err != nil {
		log.Fatalf("查询单元成本保护状态失败: %v", err)
	}
	checkResp("查询单元成本保护状态", resp.Code, resp.Message)

	if resp.Data == nil {
		fmt.Println("暂无成本保护数据")
		return
	}
	fmt.Println("=== 广告单元成本保护状态 ===")
	fmt.Printf("%+v\n", resp.Data)
}
