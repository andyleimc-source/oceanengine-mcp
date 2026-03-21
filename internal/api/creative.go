package api

import (
	"context"
	"fmt"
	"log"

	ad_open_sdk_go "github.com/oceanengine/ad_open_sdk_go"
)

// ListCreatives queries creatives (v2).
// API: 16-001 CreativeGetV2Api
func ListCreatives(c *ad_open_sdk_go.Client, accessToken string, advID int64) {
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
		log.Fatalf("查询创意列表失败: %v", err)
	}
	checkResp("查询创意列表", resp.Code, resp.Message)

	if resp.Data == nil || len(resp.Data.List) == 0 {
		fmt.Println("暂无创意数据")
		return
	}

	fmt.Println("=== 创意列表 ===")
	fmt.Println()
	for i, item := range resp.Data.List {
		fmt.Printf("#%d\n", i+1)
		printField("创意ID", item.CreativeId)
		printField("计划ID", item.AdId)
		printField("广告主ID", item.AdvertiserId)
		printField("标题", item.Title)
		printField("状态", item.Status)
		printField("素材类型", item.ImageMode)
		printField("创建时间", item.CreativeCreateTime)
		fmt.Println("  ----------")
	}
	printPageInfo(resp.Data.PageInfo)
}

// GetCreativeDetail queries creative detail (v3).
// API: 16-002 CreativeDetailGetV30Api
func GetCreativeDetail(c *ad_open_sdk_go.Client, accessToken string, advID int64, adID int64) {
	ctx := context.Background()
	resp, _, err := c.CreativeDetailGetV30Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(advID).
		AdId(adID).
		Execute()
	if err != nil {
		log.Fatalf("查询创意详情失败: %v", err)
	}
	checkResp("查询创意详情", resp.Code, resp.Message)

	if resp.Data == nil {
		fmt.Println("暂无创意详情数据")
		return
	}
	fmt.Println("=== 创意详情 ===")
	printField("广告ID", resp.Data.AdId)
	printField("广告主ID", resp.Data.AdvertiserId)
	if resp.Data.Creative != nil {
		fmt.Printf("  创意数据: %+v\n", resp.Data.Creative)
	}
	if len(resp.Data.CreativeList) > 0 {
		fmt.Printf("  创意列表: %d 条\n", len(resp.Data.CreativeList))
		for i, cr := range resp.Data.CreativeList {
			fmt.Printf("    [%d] %+v\n", i+1, cr)
		}
	}
}

// GetCreativeRejectReason queries creative rejection reasons (v2).
// API: 16-007 CreativeRejectReasonV2Api
func GetCreativeRejectReason(c *ad_open_sdk_go.Client, accessToken string, advID int64, creativeIDs []int64) {
	ctx := context.Background()
	resp, _, err := c.CreativeRejectReasonV2Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(advID).
		CreativeIds(creativeIDs).
		Execute()
	if err != nil {
		log.Fatalf("查询创意审核拒绝原因失败: %v", err)
	}
	checkResp("查询创意审核拒绝原因", resp.Code, resp.Message)

	if resp.Data == nil {
		fmt.Println("暂无拒绝原因数据")
		return
	}
	fmt.Println("=== 创意审核拒绝原因 ===")
	fmt.Printf("%+v\n", resp.Data)
}
