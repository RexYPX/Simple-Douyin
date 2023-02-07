namespace go user

struct User {
    1: i64 id
    2: string name
    3: i64 follow_count
    4: i64 follower_count
    5: bool is_follow
}

// 用户注册
struct UserRegisterRequest {
    1: string username
    2: string password
}

struct UserRegisterResponse {
    1: i32 status_code
    2: string status_msg
    3: i64 user_id
    4: string token
}

// 用户登录
struct UserLoginRequest {
    1: string username
    2: string password
}

struct UserLoginResponse {
    1: i32 status_code
    2: string status_msg
    3: i64 user_id
    4: string token
}

// 用户信息
struct UserInfoRequest {
    1: i64 user_id
    2: string token
}

struct UserInfoResponse {
    1: i32 status_code
    2: string status_msg
    3:i64 id
	4:string name
	5:i64 follow_count
	6:i64 follower_count
	7:bool is_follow
}

service UserService {
    UserRegisterResponse UserRegister(1: UserRegisterRequest req)
    UserLoginResponse UserLogin(1: UserLoginRequest req)
    UserInfoResponse UserInfo(1: UserInfoRequest req)
}