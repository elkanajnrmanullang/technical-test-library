package middleware

import (
	"strings"

	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/infrastructure/config"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/infrastructure/utility"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"errors"
)

type JWTAuthMiddleware struct {
	jwtUtility *utility.JWTUtility
	jwtConfig  config.JwtConfig
}

func NewJWTAuthMiddleware(jwtUtil *utility.JWTUtility, jwtCfg config.JwtConfig) *JWTAuthMiddleware {
	return &JWTAuthMiddleware{
		jwtUtility: jwtUtil,
		jwtConfig:  jwtCfg,
	}
}

func (j *JWTAuthMiddleware) Protected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": true, "msg": "Missing or malformed JWT"})
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": true, "msg": "Missing or malformed JWT"})
		}

		tokenString := parts[1]
		claims, err := j.jwtUtility.ParseToken(tokenString)
		if err != nil {
			if errors.Is(err, jwt.ErrTokenExpired) {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": true, "msg": "JWT token is expired"})
			}
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": true, "msg": "Invalid JWT token"})
		}

		// Store user information from token in context if needed
		c.Locals("user_id", claims["user_id"])
		return c.Next()
	}
}
