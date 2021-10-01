package services

import (
	"FoodOrderingSystem/api_request"
	"FoodOrderingSystem/repositories"
)

type Service interface {
	GetAllOrders() ([]api_request.OrderSummary, error)
	GetOrderById(string) (api_request.OrderSummary, error)
	CreateOrder(api_request.OrderCreateRequest) (api_request.OrderSummary, error)
	UpdateOrder(api_request.OrderUpdateRequest, string) (api_request.OrderSummary, error)
	CancelOrder(string) (string, error)
}

type service struct {
	repo repositories.Repository
}

func NewService(repo repositories.Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetAllOrders() ([]api_request.OrderSummary, error) {
	return s.repo.GetAllOrders()
}

func (s *service) GetOrderById(id string) (api_request.OrderSummary, error) {
	return s.repo.GetOrderById(id)
}

func (s *service) CreateOrder(request api_request.OrderCreateRequest) (api_request.OrderSummary, error) {
	return s.repo.CreateOrder(request)
}

func (s service) UpdateOrder(request api_request.OrderUpdateRequest, id string) (api_request.OrderSummary, error) {
	return s.repo.UpdateOrder(request, id)
}

func (s *service) CancelOrder(id string) (string, error) {
	return s.repo.CancelOrder(id)
}
