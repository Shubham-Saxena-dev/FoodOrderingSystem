package routes

import (
	"FoodOrderingSystem/controllers"
	"github.com/gin-gonic/gin"
)

type Routes interface {
	RegisterHandlers()
}

type route struct {
	engine     *gin.Engine
	controller controllers.Controller
}

func NewRoutesHandler(engine *gin.Engine, controller controllers.Controller) Routes {
	return &route{engine: engine, controller: controller}
}

func (r *route) RegisterHandlers() {
	r.engine.GET("/orders", r.controller.GetAllOrders())
	r.engine.GET("/order/:id", r.controller.GetOrderById())
	r.engine.POST("/order", r.controller.CreateOrder())
	r.engine.PUT("/order/:id", r.controller.UpdateOrder())
	r.engine.DELETE("/order/:id", r.controller.CancelOrder())
}
