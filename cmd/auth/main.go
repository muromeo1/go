package main

import (
	"flag"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/muromeo1/go/pkg/auth"
	"github.com/muromeo1/go/pkg/config"
)

func main() {
	databaseUrl := flag.String("database-url", "", "Postgres database url")
	jwtSecret := flag.String("secret", "", "JWT Secret")
	ginMode := flag.String("mode", "debug", "Gin mode")
	port := flag.String("port", "8080", "Port to run the http server")

	flag.Parse()

	if *databaseUrl == "" {
		log.Fatalf("Database url is missing")
	}

	if *jwtSecret == "" {
		log.Fatalf("Secret is missing")
	}

	config.Values = config.Struct{
		DatabaseUrl: *databaseUrl,
		JWTSecret:   *jwtSecret,
		GinMode:     *ginMode,
		Port:        *port,
	}

	gin.SetMode(config.Values.GinMode)

	r := gin.Default()
	db := config.PG()
	db.AutoMigrate(&auth.User{})

	r.GET("/api/auth/health", auth.HealthCheckHandler)
	r.POST("/api/users", auth.UserCreateHandler)
	r.POST("/api/sessions", auth.SessionCreateHandler)
	r.GET("/api/sessions", auth.SessionAuthenticateHandler)

	r.Run(":" + config.Values.Port)
}
