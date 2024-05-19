namespace go note.user

include "base.thrift"

struct User {
    1: i64 user_id
    2: string user_name
    3: string avatar
}

struct CreateUserReq {
    1: string user_name
    2: string password
}

struct CreateUserResp {
    1: base.BaseResp base_resp
}

struct MGetUserReq {
    1: list<i64> user_ids
}

struct MGetUserResp {
    1: list<User> users
    2: base.BaseResp base_resp
}

struct CheckUserReq {
    1: string user_name
    2: string password
}

struct CheckUserResp {
    1: i64 user_id
    2: base.BaseResp base_resp
}

service UserService {
    CreateUserResp CreateUser(1: CreateUserReq req)
    MGetUserResp MGetUser(1: MGetUserReq req)
    CheckUserResp CheckUser(1: CheckUserReq req)
}