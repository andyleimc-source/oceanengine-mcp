package api

import (
	"context"
	"fmt"
	"strings"

	ad_open_sdk_go "github.com/oceanengine/ad_open_sdk_go"
)

// ListCreatives queries creatives (v2).
// API: 16-001 CreativeGetV2Api
func ListCreatives(c *ad_open_sdk_go.Client, accessToken string, advID int64) (string, error) {
	ctx := context.Background()
	fields := []string{
		"creative_id", "ad_id", "campaign_id", "title",
		"status", "image_mode", "creative_create_time", "creative_modify_time",
	}
	resp, _, err := c.CreativeGetV2Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(advID).
		Fields(fields).
		Page(1).
		PageSize(100).
		Execute()
	if err != nil {
		return "", fmt.Errorf("查询创意列表失败: %w", err)
	}
	if err := CheckResp("查询创意列表", resp.Code, resp.Message); err != nil {
		return "", err
	}

	var b strings.Builder
	if resp.Data == nil || len(resp.Data.List) == 0 {
		b.WriteString("暂无创意数据\n")
		return b.String(), nil
	}

	b.WriteString("=== 创意列表 ===\n\n")
	for i, item := range resp.Data.List {
		fmt.Fprintf(&b, "#%d\n", i+1)
		writeField(&b, "创意ID", item.CreativeId)
		writeField(&b, "计划ID", item.AdId)
		writeField(&b, "广告主ID", item.AdvertiserId)
		writeField(&b, "标题", item.Title)
		writeField(&b, "状态", item.Status)
		writeField(&b, "素材类型", item.ImageMode)
		writeField(&b, "创建时间", item.CreativeCreateTime)
		b.WriteString("  ----------\n")
	}
	writePageInfo(&b, resp.Data.PageInfo)
	return b.String(), nil
}

// GetCreativeDetail queries creative detail (v3).
// API: 16-002 CreativeDetailGetV30Api
func GetCreativeDetail(c *ad_open_sdk_go.Client, accessToken string, advID int64, adID int64) (string, error) {
	ctx := context.Background()
	resp, _, err := c.CreativeDetailGetV30Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(advID).
		AdId(adID).
		Execute()
	if err != nil {
		return "", fmt.Errorf("查询创意详情失败: %w", err)
	}
	if err := CheckResp("查询创意详情", resp.Code, resp.Message); err != nil {
		return "", err
	}

	var b strings.Builder
	if resp.Data == nil {
		b.WriteString("暂无创意详情数据\n")
		return b.String(), nil
	}
	b.WriteString("=== 创意详情 ===\n")
	writeField(&b, "广告ID", resp.Data.AdId)
	writeField(&b, "广告主ID", resp.Data.AdvertiserId)
	if resp.Data.Creative != nil {
		fmt.Fprintf(&b, "  创意数据: %+v\n", resp.Data.Creative)
	}
	if len(resp.Data.CreativeList) > 0 {
		fmt.Fprintf(&b, "  创意列表: %d 条\n", len(resp.Data.CreativeList))
		for i, cr := range resp.Data.CreativeList {
			fmt.Fprintf(&b, "    [%d] %+v\n", i+1, cr)
		}
	}
	return b.String(), nil
}

// GetCreativeRejectReason queries creative rejection reasons (v2).
// API: 16-007 CreativeRejectReasonV2Api
func GetCreativeRejectReason(c *ad_open_sdk_go.Client, accessToken string, advID int64, creativeIDs []int64) (string, error) {
	ctx := context.Background()
	resp, _, err := c.CreativeRejectReasonV2Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(advID).
		CreativeIds(creativeIDs).
		Execute()
	if err != nil {
		return "", fmt.Errorf("查询创意审核拒绝原因失败: %w", err)
	}
	if err := CheckResp("查询创意审核拒绝原因", resp.Code, resp.Message); err != nil {
		return "", err
	}

	var b strings.Builder
	if resp.Data == nil {
		b.WriteString("暂无拒绝原因数据\n")
		return b.String(), nil
	}
	b.WriteString("=== 创意审核拒绝原因 ===\n")
	fmt.Fprintf(&b, "%+v\n", resp.Data)
	return b.String(), nil
}
