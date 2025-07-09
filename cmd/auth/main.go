package main

import (
	"github.com/gin-gonic/gin"
	"github.com/muromeo1/go/pkg/auth"
	"github.com/muromeo1/go/pkg/config"

	"flag"
	"log"
	"net/http"
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

	r.GET("/api/auth/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	r.POST("/api/users", func(c *gin.Context) {
		register := auth.Register{}

		if err := c.ShouldBindJSON(&register); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		token, err := auth.UserCreator(register)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token})
	})

	r.POST("/api/sessions", func(c *gin.Context) {
		login := auth.Login{}

		if err := c.ShouldBindJSON(&login); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		token, err := auth.SessionCreator(login)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token})
	})

	r.GET("/api/sessions", func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			return
		}

		claims, err := auth.TokenDecode(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"exp": claims["exp"]})
	})

	r.Run(":" + config.Values.Port)
}
