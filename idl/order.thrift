namespace go order

struct CreateOrderRequest {
    1: i64 userID,
    2: string product,
}

struct OrderInfo {
    1: i64 orderID,
    2: i64 userID,
    3: string product,
    4: string userName,
}

struct CreateOrderResponse {
    1: OrderInfo order,
    2: string message,
}

service OrderService {
    CreateOrderResponse CreateOrder(1: CreateOrderRequest req),
}