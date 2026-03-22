package api

import (
	"context"
	"fmt"
	"strings"

	ad_open_sdk_go "github.com/oceanengine/ad_open_sdk_go"
	"github.com/oceanengine/ad_open_sdk_go/models"
)

// ListPromotions queries promotions/ad units (v3).
// API: 12-002 PromotionListV30Api
func ListPromotions(c *ad_open_sdk_go.Client, accessToken string, advID int64) (string, error) {
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
		return "", fmt.Errorf("查询广告单元列表失败: %w", err)
	}
	if err := CheckResp("查询广告单元列表", resp.Code, resp.Message); err != nil {
		return "", err
	}

	var b strings.Builder
	if resp.Data == nil || len(resp.Data.List) == 0 {
		b.WriteString("暂无广告单元数据\n")
		return b.String(), nil
	}

	b.WriteString("=== 广告单元列表 (v3) ===\n\n")
	for i, item := range resp.Data.List {
		fmt.Fprintf(&b, "#%d\n", i+1)
		writeField(&b, "单元ID", item.PromotionId)
		writeField(&b, "名称", item.PromotionName)
		writeField(&b, "所属项目ID", item.ProjectId)
		writeField(&b, "状态", item.Status)
		writeField(&b, "一级状态", item.StatusFirst)
		writeField(&b, "预算模式", item.BudgetMode)
		writeFieldFloat(&b, "预算(元)", item.Budget)
		writeFieldFloat(&b, "出价(元)", item.Bid)
		writeFieldFloat(&b, "CPA出价(元)", item.CpaBid)
		writeFieldFloat(&b, "深度出价(元)", item.DeepCpabid)
		writeFieldFloat(&b, "ROI目标", item.RoiGoal)
		writeField(&b, "学习期状态", item.LearningPhase)
		writeField(&b, "创建时间", item.PromotionCreateTime)
		b.WriteString("  ----------\n")
	}
	writePageInfo(&b, resp.Data.PageInfo)
	return b.String(), nil
}

// UpdatePromotionStatus updates promotion status (v3).
// API: 12-005 PromotionStatusUpdateV30Api
func UpdatePromotionStatus(c *ad_open_sdk_go.Client, accessToken string, advID int64, promotionIDs []int64, optStatus string) (string, error) {
	ctx := context.Background()

	var status models.PromotionStatusUpdateV30DataOptStatus
	switch optStatus {
	case "enable":
		status = models.ENABLE_PromotionStatusUpdateV30DataOptStatus
	case "disable":
		status = models.DISABLE_PromotionStatusUpdateV30DataOptStatus
	default:
		return "", fmt.Errorf("未知状态: %s (可选: enable, disable)", optStatus)
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
		return "", fmt.Errorf("更新单元状态失败: %w", err)
	}
	if err := CheckResp("更新单元状态", resp.Code, resp.Message); err != nil {
		return "", err
	}
	return fmt.Sprintf("广告单元状态已更新为: %s\n", optStatus), nil
}

// UpdatePromotionBid batch updates promotion bids (v3).
// API: 12-006 PromotionBidUpdateV30Api
func UpdatePromotionBid(c *ad_open_sdk_go.Client, accessToken string, advID int64, promotionID int64, bid float64) (string, error) {
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
		return "", fmt.Errorf("更新单元出价失败: %w", err)
	}
	if err := CheckResp("更新单元出价", resp.Code, resp.Message); err != nil {
		return "", err
	}
	return fmt.Sprintf("广告单元 %d 出价已更新为 %.2f 元\n", promotionID, bid), nil
}

// UpdatePromotionBudget updates promotion budget (v3).
// API: 12-007 PromotionBudgetUpdateV30Api
func UpdatePromotionBudget(c *ad_open_sdk_go.Client, accessToken string, advID int64, promotionID int64, budget float64) (string, error) {
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
		return "", fmt.Errorf("更新单元预算失败: %w", err)
	}
	if err := CheckResp("更新单元预算", resp.Code, resp.Message); err != nil {
		return "", err
	}
	return fmt.Sprintf("广告单元 %d 预算已更新为 %.2f 元\n", promotionID, budget), nil
}

// DeletePromotions deletes promotions (v3).
// API: 12-004 PromotionDeleteV30Api
func DeletePromotions(c *ad_open_sdk_go.Client, accessToken string, advID int64, promotionIDs []int64) (string, error) {
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
		return "", fmt.Errorf("删除单元失败: %w", err)
	}
	if err := CheckResp("删除单元", resp.Code, resp.Message); err != nil {
		return "", err
	}
	return "广告单元已删除\n", nil
}

// GetPromotionRejectReason queries promotion rejection reasons (v3).
// API: 12-013 PromotionRejectReasonGetV30Api
func GetPromotionRejectReason(c *ad_open_sdk_go.Client, accessToken string, advID int64, promotionIDs []int64) (string, error) {
	ctx := context.Background()
	resp, _, err := c.PromotionRejectReasonGetV30Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(advID).
		PromotionIds(promotionIDs).
		Execute()
	if err != nil {
		return "", fmt.Errorf("查询单元审核拒绝原因失败: %w", err)
	}
	if err := CheckResp("查询单元审核拒绝原因", resp.Code, resp.Message); err != nil {
		return "", err
	}

	var b strings.Builder
	if resp.Data == nil || len(resp.Data.List) == 0 {
		b.WriteString("暂无拒绝原因数据\n")
		return b.String(), nil
	}
	b.WriteString("=== 广告单元审核拒绝原因 ===\n\n")
	for _, item := range resp.Data.List {
		writeField(&b, "单元ID", item.PromotionId)
		for _, pr := range item.PromotionReject {
			writeField(&b, "  审核项", pr.Content)
			for _, reason := range pr.RejectReason {
				fmt.Fprintf(&b, "  %-16s %s\n", "拒绝原因", reason)
			}
			for _, sug := range pr.Suggestion {
				fmt.Fprintf(&b, "  %-16s %s\n", "建议", sug)
			}
		}
		for _, mr := range item.MaterialReject {
			writeField(&b, "  素材类型", mr.Type)
			writeField(&b, "  素材项", mr.Item)
			for _, reason := range mr.RejectReason {
				fmt.Fprintf(&b, "  %-16s %s\n", "拒绝原因", reason)
			}
			for _, sug := range mr.Suggestion {
				fmt.Fprintf(&b, "  %-16s %s\n", "建议", sug)
			}
		}
		b.WriteString("  ----------\n")
	}
	return b.String(), nil
}

// GetPromotionCostProtectStatus queries promotion cost protection (v3).
// API: 12-014 PromotionCostProtectStatusGetV30Api
func GetPromotionCostProtectStatus(c *ad_open_sdk_go.Client, accessToken string, advID int64, promotionIDs []int64) (string, error) {
	ctx := context.Background()
	resp, _, err := c.PromotionCostProtectStatusGetV30Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(advID).
		PromotionIds(promotionIDs).
		Execute()
	if err != nil {
		return "", fmt.Errorf("查询单元成本保护状态失败: %w", err)
	}
	if err := CheckResp("查询单元成本保护状态", resp.Code, resp.Message); err != nil {
		return "", err
	}

	var b strings.Builder
	if resp.Data == nil || len(resp.Data.CompensateStatusInfoList) == 0 {
		b.WriteString("暂无成本保护数据\n")
		return b.String(), nil
	}
	b.WriteString("=== 广告单元成本保护状态 ===\n\n")
	for _, item := range resp.Data.CompensateStatusInfoList {
		writeField(&b, "单元ID", item.QueryId)
		writeField(&b, "保障状态", item.CompensateStatus)
		writeFieldFloat(&b, "赔付金额(元)", item.CompensateAmount)
		writeField(&b, "查询状态", item.Status)
		writeField(&b, "赔付规则", item.Url)
		b.WriteString("  ----------\n")
	}
	return b.String(), nil
}
