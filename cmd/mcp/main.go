package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	"github.com/andyleimc-source/oceanengine-mcp/internal/api"
	"github.com/andyleimc-source/oceanengine-mcp/internal/auth"
	"github.com/andyleimc-source/oceanengine-mcp/internal/client"
)

func main() {
	// 加载 .env
	_ = godotenv.Load()

	appIDStr := os.Getenv("APP_ID")
	secret := os.Getenv("APP_SECRET")
	if appIDStr == "" || secret == "" {
		log.Fatal("请在 .env 中设置 APP_ID 和 APP_SECRET")
	}
	appID, err := strconv.ParseInt(appIDStr, 10, 64)
	if err != nil {
		log.Fatalf("APP_ID 格式错误: %v", err)
	}

	var defaultAdvID int64
	if advIDStr := os.Getenv("ADVERTISER_ID"); advIDStr != "" {
		defaultAdvID, err = strconv.ParseInt(advIDStr, 10, 64)
		if err != nil {
			log.Fatalf("ADVERTISER_ID 格式错误: %v", err)
		}
	}

	// 初始化 SDK 客户端
	c := client.New()

	// 加载并自动刷新 token（需要先运行 ocean auth）
	token, err := client.GetValidToken(c, appID, secret)
	if err != nil {
		log.Fatalf("获取 token 失败，请先运行 ocean auth: %v", err)
	}
	accessToken := token.AccessToken

	// ── MCP server ────────────────────────────────────────────────────────────

	s := server.NewMCPServer(
		"oceanengine-mcp",
		"0.1.0",
		server.WithToolCapabilities(true),
	)

	// ── 辅助函数 ──────────────────────────────────────────────────────────────

	// resolveAdvID 优先使用工具参数中的 advertiser_id，否则用 .env 的默认值
	resolveAdvID := func(args map[string]interface{}) (int64, error) {
		if v, ok := args["advertiser_id"]; ok {
			s, ok := v.(string)
			if !ok {
				return 0, fmt.Errorf("advertiser_id 必须为字符串")
			}
			if s != "" {
				id, err := strconv.ParseInt(s, 10, 64)
				if err != nil {
					return 0, fmt.Errorf("advertiser_id 格式错误: %v", err)
				}
				return id, nil
			}
		}
		if defaultAdvID == 0 {
			return 0, fmt.Errorf("请在 .env 中设置 ADVERTISER_ID 或在调用时传入 advertiser_id")
		}
		return defaultAdvID, nil
	}

	parseIDs := func(s string) ([]int64, error) {
		parts := strings.Split(s, ",")
		var ids []int64
		for _, p := range parts {
			p = strings.TrimSpace(p)
			if p == "" {
				continue
			}
			id, err := strconv.ParseInt(p, 10, 64)
			if err != nil {
				return nil, fmt.Errorf("ID 格式错误: %s", p)
			}
			ids = append(ids, id)
		}
		if len(ids) == 0 {
			return nil, fmt.Errorf("请提供至少一个 ID")
		}
		return ids, nil
	}

	// getArgs 将 req.Params.Arguments (any) 转为 map，屏蔽 mcp-go 版本差异
	getArgs := func(raw interface{}) map[string]interface{} {
		if m, ok := raw.(map[string]interface{}); ok {
			return m
		}
		return map[string]interface{}{}
	}

	getString := func(args map[string]interface{}, key string) string {
		if v, ok := args[key]; ok {
			if s, ok := v.(string); ok {
				return s
			}
		}
		return ""
	}

	toolResult := func(result string, err error) (*mcp.CallToolResult, error) {
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return mcp.NewToolResultText(result), nil
	}

	// ── 账户 ──────────────────────────────────────────────────────────────────

	s.AddTool(
		mcp.NewTool("list_accounts",
			mcp.WithDescription("查询已授权的巨量引擎广告账户列表"),
		),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			return toolResult(auth.ListAccounts(c, accessToken))
		},
	)

	// ── 报表 ──────────────────────────────────────────────────────────────────

	s.AddTool(
		mcp.NewTool("get_report",
			mcp.WithDescription("拉取巨量引擎广告报表数据，支持广告主/项目/单元三个维度"),
			mcp.WithString("start_date",
				mcp.Description("开始日期，格式 2006-01-02，默认近7天"),
			),
			mcp.WithString("end_date",
				mcp.Description("结束日期，格式 2006-01-02，默认今天"),
			),
			mcp.WithString("level",
				mcp.Description("报表维度：advertiser（广告主汇总）、campaign（项目）、ad（广告单元），默认 advertiser"),
			),
			mcp.WithString("advertiser_id",
				mcp.Description("广告主ID，不填则使用 .env 中的 ADVERTISER_ID"),
			),
		),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := getArgs(req.Params.Arguments)
			advID, err := resolveAdvID(args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			endDate := time.Now().Format("2006-01-02")
			startDate := time.Now().AddDate(0, 0, -7).Format("2006-01-02")
			if v := getString(args, "start_date"); v != "" {
				startDate = v
			}
			if v := getString(args, "end_date"); v != "" {
				endDate = v
			}
			level := api.LevelAdvertiser
			switch getString(args, "level") {
			case "campaign", "project":
				level = api.LevelProject
			case "ad", "promotion":
				level = api.LevelPromotion
			}
			return toolResult(api.FetchReport(api.ReportParams{
				Client:       c,
				AccessToken:  accessToken,
				AdvertiserID: advID,
				StartDate:    startDate,
				EndDate:      endDate,
				Level:        level,
			}))
		},
	)

	s.AddTool(
		mcp.NewTool("get_report_config",
			mcp.WithDescription("查询巨量引擎报表可用的指标（metrics）和维度（dimensions）列表"),
			mcp.WithString("advertiser_id",
				mcp.Description("广告主ID，不填则使用 .env 中的 ADVERTISER_ID"),
			),
		),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			advID, err := resolveAdvID(getArgs(req.Params.Arguments))
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			return toolResult(api.FetchReportConfig(c, accessToken, advID))
		},
	)

	// ── v2 广告组 Campaign ────────────────────────────────────────────────────

	s.AddTool(
		mcp.NewTool("list_campaigns",
			mcp.WithDescription("查询 v2 广告组列表"),
			mcp.WithString("advertiser_id",
				mcp.Description("广告主ID，不填则使用 .env 中的 ADVERTISER_ID"),
			),
		),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			advID, err := resolveAdvID(getArgs(req.Params.Arguments))
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			return toolResult(api.ListCampaigns(c, accessToken, advID))
		},
	)

	s.AddTool(
		mcp.NewTool("update_campaign_status",
			mcp.WithDescription("批量更新 v2 广告组状态（启用/暂停/删除）"),
			mcp.WithString("campaign_ids",
				mcp.Required(),
				mcp.Description("广告组ID列表，逗号分隔，如 123,456"),
			),
			mcp.WithString("status",
				mcp.Required(),
				mcp.Description("目标状态：enable（启用）、disable（暂停）、delete（删除）"),
			),
			mcp.WithString("advertiser_id",
				mcp.Description("广告主ID，不填则使用 .env 中的 ADVERTISER_ID"),
			),
		),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := getArgs(req.Params.Arguments)
			advID, err := resolveAdvID(args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			ids, err := parseIDs(getString(args, "campaign_ids"))
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			return toolResult(api.UpdateCampaignStatus(c, accessToken, advID, ids, getString(args, "status")))
		},
	)

	// ── v2 广告计划 Ad ────────────────────────────────────────────────────────

	s.AddTool(
		mcp.NewTool("list_ads",
			mcp.WithDescription("查询 v2 广告计划列表（注：SDK 限制，当前仅返回单条）"),
			mcp.WithString("advertiser_id",
				mcp.Description("广告主ID，不填则使用 .env 中的 ADVERTISER_ID"),
			),
		),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			advID, err := resolveAdvID(getArgs(req.Params.Arguments))
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			return toolResult(api.ListAds(c, accessToken, advID))
		},
	)

	/* update_ad_bid disabled
	s.AddTool(
		mcp.NewTool("update_ad_bid", ...),
	)
	*/

	s.AddTool(
		mcp.NewTool("update_ad_budget",
			mcp.WithDescription("更新 v2 广告计划预算"),
			mcp.WithString("ad_id",
				mcp.Required(),
				mcp.Description("广告计划ID"),
			),
			mcp.WithNumber("budget",
				mcp.Required(),
				mcp.Description("新预算金额（元）"),
			),
			mcp.WithString("advertiser_id",
				mcp.Description("广告主ID，不填则使用 .env 中的 ADVERTISER_ID"),
			),
		),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := getArgs(req.Params.Arguments)
			advID, err := resolveAdvID(args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			adID, err := strconv.ParseInt(getString(args, "ad_id"), 10, 64)
			if err != nil {
				return mcp.NewToolResultError("ad_id 格式错误"), nil
			}
			budget, _ := args["budget"].(float64)
			return toolResult(api.UpdateAdBudget(c, accessToken, advID, adID, budget))
		},
	)

	/* get_ad_reject_reason disabled */
	/* get_ad_cost_protect_status disabled */

	// ── v3 项目 Project ───────────────────────────────────────────────────────

	/* create_project disabled
	s.AddTool(
		mcp.NewTool("create_project",
			mcp.WithDescription("创建 v3 广告项目"),
			mcp.WithString("name",
				mcp.Required(),
				mcp.Description("项目名称"),
			),
			mcp.WithNumber("budget",
				mcp.Description("项目预算（元），不填表示不限预算"),
			),
			mcp.WithString("budget_mode",
				mcp.Description("预算类型：day（日预算）、infinite（不限），默认 infinite"),
			),
			mcp.WithString("landing_type",
				mcp.Description("推广目的：LINK（落地页）、APP（应用推广）、SHOP（商品推广）等，默认 LINK"),
			),
			mcp.WithString("start_time",
				mcp.Description("投放开始时间，格式 2006-01-02"),
			),
			mcp.WithString("end_time",
				mcp.Description("投放结束时间，格式 2006-01-02"),
			),
			mcp.WithString("advertiser_id",
				mcp.Description("广告主ID，不填则使用 .env 中的 ADVERTISER_ID"),
			),
		),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := getArgs(req.Params.Arguments)
			advID, err := resolveAdvID(args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			name := getString(args, "name")
			if name == "" {
				return mcp.NewToolResultError("name 不能为空"), nil
			}
			budget, _ := args["budget"].(float64)
			budgetMode := getString(args, "budget_mode")
			if budgetMode == "" {
				budgetMode = "infinite"
			}
			landingType := getString(args, "landing_type")
			if landingType == "" {
				landingType = "LINK"
			}
			return toolResult(api.CreateProject(c, accessToken, advID, name, budget, budgetMode, landingType, getString(args, "start_time"), getString(args, "end_time")))
		},
	)
	*/

	s.AddTool(
		mcp.NewTool("list_projects",
			mcp.WithDescription("查询 v3 项目列表"),
			mcp.WithString("advertiser_id",
				mcp.Description("广告主ID，不填则使用 .env 中的 ADVERTISER_ID"),
			),
		),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			advID, err := resolveAdvID(getArgs(req.Params.Arguments))
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			return toolResult(api.ListProjects(c, accessToken, advID))
		},
	)

	s.AddTool(
		mcp.NewTool("update_project_status",
			mcp.WithDescription("批量更新 v3 项目状态（启用/暂停）"),
			mcp.WithString("project_ids",
				mcp.Required(),
				mcp.Description("项目ID列表，逗号分隔，如 123,456"),
			),
			mcp.WithString("status",
				mcp.Required(),
				mcp.Description("目标状态：enable（启用）、disable（暂停）"),
			),
			mcp.WithString("advertiser_id",
				mcp.Description("广告主ID，不填则使用 .env 中的 ADVERTISER_ID"),
			),
		),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := getArgs(req.Params.Arguments)
			advID, err := resolveAdvID(args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			ids, err := parseIDs(getString(args, "project_ids"))
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			return toolResult(api.UpdateProjectStatus(c, accessToken, advID, ids, getString(args, "status")))
		},
	)

	s.AddTool(
		mcp.NewTool("update_project_budget",
			mcp.WithDescription("更新 v3 项目预算"),
			mcp.WithString("project_id",
				mcp.Required(),
				mcp.Description("项目ID"),
			),
			mcp.WithNumber("budget",
				mcp.Required(),
				mcp.Description("新预算金额（元）"),
			),
			mcp.WithString("budget_mode",
				mcp.Required(),
				mcp.Description("预算模式：day（日预算）、infinite（不限预算）"),
			),
			mcp.WithString("advertiser_id",
				mcp.Description("广告主ID，不填则使用 .env 中的 ADVERTISER_ID"),
			),
		),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := getArgs(req.Params.Arguments)
			advID, err := resolveAdvID(args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			projectID, err := strconv.ParseInt(getString(args, "project_id"), 10, 64)
			if err != nil {
				return mcp.NewToolResultError("project_id 格式错误"), nil
			}
			budget, _ := args["budget"].(float64)
			return toolResult(api.UpdateProjectBudget(c, accessToken, advID, projectID, budget, getString(args, "budget_mode")))
		},
	)

	/* delete_projects disabled */
	/* get_project_cost_protect_status disabled */

	// ── v3 广告单元 Promotion ─────────────────────────────────────────────────

	/* create_promotion disabled
	s.AddTool(
		mcp.NewTool("create_promotion",
			mcp.WithDescription("创建 v3 广告单元"),
			mcp.WithString("project_id",
				mcp.Required(),
				mcp.Description("所属项目ID"),
			),
			mcp.WithString("name",
				mcp.Required(),
				mcp.Description("单元名称"),
			),
			mcp.WithNumber("budget",
				mcp.Description("单元预算（元），不填表示不限预算"),
			),
			mcp.WithString("budget_mode",
				mcp.Description("预算类型：day（日预算）、infinite（不限），默认 infinite"),
			),
			mcp.WithNumber("bid",
				mcp.Description("出价（元），bid_type 为 CUSTOM 时必填"),
			),
			mcp.WithString("bid_type",
				mcp.Description("出价方式：CUSTOM（手动出价）、NO_BID（自动出价），默认 NO_BID"),
			),
			mcp.WithString("audience_mode",
				mcp.Description("定向模式：CUSTOM（自定义）、AUTO（广定向），默认 AUTO"),
			),
			mcp.WithString("start_time",
				mcp.Description("投放开始时间，格式 2006-01-02"),
			),
			mcp.WithString("end_time",
				mcp.Description("投放结束时间，格式 2006-01-02"),
			),
			mcp.WithString("advertiser_id",
				mcp.Description("广告主ID，不填则使用 .env 中的 ADVERTISER_ID"),
			),
		),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := getArgs(req.Params.Arguments)
			advID, err := resolveAdvID(args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			projectID, err := strconv.ParseInt(getString(args, "project_id"), 10, 64)
			if err != nil {
				return mcp.NewToolResultError("project_id 格式错误"), nil
			}
			name := getString(args, "name")
			if name == "" {
				return mcp.NewToolResultError("name 不能为空"), nil
			}
			budget, _ := args["budget"].(float64)
			budgetMode := getString(args, "budget_mode")
			if budgetMode == "" {
				budgetMode = "infinite"
			}
			bid, _ := args["bid"].(float64)
			bidType := getString(args, "bid_type")
			audienceMode := getString(args, "audience_mode")
			return toolResult(api.CreatePromotion(c, accessToken, advID, projectID, name, budget, budgetMode, bid, bidType, audienceMode, getString(args, "start_time"), getString(args, "end_time")))
		},
	)
	*/

	s.AddTool(
		mcp.NewTool("list_promotions",
			mcp.WithDescription("查询 v3 广告单元列表"),
			mcp.WithString("advertiser_id",
				mcp.Description("广告主ID，不填则使用 .env 中的 ADVERTISER_ID"),
			),
		),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			advID, err := resolveAdvID(getArgs(req.Params.Arguments))
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			return toolResult(api.ListPromotions(c, accessToken, advID))
		},
	)

	s.AddTool(
		mcp.NewTool("update_promotion_status",
			mcp.WithDescription("批量更新 v3 广告单元状态（启用/暂停）"),
			mcp.WithString("promotion_ids",
				mcp.Required(),
				mcp.Description("广告单元ID列表，逗号分隔，如 123,456"),
			),
			mcp.WithString("status",
				mcp.Required(),
				mcp.Description("目标状态：enable（启用）、disable（暂停）"),
			),
			mcp.WithString("advertiser_id",
				mcp.Description("广告主ID，不填则使用 .env 中的 ADVERTISER_ID"),
			),
		),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := getArgs(req.Params.Arguments)
			advID, err := resolveAdvID(args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			ids, err := parseIDs(getString(args, "promotion_ids"))
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			return toolResult(api.UpdatePromotionStatus(c, accessToken, advID, ids, getString(args, "status")))
		},
	)

	/* update_promotion_bid disabled */

	s.AddTool(
		mcp.NewTool("update_promotion_budget",
			mcp.WithDescription("更新 v3 广告单元预算"),
			mcp.WithString("promotion_id",
				mcp.Required(),
				mcp.Description("广告单元ID"),
			),
			mcp.WithNumber("budget",
				mcp.Required(),
				mcp.Description("新预算金额（元）"),
			),
			mcp.WithString("advertiser_id",
				mcp.Description("广告主ID，不填则使用 .env 中的 ADVERTISER_ID"),
			),
		),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := getArgs(req.Params.Arguments)
			advID, err := resolveAdvID(args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			promotionID, err := strconv.ParseInt(getString(args, "promotion_id"), 10, 64)
			if err != nil {
				return mcp.NewToolResultError("promotion_id 格式错误"), nil
			}
			budget, _ := args["budget"].(float64)
			return toolResult(api.UpdatePromotionBudget(c, accessToken, advID, promotionID, budget))
		},
	)

	/* delete_promotions disabled */
	/* get_promotion_reject_reason disabled */
	/* get_promotion_cost_protect_status disabled */

	// ── 创意 Creative ─────────────────────────────────────────────────────────

	s.AddTool(
		mcp.NewTool("list_creatives",
			mcp.WithDescription("查询 v2 创意列表"),
			mcp.WithString("advertiser_id",
				mcp.Description("广告主ID，不填则使用 .env 中的 ADVERTISER_ID"),
			),
		),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			advID, err := resolveAdvID(getArgs(req.Params.Arguments))
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			return toolResult(api.ListCreatives(c, accessToken, advID))
		},
	)

	s.AddTool(
		mcp.NewTool("get_creative_detail",
			mcp.WithDescription("查询 v3 创意详情"),
			mcp.WithString("ad_id",
				mcp.Required(),
				mcp.Description("广告计划ID"),
			),
			mcp.WithString("advertiser_id",
				mcp.Description("广告主ID，不填则使用 .env 中的 ADVERTISER_ID"),
			),
		),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := getArgs(req.Params.Arguments)
			advID, err := resolveAdvID(args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			adID, err := strconv.ParseInt(getString(args, "ad_id"), 10, 64)
			if err != nil {
				return mcp.NewToolResultError("ad_id 格式错误"), nil
			}
			return toolResult(api.GetCreativeDetail(c, accessToken, advID, adID))
		},
	)

	/* get_creative_reject_reason disabled */

	// ── 启动 stdio server ──────────────────────────────────────────────────────
	if err := server.ServeStdio(s); err != nil {
		log.Fatalf("MCP server error: %v", err)
	}
}
