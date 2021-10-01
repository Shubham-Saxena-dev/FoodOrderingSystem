package repositories

import (
	"FoodOrderingSystem/api_request"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	GetAllOrders() ([]api_request.OrderSummary, error)
	GetOrderById(string) (api_request.OrderSummary, error)
	CreateOrder(api_request.OrderCreateRequest) (api_request.OrderSummary, error)
	UpdateOrder(api_request.OrderUpdateRequest, string) (api_request.OrderSummary, error)
	CancelOrder(string) (string, error)
}

type repository struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewMongoRepository(collection *mongo.Collection, ctx context.Context) Repository {
	return &repository{
		ctx:        ctx,
		collection: collection,
	}
}

func (r *repository) GetAllOrders() ([]api_request.OrderSummary, error) {
	cursor, err := r.collection.Find(r.ctx, bson.D{})
	defer cursor.Close(r.ctx)

	if err != nil {
		return []api_request.OrderSummary{}, err
	}

	var orders []api_request.OrderCreateRequest

	if cursor.All(r.ctx, &orders); err != nil {
		return []api_request.OrderSummary{}, err
	}

	var orderSummary []api_request.OrderSummary
	for _, order := range orders {

		orderSummary = append(orderSummary, api_request.OrderSummary{
			OrderId:           order.OrderId,
			RestaurantName:    order.OrderDetails.RestaurantName,
			RestaurantAddress: order.OrderDetails.RestaurantAddress,
			Address:           order.Address,
			FoodName:          order.OrderDetails.FoodName,
		})
	}
	return orderSummary, nil
}

func (r *repository) GetOrderById(id string) (api_request.OrderSummary, error) {
	uid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return api_request.EmptyOrderSummary, err
	}
	var order api_request.OrderCreateRequest
	err = r.collection.FindOne(r.ctx, bson.M{
		"_id": uid,
	}).Decode(&order)

	if err != nil {
		return api_request.EmptyOrderSummary, err
	}

	return api_request.OrderSummary{
		OrderId:           order.OrderId,
		RestaurantName:    order.OrderDetails.RestaurantName,
		RestaurantAddress: order.OrderDetails.RestaurantAddress,
		Address:           order.Address,
		FoodName:          order.OrderDetails.FoodName,
	}, nil
}

func (r *repository) CreateOrder(request api_request.OrderCreateRequest) (api_request.OrderSummary, error) {
	request.OrderId = primitive.NewObjectID()
	_, err := r.collection.InsertOne(r.ctx, request)
	if err != nil {
		return api_request.EmptyOrderSummary, err
	}
	return r.GetOrderById(request.OrderId.Hex())
}

func (r *repository) UpdateOrder(request api_request.OrderUpdateRequest, id string) (api_request.OrderSummary, error) {
	uid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return api_request.EmptyOrderSummary, err
	}
	if _, err = r.collection.UpdateOne(r.ctx, bson.M{
		"_id": uid,
	}, bson.M{
		"$set": request,
	}); err != nil {
		return api_request.EmptyOrderSummary, nil
	}

	return r.GetOrderById(id)
}

func (r *repository) CancelOrder(id string) (string, error) {
	uid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return "uid conversion error", err
	}
	_, err = r.collection.DeleteOne(r.ctx, bson.M{
		"_id": uid,
	})
	if err != nil {
		return "unable to delete", err
	}

	return "order deleted successfully", err
}
