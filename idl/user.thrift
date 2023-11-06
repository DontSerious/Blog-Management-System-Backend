namespace go user

struct BaseResponse {
    1:i64 status_code
    2:string status_msg
}

struct CreateUserRequest {
    1:string username
    2:string password
}

struct CreateUserResponse {
    1:BaseResponse base_resp
    2:string user_id
}

struct CheckUserRequest {
    1:string username
    2:string password
}

struct CheckUserResponse {
    1:BaseResponse base_resp
    2:string user_id
}

service UserService {
    CreateUserResponse CreateUser(1:CreateUserRequest req)
    CheckUserResponse CheckUser(1:CheckUserRequest req)
}