package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"time"
)

func main() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"POST", "GET", "PUT", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	v1 := router.Group("/api")
	{
		v1.GET("/info", infoHandler)

		v1.POST("/register", Register)

		v1.POST("/login", Login)
	}

	authRoutes := v1.Group("/auth/", TokenMiddleware)
	{
		authRoutes.POST("/submitUserFoodChoice", submitUserFoodChoiceHandler)
		authRoutes.GET("/getUserFoodChoice", getUserFoodChoiceHandler)
		authRoutes.GET("/getCurrentWeekFoodList", getCurrentWeekFoodListHandler)
	}

	adminRoutes := v1.Group("/auth/admin/", TokenMiddleware)
	{
		adminRoutes.POST("/updateFoodChoiceList", updateFoodChoiceListHandler)
		adminRoutes.GET("/getTodaySummary", getTodaySummaryHandler)
	}

	router.Run(":3002")
}

func infoHandler(context *gin.Context) {
	context.JSON(200,
		gin.H{
			"status": "sama",
		},
	)
}
