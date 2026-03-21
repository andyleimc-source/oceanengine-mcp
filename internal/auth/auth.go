package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"

	ad_open_sdk_go "github.com/oceanengine/ad_open_sdk_go"
	"github.com/oceanengine/ad_open_sdk_go/models"
	"ocean/internal/client"
)

func buildAuthURL(appID string, redirectURI string) string {
	params := url.Values{
		"app_id":       {appID},
		"redirect_uri": {redirectURI},
		"state":        {"ocean_report"},
	}
	return "https://open.oceanengine.com/audit/oauth.html?" + params.Encode()
}

func exchangeToken(c *ad_open_sdk_go.Client, appID int64, secret, authCode string) (*client.TokenStore, error) {
	ctx := context.Background()
	req := models.Oauth2AccessTokenRequest{
		AppId:    models.PtrInt64(appID),
		Secret:   secret,
		AuthCode: authCode,
	}
	resp, _, err := c.Oauth2AccessTokenApi().
		Post(ctx).
		Oauth2AccessTokenRequest(req).
		Execute()
	if err != nil {
		return nil, fmt.Errorf("exchange token: %w", err)
	}
	if resp.Code != nil && *resp.Code != 0 {
		msg := ""
		if resp.Message != nil {
			msg = *resp.Message
		}
		return nil, fmt.Errorf("exchange token failed: code=%d msg=%s", *resp.Code, msg)
	}
	if resp.Data == nil || resp.Data.AccessToken == nil {
		return nil, fmt.Errorf("exchange token: empty response data")
	}

	if len(resp.Data.AdvertiserIds) > 0 {
		fmt.Println("已授权的广告主 ID:")
		for _, id := range resp.Data.AdvertiserIds {
			fmt.Printf("  %d\n", id)
		}
	}

	return client.NewTokenFromAccessToken(resp.Data), nil
}

// StartAuthServer starts local OAuth2 callback server.
func StartAuthServer(c *ad_open_sdk_go.Client, appID int64, secret string) {
	redirectURI := "http://localhost:9527/callback"
	authURL := buildAuthURL(fmt.Sprintf("%d", appID), redirectURI)

	fmt.Println("=== OAuth2 授权 ===")
	fmt.Println("请在浏览器打开以下链接完成授权:")
	fmt.Println()
	fmt.Println(authURL)
	fmt.Println()
	fmt.Println("等待回调中... (监听 localhost:9527)")

	done := make(chan struct{})

	mux := http.NewServeMux()
	mux.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		authCode := r.URL.Query().Get("auth_code")
		if authCode == "" {
			http.Error(w, "missing auth_code", http.StatusBadRequest)
			return
		}
		token, err := exchangeToken(c, appID, secret, authCode)
		if err != nil {
			errMsg := fmt.Sprintf("换取 token 失败: %v", err)
			http.Error(w, errMsg, http.StatusInternalServerError)
			log.Println(errMsg)
			return
		}
		if err := client.SaveToken(token); err != nil {
			log.Printf("保存 token 失败: %v", err)
		}
		fmt.Fprintf(w, "授权成功! token 已保存，可以关闭此页面。")
		fmt.Println()
		fmt.Println("授权成功!", token)
		close(done)
	})

	server := &http.Server{Addr: ":9527", Handler: mux}
	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("auth server: %v", err)
		}
	}()
	<-done
	server.Close()
}

// ListAccounts prints authorized accounts.
func ListAccounts(c *ad_open_sdk_go.Client, accessToken string) {
	ctx := context.Background()
	resp, _, err := c.Oauth2AdvertiserGetApi().
		Get(ctx).
		AccessToken(accessToken).
		Execute()
	if err != nil {
		log.Fatalf("查询授权账户失败: %v", err)
	}
	if resp.Code != nil && *resp.Code != 0 {
		msg := ""
		if resp.Message != nil {
			msg = *resp.Message
		}
		log.Fatalf("查询授权账户错误: code=%d msg=%s", *resp.Code, msg)
	}
	if resp.Data == nil || len(resp.Data.List) == 0 {
		fmt.Println("没有已授权的账户")
		return
	}

	fmt.Println("=== 已授权账户列表 ===")
	fmt.Println()
	for _, item := range resp.Data.List {
		var id int64
		if item.AccountId != nil {
			id = *item.AccountId
		}
		name := ""
		if item.AccountName != nil {
			name = *item.AccountName
		}
		accType := "-"
		if item.AccountType != nil {
			accType = string(*item.AccountType)
		}
		valid := "?"
		if item.IsValid != nil {
			if *item.IsValid {
				valid = "有效"
			} else {
				valid = "无效"
			}
		}
		fmt.Printf("  ID: %d\n  名称: %s\n  类型: %s\n  状态: %s\n", id, name, accType, valid)
		fmt.Println("  ----------")
	}
	fmt.Println()
	fmt.Println("请将类型为 ADVERTISER 的账户 ID 填入 .env 的 ADVERTISER_ID")
}
