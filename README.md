# FoodOrderingSystem
online food ordering system


**GET ALL**: /orders

**GET By Id**: /order/:id

**POST**: /order
Body: 
{
    "name": "Name",
    "user_address": "Address",
    "order_details": {
        "restaurant_name": "Restaurant",
        "restaurant_address": "Address",
        "food_name": "anyfood"
    }
}


**DELETE**: /order/:id

**PUT**: /order/:id
Body:
{
    "user_address": "India",
    "order_details": {
        "restaurant_name": "Restaurant",
        "restaurant_address": "Address",
        "food_name": "anyfood"
    }
}

**Dockerisation**

To build image: 
docker build -t foodorderingsystem -p 8080:8080 .

To compose: 
docker-compose up

**Deploy to Kubernetes single cluster(minikube)**

1) minikube start
2) kubectl apply -f foodorderingsystem-deployment.yaml,foodorderingsystem-service.yaml, mongodb-deployment.yaml, mongodb-persistentvolumeclaim.yaml, mongodb-service.yaml
3) kubectl get pods/service/nodes/deployments/all
3) minikube service foodorderingsystem --url
