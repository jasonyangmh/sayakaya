package router

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jasonyangmh/sayakaya/handler"
	"github.com/jasonyangmh/sayakaya/middleware"
	"github.com/jasonyangmh/sayakaya/repository"
	"github.com/jasonyangmh/sayakaya/usecase"
	"gorm.io/gorm"
)

type Router struct {
	router *gin.Engine
}

func New(db *gorm.DB) *Router {
	r := gin.Default()
	r.Use(middleware.ErrorMiddleware())
	r.ContextWithFallback = true

	ur := repository.NewUserRepository(db)
	uu := usecase.NewUserUsecase(ur)
	uh := handler.NewUserHandler(uu)

	users := r.Group("/users")
	{
		users.POST("", uh.PostUser)
	}

	return &Router{
		router: r,
	}
}

func (r *Router) Start() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r.router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()

	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}
