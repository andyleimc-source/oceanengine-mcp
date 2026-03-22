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

	s.AddTool(
		mcp.NewTool("update_ad_bid",
			mcp.WithDescription("更新 v2 广告计划出价"),
			mcp.WithString("ad_id",
				mcp.Required(),
				mcp.Description("广告计划ID"),
			),
			mcp.WithNumber("bid",
				mcp.Required(),
				mcp.Description("新出价金额（元）"),
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
			bid, _ := args["bid"].(float64)
			return toolResult(api.UpdateAdBid(c, accessToken, advID, adID, bid))
		},
	)

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

	s.AddTool(
		mcp.NewTool("get_ad_reject_reason",
			mcp.WithDescription("查询 v2 广告计划审核拒绝原因"),
			mcp.WithString("ad_ids",
				mcp.Required(),
				mcp.Description("广告计划ID列表，逗号分隔，如 123,456"),
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
			ids, err := parseIDs(getString(args, "ad_ids"))
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			return toolResult(api.GetAdRejectReason(c, accessToken, advID, ids))
		},
	)

	s.AddTool(
		mcp.NewTool("get_ad_cost_protect_status",
			mcp.WithDescription("查询 v2 广告计划成本保护状态"),
			mcp.WithString("ad_ids",
				mcp.Required(),
				mcp.Description("广告计划ID列表，逗号分隔，如 123,456"),
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
			ids, err := parseIDs(getString(args, "ad_ids"))
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			return toolResult(api.GetAdCostProtectStatus(c, accessToken, advID, ids))
		},
	)

	// ── v3 项目 Project ───────────────────────────────────────────────────────

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

	s.AddTool(
		mcp.NewTool("delete_projects",
			mcp.WithDescription("删除 v3 项目"),
			mcp.WithString("project_ids",
				mcp.Required(),
				mcp.Description("项目ID列表，逗号分隔，如 123,456"),
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
			return toolResult(api.DeleteProjects(c, accessToken, advID, ids))
		},
	)

	s.AddTool(
		mcp.NewTool("get_project_cost_protect_status",
			mcp.WithDescription("查询 v3 项目成本保护状态"),
			mcp.WithString("project_ids",
				mcp.Required(),
				mcp.Description("项目ID列表，逗号分隔，如 123,456"),
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
			return toolResult(api.GetProjectCostProtectStatus(c, accessToken, advID, ids))
		},
	)

	// ── v3 广告单元 Promotion ─────────────────────────────────────────────────

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

	s.AddTool(
		mcp.NewTool("update_promotion_bid",
			mcp.WithDescription("更新 v3 广告单元出价"),
			mcp.WithString("promotion_id",
				mcp.Required(),
				mcp.Description("广告单元ID"),
			),
			mcp.WithNumber("bid",
				mcp.Required(),
				mcp.Description("新出价金额（元）"),
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
			bid, _ := args["bid"].(float64)
			return toolResult(api.UpdatePromotionBid(c, accessToken, advID, promotionID, bid))
		},
	)

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

	s.AddTool(
		mcp.NewTool("delete_promotions",
			mcp.WithDescription("删除 v3 广告单元"),
			mcp.WithString("promotion_ids",
				mcp.Required(),
				mcp.Description("广告单元ID列表，逗号分隔，如 123,456"),
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
			return toolResult(api.DeletePromotions(c, accessToken, advID, ids))
		},
	)

	s.AddTool(
		mcp.NewTool("get_promotion_reject_reason",
			mcp.WithDescription("查询 v3 广告单元审核拒绝原因"),
			mcp.WithString("promotion_ids",
				mcp.Required(),
				mcp.Description("广告单元ID列表，逗号分隔，如 123,456"),
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
			return toolResult(api.GetPromotionRejectReason(c, accessToken, advID, ids))
		},
	)

	s.AddTool(
		mcp.NewTool("get_promotion_cost_protect_status",
			mcp.WithDescription("查询 v3 广告单元成本保护状态"),
			mcp.WithString("promotion_ids",
				mcp.Required(),
				mcp.Description("广告单元ID列表，逗号分隔，如 123,456"),
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
			return toolResult(api.GetPromotionCostProtectStatus(c, accessToken, advID, ids))
		},
	)

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

	s.AddTool(
		mcp.NewTool("get_creative_reject_reason",
			mcp.WithDescription("查询 v2 创意审核拒绝原因"),
			mcp.WithString("creative_ids",
				mcp.Required(),
				mcp.Description("创意ID列表，逗号分隔，如 123,456"),
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
			ids, err := parseIDs(getString(args, "creative_ids"))
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			return toolResult(api.GetCreativeRejectReason(c, accessToken, advID, ids))
		},
	)

	// ── 启动 stdio server ──────────────────────────────────────────────────────
	if err := server.ServeStdio(s); err != nil {
		log.Fatalf("MCP server error: %v", err)
	}
}
