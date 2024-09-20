package metrics

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/adaptor"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// SetupMetrics returns a new Fiber app that serves prometheus metrics on
// port 2112. The app is started in a goroutine to not block the main thread.
// The metrics path is /metrics.
func SetupMetrics() *fiber.App {
	app := fiber.New()
	app.Get("/metrics", adaptor.HTTPHandler(promhttp.Handler()))
	return app
}
