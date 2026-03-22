package format

import (
	"fmt"
	"sort"
	"strings"
)

// MetricNames maps API field names to Chinese display names.
var MetricNames = map[string]string{
	"stat_cost": "消耗(元)", "show_cnt": "展示数", "cpm_platform": "千次展现费用(元)",
	"click_cnt": "点击数", "ctr": "点击率", "cpc_platform": "点击单价(元)",
	"convert_cnt": "转化数", "conversion_cost": "转化成本", "conversion_rate": "转化率",
	"deep_convert_cnt": "深度转化数", "deep_convert_cost": "深度转化成本", "deep_convert_rate": "深度转化率",
	"attribution_convert_cnt": "转化数(计费)", "attribution_convert_cost": "转化成本(计费)", "attribution_conversion_rate": "转化率(计费)",
	"attribution_deep_convert_cnt": "深度转化数(计费)", "attribution_deep_convert_cost": "深度转化成本(计费)", "attribution_deep_convert_rate": "深度转化率(计费)",
	"total_play": "播放量", "valid_play": "有效播放", "valid_play_cost": "有效播放成本", "valid_play_rate": "有效播放率",
	"play_25_feed_break": "25%播放", "play_50_feed_break": "50%播放", "play_75_feed_break": "75%播放", "play_99_feed_break": "99%播放",
	"average_play_time_per_play": "平均播放时长", "play_over_rate": "完播率",
	"dy_like": "点赞", "dy_comment": "评论", "dy_share": "分享", "dy_follow": "新增粉丝", "dy_home_visited": "主页访问",
	"click_landing_page": "落地页访问", "form": "表单提交", "form_submit": "附加表单提交",
	"form_and_submit_count": "表单总提交数", "form_and_submit_cost": "表单提交成本",
	"phone": "电话点击", "phone_confirm": "电话确认拨打", "phone_connect": "电话确认接通", "phone_effective": "电话有效接通",
	"consult": "在线咨询", "consult_effective": "有效咨询", "consult_clue": "留咨咨询",
	"message_action": "私信开口", "message_enter_chat": "进入私信", "clue_message_count": "私信留资",
	"attribution_all_convert_clue_count": "线索数(计费)", "attribution_all_convert_clue_cost": "线索成本(计费)",
	"customer_effective": "有效获客", "attribution_customer_effective": "有效获客(计费)", "attribution_customer_effective_cost": "有效获客成本(计费)",
	"clue_dialed_count": "拨打线索数", "clue_connected_rate": "接通率", "clue_connected_cost": "接通成本",
	"clue_connected_30s_count": "30秒沟通数", "clue_connected_30s_rate": "30秒沟通率", "clue_connected_average_duration": "平均通话时长(秒)",
	"luban_live_enter_cnt": "直播观看", "live_watch_one_minute_count": "直播>1分钟",
	"luban_live_follow_cnt": "直播关注", "luban_live_comment_cnt": "直播评论",
	"search_after_read_pv": "看后搜次数", "search_after_pv_rate": "看后搜率",
}

func isPercentField(key string) bool {
	return strings.Contains(key, "rate") || strings.Contains(key, "ctr") || strings.Contains(key, "roi")
}

func isZero(val string) bool {
	return val == "" || val == "0" || val == "0.00" || val == "0.0"
}

func displayName(key string) string {
	if n, ok := MetricNames[key]; ok {
		return n
	}
	return key
}

// PrintMetricsOrdered writes metrics in the given order to b, skipping zeros.
func PrintMetricsOrdered(b *strings.Builder, metrics map[string]string, order []string) {
	for _, key := range order {
		val, ok := metrics[key]
		if !ok || isZero(val) {
			continue
		}
		if isPercentField(key) {
			fmt.Fprintf(b, "  %-24s %s%%\n", displayName(key), val)
		} else {
			fmt.Fprintf(b, "  %-24s %s\n", displayName(key), val)
		}
	}
}

// PrintMetricsSorted writes all non-zero metrics sorted alphabetically to b.
func PrintMetricsSorted(b *strings.Builder, m map[string]string) {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, key := range keys {
		val := m[key]
		if isZero(val) {
			continue
		}
		if isPercentField(key) {
			fmt.Fprintf(b, "  %-24s %s%%\n", displayName(key), val)
		} else {
			fmt.Fprintf(b, "  %-24s %s\n", displayName(key), val)
		}
	}
}

// Separator writes a horizontal line to b.
func Separator(b *strings.Builder) {
	b.WriteString(strings.Repeat("-", 60) + "\n")
}
