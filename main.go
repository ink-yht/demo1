package main

import (
	"demo-1/internal/service"
	"demo-1/internal/web"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

func main() {
	server := gin.Default()
	server.Use(cors.New(cors.Config{
		//AllowOrigins:     []string{"https://foo.com"},
		//AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Content-Type", "multipart/form-data"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "http://localhost") {
				return true
			}
			return strings.Contains(origin, "your.com")
		},
		MaxAge: 12 * time.Hour,
	}))
	svc := service.NewService()
	hd := web.NewHandler(svc)
	hd.RegisterRoutes(server)
	server.Run(":8088")
}
