# Order Management

## How To Run
`go run main.go`

## Table Design
![ERD](https://github.com/ktiarad/order-management/blob/main/assets/erd.png)

## API List / Features (with View)
### Sign In
```
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"email":"test@mail.com", "password":"test", "name":"test"}' \
  http://localhost:8080/signin
```

### Log In
```
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"email":"test@mail.com", "password":"test"}' \
  http://localhost:8080/login
```

### Get Customer Detail
```
curl --header "Content-Type: application/json" \
  --header "Authorization: Bearer <token>" \
  --request GET \
  http://localhost:8080/customers/detail/1
```

### Get Customer By Name
```
curl --header "Content-Type: application/json" \
  --header "Authorization: Bearer <token>" \
  --request GET \
  http://localhost:8080/customers/search/test
```

### Get All Customers
```
curl --header "Content-Type: application/json" \
  --header "Authorization: Bearer <token>" \
  --request GET \
  http://localhost:8080/customers/list/10/0
```

### Create Customer
```
curl --header "Content-Type: application/json" \
  --header "Authorization: Bearer <token>" \
  --request POST \
  --data '{"name":"test3"}' \
  http://localhost:8080/customers
```

### Update Customer
```
curl --header "Content-Type: application/json" \
  --header "Authorization: Bearer <token>" \
  --request PUT \
  --data '{"id":1, "name":"test2"}' \
  http://localhost:8080/customers
```

### Delete Customer
```
curl --header "Content-Type: application/json" \
  --header "Authorization: Bearer <token>" \
  --request DELETE \
  http://localhost:8080/customers/2
```

### Get Order Detail
```
curl --header "Content-Type: application/json" \
  --header "Authorization: Bearer <token>" \
  --request GET \
  http://localhost:8080/orders/detail/1
```

### Get Order By Title
```
curl --header "Content-Type: application/json" \
  --header "Authorization: Bearer <token>" \
  --request GET \
  http://localhost:8080/orders/search/test
```

### Create Order
```
curl --header "Content-Type: application/json" \
  --header "Authorization: Bearer <token>" \
  --request POST \
  --data '{"title":"test", "description":"test order", "customer_id":1, "status":"confirmed", "price":250000}' \
  http://localhost:8080/orders
```

### Update Order Status
```
curl --header "Content-Type: application/json" \
  --header "Authorization: Bearer <token>" \
  --request PUT \
  --data '{"id":1, "status":"proceed"}' \
  http://localhost:8080/orders/status
```

### Delete Order
```
curl --header "Content-Type: application/json" \
  --header "Authorization: Bearer <token>" \
  --request DELETE \
  http://localhost:8080/orders/2
```