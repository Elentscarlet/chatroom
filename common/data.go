package common

const (
	TypeLogin      = "Login"
	TypeLoginReply = "LoginReply"
	TypeRegister   = "Register"
	TypeCommonMsag = "CommonMsg"
	TypeResopnse   = "Response"

	ServerError = 500

	// status code for login
	LoginError   = 403
	NotExit      = 404
	LoginSucceed = 200

	// status code for register
	HasExited        = 403
	RegisterSucceed  = 200
	PasswordNotMatch = 402
)

type Data struct {
	User    string `json:"user"`
	Type    string `json:"type"`
	Content string `json:"content"`
}

type RegisterMsg struct {
	UserName string
	Password string
}

type LoginMsg struct {
	UserName string
	Password string
}

type CommonMsg struct {
	UserName string
	Content  string
}

type ResponseMsg struct {
	Type  string
	Code  int
	Error string
	Data  string
}
