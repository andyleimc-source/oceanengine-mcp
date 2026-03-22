package api

import (
	"context"
	"fmt"
	"strings"

	ad_open_sdk_go "github.com/oceanengine/ad_open_sdk_go"
	"github.com/oceanengine/ad_open_sdk_go/models"
)

// ListCampaigns queries ad groups (v2).
// API: 09-002 CampaignGetV2Api
func ListCampaigns(c *ad_open_sdk_go.Client, accessToken string, advID int64) (string, error) {
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
		return "", fmt.Errorf("查询广告组失败: %w", err)
	}
	if err := CheckResp("查询广告组", resp.Code, resp.Message); err != nil {
		return "", err
	}

	var b strings.Builder
	if resp.Data == nil || len(resp.Data.List) == 0 {
		b.WriteString("暂无广告组数据\n")
		return b.String(), nil
	}

	b.WriteString("=== 广告组列表 ===\n\n")
	for i, item := range resp.Data.List {
		fmt.Fprintf(&b, "#%d\n", i+1)
		writeField(&b, "广告组ID", item.Id)
		writeField(&b, "名称", item.Name)
		writeField(&b, "预算模式", item.BudgetMode)
		writeFieldFloat(&b, "预算(元)", item.Budget)
		writeField(&b, "推广目的", item.LandingType)
		writeField(&b, "状态", item.Status)
		writeField(&b, "创建时间", item.CampaignCreateTime)
		b.WriteString("  ----------\n")
	}
	writePageInfo(&b, resp.Data.PageInfo)
	return b.String(), nil
}

// UpdateCampaignStatus updates ad group status (v2).
// API: 09-004 CampaignUpdateStatusV2Api
func UpdateCampaignStatus(c *ad_open_sdk_go.Client, accessToken string, advID int64, campaignIDs []int64, optStatus string) (string, error) {
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
		return "", fmt.Errorf("未知状态: %s (可选: enable, disable, delete)", optStatus)
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
		return "", fmt.Errorf("更新广告组状态失败: %w", err)
	}
	if err := CheckResp("更新广告组状态", resp.Code, resp.Message); err != nil {
		return "", err
	}
	return fmt.Sprintf("广告组状态已更新为: %s\n", optStatus), nil
}
