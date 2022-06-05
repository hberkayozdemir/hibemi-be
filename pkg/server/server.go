package server

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gitlab.com/modanisatech/marketplace/shared/errors"
	httpkitmw "gitlab.com/modanisatech/marketplace/shared/httpkit/middlewares"
	"gitlab.com/modanisatech/marketplace/shared/log"
	"go.uber.org/zap"
)

type Handler interface {
	RegisterRoutes(app *fiber.App)
}

type Server struct {
	app    *fiber.App
	port   string
	logger *zap.Logger
}

func New(port string, handlers []Handler, logger *zap.Logger) Server {
	app := fiber.New(fiber.Config{ErrorHandler: errors.Handler(logger)})
	server := Server{app: app, port: port, logger: logger}
	server.app.Use(recover.New())
	app.Use(httpkitmw.HeaderPropagator)
	app.Use(httpkitmw.PreAuthHeader)
	server.app.Use(cors.New())
	server.app.Use(log.Middleware(logger))
	server.addRoutes()

	for _, handler := range handlers {
		handler.RegisterRoutes(server.app)
	}
	return server
}

func (s Server) addRoutes() {
	s.app.Get("/health", healthCheck)
	s.app.Get("/metrics", adaptor.HTTPHandler(promhttp.Handler()))
}

func (s Server) Run() {
	shutdownChan := make(chan os.Signal, 1)
	signal.Notify(shutdownChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		shutdownSignal := <-shutdownChan
		s.logger.Info("Received interrupt signal", zap.String("shutdownSignal", shutdownSignal.String()))
		if err := s.app.Shutdown(); err != nil {
			s.logger.Info("Failed to shutdown gracefully", zap.Error(err))
			return
		}
		s.logger.Info("application shutdown gracefully")
	}()
	err := s.app.Listen(s.port)
	if err != nil {
		s.logger.Panic(err.Error())
	}
}

func (s Server) Stop() {
	if err := s.app.Shutdown(); err != nil {
		s.logger.Error(err.Error())
	}
}

func healthCheck(c *fiber.Ctx) error {
	c.Status(fiber.StatusOK)
	return nil
}
