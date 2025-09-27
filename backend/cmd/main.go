package main

import (
	"ethglobal-nd-backend/internal/httphandlers"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	jwtware "github.com/gofiber/jwt/v3"
)

func main() {

	os.Setenv("JWKS_URL", "https://auth.privy.io/api/v1/apps/cmg17gsej004qla0dar9k36t8/jwks.json")
	app := fiber.New()

	app.Use(cors.New())

	keySetURLs := []string{os.Getenv("JWKS_URL")}
	var OptionalJWT = func(ctx *fiber.Ctx) bool {
		token := ctx.Get("Authorization")

		return len(token) == 0
	}

	apiGroup := app.Group("/api")
	v1Group := apiGroup.Group("/v1")
	v1Group.Use(jwtware.New(jwtware.Config{
		Filter:     OptionalJWT,
		KeySetURLs: keySetURLs,
		ContextKey: "token",
	}))

	// app.Post("/session", httphandlers.SessionHandler)

	v1Group.Post("/session", httphandlers.SessionHandler)
	v1Group.Post("/session/approve", httphandlers.SessionApprovalHandler)

	app.Listen(":5509")
}
