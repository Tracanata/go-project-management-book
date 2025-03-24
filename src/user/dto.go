package user

type ReqUserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RespSuccessLogin struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}
