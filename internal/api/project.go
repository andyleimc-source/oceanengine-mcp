package api

import (
	"context"
	"fmt"
	"log"

	ad_open_sdk_go "github.com/oceanengine/ad_open_sdk_go"
	"github.com/oceanengine/ad_open_sdk_go/models"
)

// ListProjects queries projects (v3).
// API: 11-002 ProjectListV30Api
func ListProjects(c *ad_open_sdk_go.Client, accessToken string, advID int64) {
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
		log.Fatalf("查询项目列表失败: %v", err)
	}
	checkResp("查询项目列表", resp.Code, resp.Message)

	if resp.Data == nil || len(resp.Data.List) == 0 {
		fmt.Println("暂无项目数据")
		return
	}

	fmt.Println("=== 项目列表 (v3) ===")
	fmt.Println()
	for i, item := range resp.Data.List {
		fmt.Printf("#%d\n", i+1)
		printField("项目ID", item.ProjectId)
		printField("名称", item.Name)
		printField("状态", item.Status)
		printField("一级状态", item.StatusFirst)
		printField("推广目的", item.LandingType)
		printField("营销目标", item.MarketingGoal)
		printField("投放类型", item.AdType)
		printField("投放模式", item.DeliveryMode)
		printField("创建时间", item.ProjectCreateTime)
		fmt.Println("  ----------")
	}
	printPageInfo(resp.Data.PageInfo)
}

// UpdateProjectStatus updates project status (v3).
// API: 11-005 ProjectStatusUpdateV30Api
func UpdateProjectStatus(c *ad_open_sdk_go.Client, accessToken string, advID int64, projectIDs []int64, optStatus string) {
	ctx := context.Background()

	var status models.ProjectStatusUpdateV30DataOptStatus
	switch optStatus {
	case "enable":
		status = models.ENABLE_ProjectStatusUpdateV30DataOptStatus
	case "disable":
		status = models.DISABLE_ProjectStatusUpdateV30DataOptStatus
	default:
		log.Fatalf("未知状态: %s (可选: enable, disable)", optStatus)
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
		log.Fatalf("更新项目状态失败: %v", err)
	}
	checkResp("更新项目状态", resp.Code, resp.Message)
	fmt.Printf("项目状态已更新为: %s\n", optStatus)
}

// UpdateProjectBudget updates project budget (v3).
// API: 11-006 ProjectBudgetUpdateV30Api
func UpdateProjectBudget(c *ad_open_sdk_go.Client, accessToken string, advID int64, projectID int64, budget float64, budgetMode string) {
	ctx := context.Background()

	var mode models.ProjectBudgetUpdateV30DataBudgetMode
	switch budgetMode {
	case "day":
		mode = models.BUDGET_MODE_DAY_ProjectBudgetUpdateV30DataBudgetMode
	case "infinite":
		mode = models.BUDGET_MODE_INFINITE_ProjectBudgetUpdateV30DataBudgetMode
	default:
		log.Fatalf("未知预算模式: %s (可选: day, infinite)", budgetMode)
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
		log.Fatalf("更新项目预算失败: %v", err)
	}
	checkResp("更新项目预算", resp.Code, resp.Message)
	fmt.Printf("项目 %d 预算已更新为 %.2f 元 (模式: %s)\n", projectID, budget, budgetMode)
}

// DeleteProjects deletes projects (v3).
// API: 11-004 ProjectDeleteV30Api
func DeleteProjects(c *ad_open_sdk_go.Client, accessToken string, advID int64, projectIDs []int64) {
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
		log.Fatalf("删除项目失败: %v", err)
	}
	checkResp("删除项目", resp.Code, resp.Message)
	fmt.Println("项目已删除")
}

// GetProjectCostProtectStatus queries project cost protection (v3).
// API: 11-013 ProjectCostProtectStatusGetV30Api
func GetProjectCostProtectStatus(c *ad_open_sdk_go.Client, accessToken string, advID int64, projectIDs []int64) {
	ctx := context.Background()
	resp, _, err := c.ProjectCostProtectStatusGetV30Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(advID).
		ProjectIds(projectIDs).
		Execute()
	if err != nil {
		log.Fatalf("查询项目成本保护状态失败: %v", err)
	}
	checkResp("查询项目成本保护状态", resp.Code, resp.Message)

	if resp.Data == nil {
		fmt.Println("暂无成本保护数据")
		return
	}
	fmt.Println("=== 项目成本保护状态 ===")
	fmt.Printf("%+v\n", resp.Data)
}
