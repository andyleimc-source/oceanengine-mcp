package api

import (
	"context"
	"fmt"
	"strings"

	ad_open_sdk_go "github.com/oceanengine/ad_open_sdk_go"
	"github.com/oceanengine/ad_open_sdk_go/models"
)

// ListAds queries ad plans (v2).
// API: 10-001 AdGetV2Api
// NOTE: SDK 将 AdGetV2ResponseData 建模为扁平结构（无 List 字段），
// 实际 API 返回列表但 SDK 只解析出单条。这是 SDK 代码生成的已知限制。
func ListAds(c *ad_open_sdk_go.Client, accessToken string, advID int64) (string, error) {
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
		return "", fmt.Errorf("查询广告计划失败: %w", err)
	}
	if err := CheckResp("查询广告计划", resp.Code, resp.Message); err != nil {
		return "", err
	}

	var b strings.Builder
	if resp.Data == nil || resp.Data.AdId == nil {
		b.WriteString("暂无广告计划数据\n")
		return b.String(), nil
	}

	b.WriteString("=== 广告计划列表 ===\n\n")
	writeField(&b, "计划ID", resp.Data.AdId)
	writeField(&b, "名称", resp.Data.Name)
	writeField(&b, "广告组ID", resp.Data.CampaignId)
	writeFieldFloat(&b, "预算(元)", &resp.Data.Budget)
	writeFieldFloat(&b, "出价(元)", resp.Data.Bid)
	writeFieldFloat(&b, "CPA出价(元)", resp.Data.CpaBid)
	writeField(&b, "状态", resp.Data.Status)
	writeField(&b, "创建时间", resp.Data.AdCreateTime)
	b.WriteString("\n(注: SDK 限制，仅返回单条记录)\n")
	return b.String(), nil
}

// UpdateAdBid batch updates ad plan bids (v2).
// API: 10-002 AdUpdateBidV2Api
func UpdateAdBid(c *ad_open_sdk_go.Client, accessToken string, advID int64, adID int64, bid float64) (string, error) {
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
		return "", fmt.Errorf("更新出价失败: %w", err)
	}
	if err := CheckResp("更新出价", resp.Code, resp.Message); err != nil {
		return "", err
	}
	return fmt.Sprintf("广告计划 %d 出价已更新为 %.2f 元\n", adID, bid), nil
}

// UpdateAdBudget batch updates ad plan budgets (v2).
// API: 10-003 AdUpdateBudgetV2Api
func UpdateAdBudget(c *ad_open_sdk_go.Client, accessToken string, advID int64, adID int64, budget float64) (string, error) {
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
		return "", fmt.Errorf("更新预算失败: %w", err)
	}
	if err := CheckResp("更新预算", resp.Code, resp.Message); err != nil {
		return "", err
	}
	return fmt.Sprintf("广告计划 %d 预算已更新为 %.2f 元\n", adID, budget), nil
}

// GetAdRejectReason queries ad rejection reasons (v2).
// API: 10-005 AdRejectReasonV2Api
func GetAdRejectReason(c *ad_open_sdk_go.Client, accessToken string, advID int64, adIDs []int64) (string, error) {
	ctx := context.Background()
	resp, _, err := c.AdRejectReasonV2Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(advID).
		AdIds(adIDs).
		Execute()
	if err != nil {
		return "", fmt.Errorf("查询审核拒绝原因失败: %w", err)
	}
	if err := CheckResp("查询审核拒绝原因", resp.Code, resp.Message); err != nil {
		return "", err
	}

	var b strings.Builder
	if resp.Data == nil || len(resp.Data.List) == 0 {
		b.WriteString("暂无拒绝原因数据\n")
		return b.String(), nil
	}
	b.WriteString("=== 广告计划审核拒绝原因 ===\n\n")
	for _, item := range resp.Data.List {
		if item.AdReject != nil {
			writeField(&b, "计划ID", item.AdReject.AdId)
			for _, rd := range item.AdReject.RejectData {
				writeField(&b, "  拒绝项", rd.RejectItem)
				writeField(&b, "  拒绝原因", rd.RejectReason)
			}
		}
		for _, cr := range item.CreativeReject {
			writeField(&b, "创意ID", cr.CreativeId)
			for _, rd := range cr.RejectData {
				writeField(&b, "  拒绝项", rd.RejectItem)
				writeField(&b, "  拒绝原因", rd.RejectReason)
			}
			for _, mr := range cr.MaterialReject {
				writeField(&b, "  素材标题", mr.Title)
				writeField(&b, "  拒绝原因", mr.RejectReason)
			}
		}
		for _, mr := range item.MaterialReject {
			writeField(&b, "素材标题", mr.Title)
			writeField(&b, "拒绝原因", mr.RejectReason)
		}
		b.WriteString("  ----------\n")
	}
	return b.String(), nil
}

// GetAdCostProtectStatus queries cost protection status (v2).
// API: 10-006 AdCostProtectStatusGetV2Api
func GetAdCostProtectStatus(c *ad_open_sdk_go.Client, accessToken string, advID int64, adIDs []int64) (string, error) {
	ctx := context.Background()
	resp, _, err := c.AdCostProtectStatusGetV2Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(advID).
		AdIds(adIDs).
		Execute()
	if err != nil {
		return "", fmt.Errorf("查询成本保护状态失败: %w", err)
	}
	if err := CheckResp("查询成本保护状态", resp.Code, resp.Message); err != nil {
		return "", err
	}

	var b strings.Builder
	if resp.Data == nil || len(resp.Data.List) == 0 {
		b.WriteString("暂无成本保护数据\n")
		return b.String(), nil
	}
	b.WriteString("=== 成本保护状态 ===\n\n")
	for _, item := range resp.Data.List {
		writeField(&b, "计划ID", item.AdId)
		writeField(&b, "状态", item.Status)
		b.WriteString("  ----------\n")
	}
	return b.String(), nil
}
