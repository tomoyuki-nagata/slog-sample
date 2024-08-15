package api

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"todo-app/api/handler"
	"todo-app/api/middleware"
	"todo-app/registry"

	"github.com/gin-gonic/gin"
)

func Run() {
	engine := gin.New()

	engine.Use(middleware.WithRequestId())
	engine.Use(middleware.WithExecutorId())
	engine.Use(middleware.WithAccessLog())
	engine.Use(middleware.WithCustomGinLogger())
	engine.Use(gin.Recovery())

	// DIコンテナから依存関係を解決したハンドラーを作成し、ルーティング設定を登録する
	container := registry.BuildContainer()
	container.Invoke(func(
		UserHandler *handler.UserHandler,
	) {
		engine.GET("/users", UserHandler.GetUsers)
		engine.GET("/users/:id", UserHandler.GetUser)
		engine.POST("/users", UserHandler.PostUser)
		engine.PUT("/users/:id", UserHandler.PutUser)
		engine.DELETE("/users/:id", UserHandler.PutUser)
	})

	srv := &http.Server{Addr: ":8080", Handler: engine}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("listen: %s\n", err)
		}
	}()

	// graceful shutdown
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
}
