package internal

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"com.goa/internal/api"
	"com.goa/internal/config"
	"github.com/gin-gonic/gin"
)

func InitEngine(ctx context.Context) (e *gin.Engine, clearFunc func(), err error) {
	// gin.SetMode(gin.ReleaseMode)

	engine := gin.New()

	engine.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hi",
		})
	})

	engine.GET("/github/login", api.GithubOauth)
	engine.GET("/github/callback", api.GithubOauthCallback)

	srv := &http.Server{
		Addr:    config.C.Server.Addr,
		Handler: engine,
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			fmt.Printf("启动http服务失败: %+v", err)
		}
	}()

	return engine, func() {

		ShutdownTimeout := 10
		ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(ShutdownTimeout))

		defer cancel()

		srv.SetKeepAlivesEnabled(false)
		if err := srv.Shutdown(ctx); err != nil {
			fmt.Printf("关闭http服务失败: %+v", err)
		}
	}, nil
}
