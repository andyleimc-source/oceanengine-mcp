package api

import (
	"context"
	"fmt"
	"strings"

	ad_open_sdk_go "github.com/oceanengine/ad_open_sdk_go"
	"github.com/oceanengine/ad_open_sdk_go/models"
)

// CreateProject creates a new project (v3).
// API: 11-001 ProjectCreateV30Api
func CreateProject(c *ad_open_sdk_go.Client, accessToken string, advID int64, name string, budget float64, budgetMode string, landingType string, startTime string, endTime string) (string, error) {
	ctx := context.Background()

	var bMode models.ProjectCreateV30DeliverySettingBudgetMode
	switch budgetMode {
	case "day":
		bMode = models.BUDGET_MODE_DAY_ProjectCreateV30DeliverySettingBudgetMode
	case "infinite":
		bMode = models.BUDGET_MODE_INFINITE_ProjectCreateV30DeliverySettingBudgetMode
	default:
		return "", fmt.Errorf("未知预算模式: %s (可选: day, infinite)", budgetMode)
	}

	var lType models.ProjectCreateV30LandingType
	switch landingType {
	case "LINK":
		lType = models.LINK_ProjectCreateV30LandingType
	case "APP":
		lType = models.APP_ProjectCreateV30LandingType
	case "SHOP":
		lType = models.SHOP_ProjectCreateV30LandingType
	default:
		lType = models.ProjectCreateV30LandingType(landingType)
	}

	deliverySetting := models.ProjectCreateV30RequestDeliverySetting{
		BudgetMode: bMode,
	}
	if budget > 0 {
		deliverySetting.Budget = &budget
	}
	if startTime != "" {
		deliverySetting.StartTime = &startTime
	}
	if endTime != "" {
		deliverySetting.EndTime = &endTime
	}

	req := models.ProjectCreateV30Request{
		AdvertiserId: advID,
		Name:         name,
		LandingType:  lType,
		MarketingGoal: models.VIDEO_AND_IMAGE_ProjectCreateV30MarketingGoal,
		AdType:       models.ALL_ProjectCreateV30AdType,
		DeliveryRange: models.ProjectCreateV30RequestDeliveryRange{
			InventoryCatalog: models.UNIVERSAL_SMART_ProjectCreateV30DeliveryRangeInventoryCatalog,
		},
		DeliverySetting: deliverySetting,
	}

	resp, _, err := c.ProjectCreateV30Api().
		Post(ctx).
		AccessToken(accessToken).
		ProjectCreateV30Request(req).
		Execute()
	if err != nil {
		return "", fmt.Errorf("创建项目失败: %w", err)
	}
	if err := CheckResp("创建项目", resp.Code, resp.Message); err != nil {
		return "", err
	}

	var b strings.Builder
	b.WriteString("=== 项目创建成功 ===\n\n")
	if resp.Data != nil && resp.Data.ProjectId != nil {
		fmt.Fprintf(&b, "  %-16s %d\n", "项目ID", *resp.Data.ProjectId)
	}
	writeField(&b, "名称", name)
	writeField(&b, "推广目的", string(lType))
	writeField(&b, "预算模式", budgetMode)
	if budget > 0 {
		writeFieldFloat(&b, "预算(元)", &budget)
	}
	return b.String(), nil
}

// ListProjects queries projects (v3).
// API: 11-002 ProjectListV30Api
func ListProjects(c *ad_open_sdk_go.Client, accessToken string, advID int64) (string, error) {
	ctx := context.Background()
	fields := []string{
		"project_id", "name", "status", "status_first",
		"landing_type", "marketing_goal", "ad_type", "delivery_mode",
		"project_create_time", "project_modify_time",
	}
	resp, _, err := c.ProjectListV30Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(advID).
		Fields(fields).
		Page(1).
		PageSize(100).
		Execute()
	if err != nil {
		return "", fmt.Errorf("查询项目列表失败: %w", err)
	}
	if err := CheckResp("查询项目列表", resp.Code, resp.Message); err != nil {
		return "", err
	}

	var b strings.Builder
	if resp.Data == nil || len(resp.Data.List) == 0 {
		b.WriteString("暂无项目数据\n")
		return b.String(), nil
	}

	b.WriteString("=== 项目列表 (v3) ===\n\n")
	for i, item := range resp.Data.List {
		fmt.Fprintf(&b, "#%d\n", i+1)
		writeField(&b, "项目ID", item.ProjectId)
		writeField(&b, "名称", item.Name)
		writeField(&b, "状态", item.Status)
		writeField(&b, "一级状态", item.StatusFirst)
		writeField(&b, "推广目的", item.LandingType)
		writeField(&b, "营销目标", item.MarketingGoal)
		writeField(&b, "投放类型", item.AdType)
		writeField(&b, "投放模式", item.DeliveryMode)
		writeField(&b, "创建时间", item.ProjectCreateTime)
		b.WriteString("  ----------\n")
	}
	writePageInfo(&b, resp.Data.PageInfo)
	return b.String(), nil
}

// UpdateProjectStatus updates project status (v3).
// API: 11-005 ProjectStatusUpdateV30Api
func UpdateProjectStatus(c *ad_open_sdk_go.Client, accessToken string, advID int64, projectIDs []int64, optStatus string) (string, error) {
	ctx := context.Background()

	var status models.ProjectStatusUpdateV30DataOptStatus
	switch optStatus {
	case "enable":
		status = models.ENABLE_ProjectStatusUpdateV30DataOptStatus
	case "disable":
		status = models.DISABLE_ProjectStatusUpdateV30DataOptStatus
	default:
		return "", fmt.Errorf("未知状态: %s (可选: enable, disable)", optStatus)
	}

	data := make([]*models.ProjectStatusUpdateV30RequestDataInner, len(projectIDs))
	for i, id := range projectIDs {
		data[i] = &models.ProjectStatusUpdateV30RequestDataInner{
			ProjectId: id,
			OptStatus: status,
		}
	}

	req := models.ProjectStatusUpdateV30Request{
		AdvertiserId: advID,
		Data:         data,
	}
	resp, _, err := c.ProjectStatusUpdateV30Api().
		Post(ctx).
		AccessToken(accessToken).
		ProjectStatusUpdateV30Request(req).
		Execute()
	if err != nil {
		return "", fmt.Errorf("更新项目状态失败: %w", err)
	}
	if err := CheckResp("更新项目状态", resp.Code, resp.Message); err != nil {
		return "", err
	}
	return fmt.Sprintf("项目状态已更新为: %s\n", optStatus), nil
}

// UpdateProjectBudget updates project budget (v3).
// API: 11-006 ProjectBudgetUpdateV30Api
func UpdateProjectBudget(c *ad_open_sdk_go.Client, accessToken string, advID int64, projectID int64, budget float64, budgetMode string) (string, error) {
	ctx := context.Background()

	var mode models.ProjectBudgetUpdateV30DataBudgetMode
	switch budgetMode {
	case "day":
		mode = models.BUDGET_MODE_DAY_ProjectBudgetUpdateV30DataBudgetMode
	case "infinite":
		mode = models.BUDGET_MODE_INFINITE_ProjectBudgetUpdateV30DataBudgetMode
	default:
		return "", fmt.Errorf("未知预算模式: %s (可选: day, infinite)", budgetMode)
	}

	req := models.ProjectBudgetUpdateV30Request{
		AdvertiserId: advID,
		Data: []*models.ProjectBudgetUpdateV30RequestDataInner{
			{
				ProjectId:  projectID,
				BudgetMode: mode,
				Budget:     &budget,
			},
		},
	}
	resp, _, err := c.ProjectBudgetUpdateV30Api().
		Post(ctx).
		AccessToken(accessToken).
		ProjectBudgetUpdateV30Request(req).
		Execute()
	if err != nil {
		return "", fmt.Errorf("更新项目预算失败: %w", err)
	}
	if err := CheckResp("更新项目预算", resp.Code, resp.Message); err != nil {
		return "", err
	}
	return fmt.Sprintf("项目 %d 预算已更新为 %.2f 元 (模式: %s)\n", projectID, budget, budgetMode), nil
}

// DeleteProjects deletes projects (v3).
// API: 11-004 ProjectDeleteV30Api
func DeleteProjects(c *ad_open_sdk_go.Client, accessToken string, advID int64, projectIDs []int64) (string, error) {
	ctx := context.Background()
	req := models.ProjectDeleteV30Request{
		AdvertiserId: advID,
		ProjectIds:   projectIDs,
	}
	resp, _, err := c.ProjectDeleteV30Api().
		Post(ctx).
		AccessToken(accessToken).
		ProjectDeleteV30Request(req).
		Execute()
	if err != nil {
		return "", fmt.Errorf("删除项目失败: %w", err)
	}
	if err := CheckResp("删除项目", resp.Code, resp.Message); err != nil {
		return "", err
	}
	return "项目已删除\n", nil
}

// GetProjectCostProtectStatus queries project cost protection (v3).
// API: 11-013 ProjectCostProtectStatusGetV30Api
func GetProjectCostProtectStatus(c *ad_open_sdk_go.Client, accessToken string, advID int64, projectIDs []int64) (string, error) {
	ctx := context.Background()
	resp, _, err := c.ProjectCostProtectStatusGetV30Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(advID).
		ProjectIds(projectIDs).
		Execute()
	if err != nil {
		return "", fmt.Errorf("查询项目成本保护状态失败: %w", err)
	}
	if err := CheckResp("查询项目成本保护状态", resp.Code, resp.Message); err != nil {
		return "", err
	}

	var b strings.Builder
	if resp.Data == nil || len(resp.Data.CompensateStatusInfoList) == 0 {
		b.WriteString("暂无成本保护数据\n")
		return b.String(), nil
	}
	b.WriteString("=== 项目成本保护状态 ===\n\n")
	for _, item := range resp.Data.CompensateStatusInfoList {
		writeField(&b, "项目ID", item.ProjectId)
		writeField(&b, "保障状态", item.CompensateStatus)
		writeFieldFloat(&b, "赔付金额(元)", item.CompensateAmount)
		if len(item.CompensateEndReasons) > 0 {
			fmt.Fprintf(&b, "  %-16s", "结束原因")
			for _, r := range item.CompensateEndReasons {
				fmt.Fprintf(&b, " %s", r)
			}
			b.WriteString("\n")
		}
		if len(item.CompensateInvalidReasons) > 0 {
			fmt.Fprintf(&b, "  %-16s", "失效原因")
			for _, r := range item.CompensateInvalidReasons {
				fmt.Fprintf(&b, " %s", r)
			}
			b.WriteString("\n")
		}
		writeField(&b, "赔付规则", item.CompensateUrl)
		b.WriteString("  ----------\n")
	}
	return b.String(), nil
}
