package controllers

import (
	"FoodOrderingSystem/api_request"
	"FoodOrderingSystem/services"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Controller interface {
	GetAllOrders() gin.HandlerFunc
	GetOrderById() gin.HandlerFunc
	CreateOrder() gin.HandlerFunc
	UpdateOrder() gin.HandlerFunc
	CancelOrder() gin.HandlerFunc
}

type controller struct {
	service services.Service
}

func NewController(service services.Service) Controller {
	return &controller{
		service: service,
	}
}

func (c *controller) GetAllOrders() gin.HandlerFunc {
	return func(context *gin.Context) {
		response, err := c.service.GetAllOrders()
		if err != nil {
			handleError(context, err, http.StatusInternalServerError)
		}
		context.JSON(http.StatusOK, response)
	}
}

func (c *controller) GetOrderById() gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Param("id")
		response, err := c.service.GetOrderById(id)
		if err != nil {
			handleError(context, err, http.StatusInternalServerError)
		}
		context.JSON(http.StatusOK, response)
	}
}

func (c *controller) CreateOrder() gin.HandlerFunc {
	return func(context *gin.Context) {
		var order api_request.OrderCreateRequest
		if err := context.ShouldBindJSON(&order); err != nil {
			handleError(context, err, http.StatusBadRequest)
			return
		}
		response, err := c.service.CreateOrder(order)
		if err != nil {
			handleError(context, err, http.StatusInternalServerError)
		}
		context.JSON(http.StatusOK, response)
	}
}

func (c *controller) UpdateOrder() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var order api_request.OrderUpdateRequest
		id := ctx.Param("id")
		if err := ctx.ShouldBindJSON(&order); err != nil {
			handleError(ctx, err, http.StatusBadRequest)
			return
		}
		response, err := c.service.UpdateOrder(order, id)
		if err != nil {
			handleError(ctx, err, http.StatusInternalServerError)
		}
		ctx.JSON(http.StatusOK, response)
	}
}

func (c *controller) CancelOrder() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		response, err := c.service.CancelOrder(id)
		if err != nil {
			handleError(ctx, err, http.StatusInternalServerError)
		}
		ctx.JSON(http.StatusOK, response)
	}
}

func handleError(context *gin.Context, err error, statusCode int) {
	log.Error(err)
	context.JSON(statusCode, gin.H{
		"error": err.Error(),
	})
}
