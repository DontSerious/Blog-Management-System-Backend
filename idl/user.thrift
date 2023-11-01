namespace go user

struct BaseResponse {
    1:i64 status_code
    2:string status_msg
}

struct UserInfo {
    1:list<string> categories
    2:list<string> tags
    3:list<string> ips
}

struct CreateUserRequest {
    1:string username
    2:string password
}

struct CreateUserResponse {
    1:BaseResponse base_resp
    2:i64 user_id
}

struct CheckUserRequest {
    1:string username
    2:string password
}

struct CheckUserResponse {
    1:BaseResponse base_resp
    2:i64 user_id
}

struct GetUserInfoRequest {
    1:i64 user_id
}

struct GetUserInfoResponse {
    1:BaseResponse base_resp
    2:UserInfo user_info
}

service UserService {
    CreateUserResponse CreateUser(1:CreateUserRequest req)
    CheckUserResponse CheckUser(1:CheckUserRequest req)
    GetUserInfoResponse GetUserInfo(1:GetUserInfoRequest req)
}