namespace go userdemo



struct BaseResp {
    1: i64 status_code
    2: string status_message
    3: i64 service_time
}

struct User {
    1: i64 user_id
    2: string user_name
    3: string avatar
}

struct CreateUserRequest {
    1: string user_name
    2: string password
}

struct CreateUserResponse {
    1: BaseResp base_resp
}

struct MGetUserRequest {
    1: list<i64> user_ids
}

struct MGetUserResponse {
    1: list<User> users
    2: BaseResp base_resp
}

struct CheckUserRequest {
    1: string user_name
    2: string password
}

struct CheckUserResponse {
    1: i64 user_id
    2: BaseResp base_resp
}

service UserService {
    CreateUserResponse CreateUser(1: CreateUserRequest req)
    MGetUserResponse MGetUser(1: MGetUserRequest req)
    CheckUserResponse CheckUser(1: CheckUserRequest req)
}