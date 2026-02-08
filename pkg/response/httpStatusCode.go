package response

const (
	// HTTP Status Codes
	ErrCodeSuccess      = 20001
	ErrCodeParamInvalid = 20003 //email is invalid
	ErrTokenInvalid     = 30001
	//Register Code
	ErrCodeUserHasExists = 50001 // user da ton tai
)

// message
var msg = map[int]string{
	ErrCodeSuccess:       "Success",
	ErrCodeParamInvalid:  "Email is invalid",
	ErrTokenInvalid:      "Token is invalid",
	ErrCodeUserHasExists: "User already exists",
}
