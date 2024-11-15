package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"com.goa/internal/config"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

var (
	// 随机生成的 OAuth 状态字符串
	oauthStateString = "TODO_GEN" // 在生产环境中应使用安全的随机字符串
)

func getGithubOAuthConfig() oauth2.Config {
	return oauth2.Config{
		ClientID:     config.C.Oauth.Github.ClientId,         // GitHub OAuth 应用的 Client ID
		ClientSecret: config.C.Oauth.Github.ClientSecret,     // GitHub OAuth 应用的 Client Secret
		RedirectURL:  config.C.Oauth.Github.RedirectUrl,      // 回调 URL（用于本地测试）
		Scopes:       []string{config.C.Oauth.Github.Scopes}, // 请求 GitHub 用户邮箱权限
		Endpoint:     github.Endpoint,                        // GitHub OAuth 端点
	}
}

func GithubOauth(c *gin.Context) {
	githubOauthConfig := getGithubOAuthConfig()

	url := githubOauthConfig.AuthCodeURL(oauthStateString, oauth2.AccessTypeOffline)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func GithubOauthCallback(c *gin.Context) {
	githubOauthConfig := getGithubOAuthConfig()
	code := c.Query("code")

	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Code not found"})
	}

	// 使用授权码换取 access token
	token, err := githubOauthConfig.Exchange(c.Request.Context(), code)
	if err != nil {
		// code 无效
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to exchange token: %s", err.Error())})
		return
	}

	// 使用获取的 access token 访问 GitHub 用户信息
	client := githubOauthConfig.Client(c.Request.Context(), token)
	resp, err := client.Get("https://api.github.com/user")
	if err != nil {
		// 获取失败
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to get user info: %s", err.Error())})
		return
	}
	defer resp.Body.Close()

	// 解析 GitHub 返回的用户信息
	var userInfo struct {
		ID        int64  `json:"id"`
		Login     string `json:"login"`
		Email     string `json:"email"`
		AvatarUrl string `json:"avatar_url"`
		Bio       string `json:"bio"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		// json解析失败
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to parse user info: %s", err.Error())})
		return
	}

	// 返回用户信息
	c.JSON(http.StatusOK, userInfo)
}
