namespace go userInfo

struct BaseResponse {
    1:i64 status_code
    2:string status_msg
}

struct UserInfo {
    1:list<string> categories
    2:list<string> tags
    3:list<string> ips
}

struct GetUserInfoRequest {
    1:string user_id
}

struct GetUserInfoResponse {
    1:BaseResponse base_resp
    2:UserInfo user_info
}

service UserService {
    GetUserInfoResponse GetUserInfo(1:GetUserInfoRequest req)
}