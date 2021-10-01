package api_request

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FoodType int

const (
	SeaFood FoodType = iota + 1
	Veg
	Mexican
	Italian
)

func (f FoodType) String() string {
	return [...]string{"SeaFood", "Veg", "Mexican", "Italian"}[f-1]
}

func (f FoodType) EnumIndex() int {
	return int(f)
}

var (
	EmptyOrderSummary = OrderSummary{}
)

type OrderSummary struct {
	OrderId           primitive.ObjectID `bson:"_id" json:"order_id,omitempty"`
	RestaurantName    string             `json:"restaurant_name",omitempty`
	RestaurantAddress string             `json:"restaurant_address,omitempty"`
	FoodName          string             `json:"food_name",omitempty`
	Address           string             `json:"user_address",omitempty`
}

type OrderDetails struct {
	RestaurantName    string `json:"restaurant_name"`
	RestaurantAddress string `json:"restaurant_address"`
	FoodName          string `json:"food_name"`
	//FoodType          FoodType `json: "food_type"`
}

type OrderCreateRequest struct {
	OrderId      primitive.ObjectID `bson:"_id" json:"order_id,omitempty"`
	UserName     string             `json:"name"`
	Address      string             `json:"user_address"`
	OrderDetails OrderDetails       `json:"order_details"`
}

type OrderUpdateRequest struct {
	Address      string             `json:"user_address"`
	OrderDetails OrderDetails       `json:"order_details"`
}
