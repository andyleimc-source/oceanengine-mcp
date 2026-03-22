package api

import (
	"context"
	"fmt"
	"strings"

	ad_open_sdk_go "github.com/oceanengine/ad_open_sdk_go"
	"github.com/oceanengine/ad_open_sdk_go/models"
	"github.com/andyleimc-source/oceanengine-mcp/pkg/format"
)

// ReportLevel determines the grouping dimension for reports.
type ReportLevel string

const (
	LevelAdvertiser ReportLevel = "advertiser"
	LevelProject    ReportLevel = "campaign"
	LevelPromotion  ReportLevel = "ad"
)

type ReportParams struct {
	Client       *ad_open_sdk_go.Client
	AccessToken  string
	AdvertiserID int64
	StartDate    string
	EndDate      string
	Level        ReportLevel
}

// FetchReport fetches a v3 custom report at the given level.
// API: 20-005 ReportCustomGetV30Api
func FetchReport(p ReportParams) (string, error) {
	ctx := context.Background()

	var metrics []string
	var dimensions []string
	var title string

	switch p.Level {
	case LevelAdvertiser:
		metrics = AllMetrics
		dimensions = []string{"stat_time_day"}
		title = "广告主报表"
	case LevelProject:
		metrics = CoreMetrics
		dimensions = []string{"cdp_project_id", "cdp_project_name", "stat_time_day"}
		title = "项目报表 (按消耗降序 Top 20)"
	case LevelPromotion:
		metrics = CoreMetrics
		dimensions = []string{"cdp_promotion_id", "cdp_promotion_name", "stat_time_day"}
		title = "单元报表 (按消耗降序 Top 20)"
	}

	pageSize := int32(100)
	if p.Level != LevelAdvertiser {
		pageSize = 20
	}

	resp, _, err := p.Client.ReportCustomGetV30Api().
		Get(ctx).
		AccessToken(p.AccessToken).
		AdvertiserId(p.AdvertiserID).
		StartTime(p.StartDate).
		EndTime(p.EndDate).
		Metrics(metrics).
		Dimensions(dimensions).
		DataTopic(models.BASIC_DATA_ReportCustomGetV30DataTopic).
		Filters([]*models.ReportCustomGetV30FiltersInner{}).
		OrderBy([]*models.ReportCustomGetV30OrderByInner{
			{Field: "stat_cost", Type: models.DESC_ReportCustomGetV30OrderByType.Ptr()},
		}).
		Page(1).
		PageSize(pageSize).
		Execute()
	if err != nil {
		return "", fmt.Errorf("拉取%s失败: %w", title, err)
	}
	if err := CheckResp(title, resp.Code, resp.Message); err != nil {
		return "", err
	}

	var b strings.Builder
	fmt.Fprintf(&b, "=== %s ===\n", title)
	fmt.Fprintf(&b, "日期范围: %s ~ %s\n\n", p.StartDate, p.EndDate)

	if resp.Data == nil || len(resp.Data.Rows) == 0 {
		b.WriteString("暂无数据\n")
		return b.String(), nil
	}

	for i, row := range resp.Data.Rows {
		switch p.Level {
		case LevelAdvertiser:
			if d := row.Dimensions["stat_time_day"]; d != "" {
				fmt.Fprintf(&b, "日期: %s\n", d)
			}
		case LevelProject:
			name := row.Dimensions["cdp_project_name"]
			if name == "" {
				name = row.Dimensions["cdp_project_id"]
			}
			fmt.Fprintf(&b, "#%d %s\n", i+1, name)
		case LevelPromotion:
			name := row.Dimensions["cdp_promotion_name"]
			if name == "" {
				name = row.Dimensions["cdp_promotion_id"]
			}
			fmt.Fprintf(&b, "#%d %s\n", i+1, name)
		}
		format.PrintMetricsOrdered(&b, row.Metrics, metrics)
		format.Separator(&b)
	}

	if p.Level == LevelAdvertiser && resp.Data.TotalMetrics != nil && len(resp.Data.TotalMetrics) > 0 {
		b.WriteString("=== 汇总 ===\n")
		format.PrintMetricsSorted(&b, resp.Data.TotalMetrics)
	}
	return b.String(), nil
}

// FetchReportConfig queries available metrics and dimensions.
// API: 20-007 ReportCustomConfigGetV30Api
func FetchReportConfig(client *ad_open_sdk_go.Client, accessToken string, advertiserID int64) (string, error) {
	ctx := context.Background()
	topics := []*models.ReportCustomConfigGetV30DataTopics{
		models.BASIC_DATA_ReportCustomConfigGetV30DataTopics.Ptr(),
	}
	resp, _, err := client.ReportCustomConfigGetV30Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(advertiserID).
		DataTopics(topics).
		Execute()
	if err != nil {
		return "", fmt.Errorf("查询报表配置失败: %w", err)
	}
	if err := CheckResp("查询报表配置", resp.Code, resp.Message); err != nil {
		return "", err
	}

	var b strings.Builder
	if resp.Data == nil {
		b.WriteString("无配置数据\n")
		return b.String(), nil
	}
	for _, item := range resp.Data.List {
		topic := "-"
		if item.DataTopic != nil {
			topic = string(*item.DataTopic)
		}
		fmt.Fprintf(&b, "=== 数据集: %s ===\n\n", topic)

		b.WriteString("可用维度:\n")
		for _, d := range item.Dimensions {
			field, name := "", ""
			if d.Field != nil {
				field = *d.Field
			}
			if d.Name != nil {
				name = *d.Name
			}
			fmt.Fprintf(&b, "  %-30s %s\n", field, name)
		}

		b.WriteString("\n可用指标:\n")
		for _, m := range item.Metrics {
			field, name := "", ""
			if m.Field != nil {
				field = *m.Field
			}
			if m.Name != nil {
				name = *m.Name
			}
			fmt.Fprintf(&b, "  %-30s %s\n", field, name)
		}
		b.WriteString("\n")
	}
	return b.String(), nil
}
