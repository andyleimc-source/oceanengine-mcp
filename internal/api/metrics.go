package api

// AllMetrics defines the full set of metrics to fetch for reports.
var AllMetrics = []string{
	// 基础消耗
	"stat_cost", "show_cnt", "cpm_platform", "click_cnt", "ctr", "cpc_platform",
	// 转化（回传时间）
	"convert_cnt", "conversion_cost", "conversion_rate",
	"deep_convert_cnt", "deep_convert_cost", "deep_convert_rate",
	// 转化（计费时间）
	"attribution_convert_cnt", "attribution_convert_cost", "attribution_conversion_rate",
	"attribution_deep_convert_cnt", "attribution_deep_convert_cost", "attribution_deep_convert_rate",
	// 视频互动
	"total_play", "valid_play", "valid_play_cost", "valid_play_rate",
	"play_25_feed_break", "play_50_feed_break", "play_75_feed_break", "play_99_feed_break",
	"average_play_time_per_play", "play_over_rate",
	// 社交互动
	"dy_like", "dy_comment", "dy_share", "dy_follow", "dy_home_visited",
	// 落地页 & 表单
	"click_landing_page", "form", "form_submit", "form_and_submit_count", "form_and_submit_cost",
	"phone", "phone_confirm", "phone_connect", "phone_effective",
	"consult", "consult_effective", "consult_clue",
	// 私信
	"message_action", "message_enter_chat", "clue_message_count",
	// 线索
	"attribution_all_convert_clue_count", "attribution_all_convert_clue_cost",
	"customer_effective", "attribution_customer_effective", "attribution_customer_effective_cost",
	// 电话线索
	"clue_dialed_count", "clue_connected_rate", "clue_connected_cost",
	"clue_connected_30s_count", "clue_connected_30s_rate", "clue_connected_average_duration",
	// 直播
	"luban_live_enter_cnt", "live_watch_one_minute_count",
	"luban_live_follow_cnt", "luban_live_comment_cnt",
	// 搜索
	"search_after_read_pv", "search_after_pv_rate",
}

// CoreMetrics is a smaller set for campaign/promotion level reports.
var CoreMetrics = []string{
	"stat_cost", "show_cnt", "click_cnt", "ctr", "cpc_platform",
	"convert_cnt", "conversion_cost", "conversion_rate",
}
