package client

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	ad_open_sdk_go "github.com/oceanengine/ad_open_sdk_go"
	"github.com/oceanengine/ad_open_sdk_go/models"
	"context"
)

type TokenStore struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresAt    time.Time `json:"expires_at"`
	RefreshedAt  time.Time `json:"refreshed_at"`
}

const tokenFile = "token.json"

func tokenFilePath() string {
	// Support override via env var for MCP server launched from arbitrary cwd
	if p := os.Getenv("TOKEN_FILE"); p != "" {
		return p
	}
	// Default: same directory as the running binary
	exe, err := os.Executable()
	if err == nil {
		return exe[:len(exe)-len("/"+filepath.Base(exe))] + "/" + tokenFile
	}
	return tokenFile
}

func LoadToken() (*TokenStore, error) {
	data, err := os.ReadFile(tokenFilePath())
	if err != nil {
		return nil, err
	}
	var t TokenStore
	if err := json.Unmarshal(data, &t); err != nil {
		return nil, err
	}
	return &t, nil
}

func SaveToken(t *TokenStore) error {
	data, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(tokenFilePath(), data, 0600)
}

func (t *TokenStore) IsExpired() bool {
	return time.Now().After(t.ExpiresAt.Add(-5 * time.Minute))
}

func (t *TokenStore) String() string {
	return fmt.Sprintf("token expires at %s (refreshed %s)", t.ExpiresAt.Format(time.DateTime), t.RefreshedAt.Format(time.DateTime))
}

func NewTokenFromAccessToken(resp *models.Oauth2AccessTokenResponseData) *TokenStore {
	now := time.Now()
	return &TokenStore{
		AccessToken:  *resp.AccessToken,
		RefreshToken: *resp.RefreshToken,
		ExpiresAt:    now.Add(time.Duration(*resp.ExpiresIn) * time.Second),
		RefreshedAt:  now,
	}
}

func NewTokenFromRefresh(resp *models.Oauth2RefreshTokenResponseData) *TokenStore {
	now := time.Now()
	return &TokenStore{
		AccessToken:  *resp.AccessToken,
		RefreshToken: *resp.RefreshToken,
		ExpiresAt:    now.Add(time.Duration(*resp.ExpiresIn) * time.Second),
		RefreshedAt:  now,
	}
}

// GetValidToken loads token, refreshes if expired, saves back.
func GetValidToken(c *ad_open_sdk_go.Client, appID int64, secret string) (*TokenStore, error) {
	token, err := LoadToken()
	if err != nil {
		return nil, fmt.Errorf("no saved token, run 'ocean auth' first: %w", err)
	}
	if !token.IsExpired() {
		return token, nil
	}
	log.Println("token expired, refreshing...")
	newToken, err := refreshAccessToken(c, appID, secret, token)
	if err != nil {
		return nil, fmt.Errorf("refresh failed (may need re-auth): %w", err)
	}
	if err := SaveToken(newToken); err != nil {
		return nil, fmt.Errorf("save refreshed token: %w", err)
	}
	log.Println("token refreshed:", newToken)
	return newToken, nil
}

func refreshAccessToken(c *ad_open_sdk_go.Client, appID int64, secret string, token *TokenStore) (*TokenStore, error) {
	ctx := context.Background()
	req := models.Oauth2RefreshTokenRequest{
		AppId:        models.PtrInt64(appID),
		Secret:       secret,
		RefreshToken: token.RefreshToken,
	}
	resp, _, err := c.Oauth2RefreshTokenApi().
		Post(ctx).
		Oauth2RefreshTokenRequest(req).
		Execute()
	if err != nil {
		return nil, fmt.Errorf("refresh token: %w", err)
	}
	if resp.Code != nil && *resp.Code != 0 {
		return nil, fmt.Errorf("refresh token: code=%d msg=%s", *resp.Code, ptrStr(resp.Message))
	}
	if resp.Data == nil || resp.Data.AccessToken == nil {
		return nil, fmt.Errorf("refresh token: empty response")
	}
	return NewTokenFromRefresh(resp.Data), nil
}

func ptrStr(p *string) string {
	if p == nil {
		return ""
	}
	return *p
}
