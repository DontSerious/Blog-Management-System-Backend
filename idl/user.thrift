namespace go user

struct UserInfo {
    1:list<string> categories
    2:list<string> tags
}

struct User {
    1:string user_id
    2:string username
    3:UserInfo user_info
}

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

struct ChangePWDRequest {
    1:string user_id
    2:string password
}

struct ChangePWDResponse {
    1:BaseResponse base_resp
}

struct DelUserRequest {
    1:string user_id
}

struct DelUserResponse {
    1:BaseResponse base_resp
}

struct GetAllUserResponse {
    1:BaseResponse base_resp
    2:list<User> userList
}

service UserService {
    CreateUserResponse CreateUser(1:CreateUserRequest req)
    CheckUserResponse CheckUser(1:CheckUserRequest req)
    ChangePWDResponse ChangePWD(1:ChangePWDRequest req)
    DelUserResponse DelUser(1:DelUserRequest req)
    GetAllUserResponse GetAllUser()
}