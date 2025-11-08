namespace go user

struct GetUserRequest {
    1: i64 userID,
}

struct User {
    1: i64 userID,
    2: string name,
    3: string email,
}

struct GetUserResponse {
    1: User user,
    2: string message,
}

service UserService {
    GetUserResponse GetUserInfo(1: GetUserRequest req),
}