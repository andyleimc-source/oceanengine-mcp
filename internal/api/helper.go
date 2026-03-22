package api

import (
	"fmt"
	"reflect"
	"strings"
)

// CheckResp checks API response code and returns an error if non-zero.
func CheckResp(action string, code *int64, message *string) error {
	if code != nil && *code != 0 {
		msg := ""
		if message != nil {
			msg = *message
		}
		return fmt.Errorf("%s错误: code=%d msg=%s", action, *code, msg)
	}
	return nil
}

func writeField(b *strings.Builder, label string, val interface{}) {
	if val == nil {
		return
	}

	rv := reflect.ValueOf(val)

	// Handle pointer types (including *CustomEnumType which is really *string underneath)
	if rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			return
		}
		elem := rv.Elem()
		switch elem.Kind() {
		case reflect.Int64:
			fmt.Fprintf(b, "  %-16s %d\n", label, elem.Int())
		case reflect.Float64:
			fmt.Fprintf(b, "  %-16s %.2f\n", label, elem.Float())
		case reflect.String:
			s := elem.String()
			if s != "" {
				fmt.Fprintf(b, "  %-16s %s\n", label, s)
			}
		default:
			fmt.Fprintf(b, "  %-16s %v\n", label, elem.Interface())
		}
		return
	}

	// Handle value types
	switch rv.Kind() {
	case reflect.Int64:
		fmt.Fprintf(b, "  %-16s %d\n", label, rv.Int())
	case reflect.Float64:
		fmt.Fprintf(b, "  %-16s %.2f\n", label, rv.Float())
	case reflect.String:
		s := rv.String()
		if s != "" {
			fmt.Fprintf(b, "  %-16s %s\n", label, s)
		}
	default:
		s := fmt.Sprintf("%v", val)
		if s != "" && s != "<nil>" {
			fmt.Fprintf(b, "  %-16s %s\n", label, s)
		}
	}
}

func writeFieldFloat(b *strings.Builder, label string, val *float64) {
	if val != nil && *val > 0 {
		fmt.Fprintf(b, "  %-16s %.2f\n", label, *val)
	}
}

func writePageInfo(b *strings.Builder, pi interface{}) {
	if pi == nil {
		return
	}
	b.WriteString("\n(使用 SDK 分页获取更多数据)\n")
}
