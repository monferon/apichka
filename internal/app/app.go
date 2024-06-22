package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/monferon/apichka/config"
	v1 "github.com/monferon/apichka/internal/controller/http/v1"
	"github.com/monferon/apichka/pkg/httpserver"
	"github.com/monferon/apichka/pkg/logger"
	"os"
	"os/signal"
	"syscall"
)

func Run(cfg *config.Config) {
	l := logger.NewLogger(cfg.Log.Level)
	//pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	//if err != nil {
	//	l.Fatal(fmt.Errorf("(app - Run - postgres. failed to connect to postgres: %w", err))
	//}
	//defer pg.Close()

	// HTTP Server
	handler := gin.New()
	v1.NewRouter(handler, l)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err := <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}
}
