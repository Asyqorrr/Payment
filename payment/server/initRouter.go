package server

import (
	// "payment/controller"

	"payment/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter(handler *controller.ControllerManager) *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")

	api.GET("/home", func(ctx *gin.Context) {
		ctx.String(200, "Hello Gin FB")
	})

	entityRoute := api.Group("/entity")
	{
		entityRoute.GET("/",handler.GetListEntity)
		entityRoute.GET("",handler.GetListEntity)

		entityRoute.GET("/:id",handler.GetEntityById)

		entityRoute.POST("/",handler.CreateEntity)
		entityRoute.POST("",handler.CreateEntity)

		entityRoute.PUT("/:id",handler.UpdateEntity)

		entityRoute.DELETE("/:id",handler.DeleteEntity)
	}

	transactionRoute := api.Group("/transaction")
	{
		transactionRoute.POST("/", handler.CreatePaymentTransaction)
	}

	return router
}
