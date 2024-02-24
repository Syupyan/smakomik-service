package config

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnDb() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	return db
}

func AuthMiddlewareWithConfig(config jwtware.Config) fiber.Handler {
    return func(c *fiber.Ctx) error {
        // Extract token from cookie
        cookie := c.Cookies("token")
        if cookie == "" {
            return c.SendStatus(fiber.StatusUnauthorized)
        }

        // Verify token
        token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
            // Check the signing method
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
            }
            return config.SigningKey.Key, nil
        })

        if err != nil {
            return c.SendStatus(fiber.StatusUnauthorized)
        }

        // Store token claims in locals
        c.Locals("user", token)

        return c.Next()
    }
}


