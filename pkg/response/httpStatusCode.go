package response

const (
	// Success
	CodeSuccess = 20000

	// General Errors (40000 - 49999)
	CodeInvalidParams = 40001
	CodeServerBusy    = 50000

	// Auth Errors (10000 - 19999)
	CodeInvalidCredentials = 10001
	CodeTokenInvalid       = 10002
	CodeTokenExpired       = 10003

	// User Errors (20000 - 29999)
	CodeUserNotFound         = 20001
	CodeUserAlreadyExists    = 20002
	CodeUserAccountLocked    = 20003
	CodeUserEmailNotVerified = 20004
	CodeUserPhoneNotVerified = 20005
	CodeUserInvalidPassword  = 20006
	CodeUserInvalidEmail     = 20007
	CodeUserInvalidUsername  = 20008
	CodeUserCreationFailed   = 20009
	CodeUserUpdateFailed     = 20010
	CodeUserDeletionFailed   = 20011

	// OTP Errors (30000 - 39999)
	CodeOTPInvalid    = 30001
	CodeOTPExpired    = 30002
	CodeOTPNotFound   = 30003
	CodeOTPSendFailed = 30004

	// Mail Errors (50000 - 59999) - keeping original range for now but aligning pattern
	CodeMailSendFailed       = 50001
	CodeMailConfigMissing    = 50002
	CodeMailConnectionFailed = 50003
)

// msg maps error codes to user-friendly messages
var msg = map[int]string{
	// Success
	CodeSuccess: "Success",

	// General
	CodeInvalidParams: "Invalid parameters provided",
	CodeServerBusy:    "Server is busy, please try again later",

	// Auth
	CodeInvalidCredentials: "Invalid email or password",
	CodeTokenInvalid:       "Invalid authentication token",
	CodeTokenExpired:       "Authentication token has expired",

	// User
	CodeUserNotFound:         "User not found",
	CodeUserAlreadyExists:    "User already exists",
	CodeUserAccountLocked:    "User account is locked",
	CodeUserEmailNotVerified: "Email address is not verified",
	CodeUserPhoneNotVerified: "Phone number is not verified",
	CodeUserInvalidPassword:  "Invalid password format",
	CodeUserInvalidEmail:     "Invalid email format",
	CodeUserInvalidUsername:  "Invalid username format",
	CodeUserCreationFailed:   "Failed to create user account",
	CodeUserUpdateFailed:     "Failed to update user information",
	CodeUserDeletionFailed:   "Failed to delete user account",

	// OTP
	CodeOTPInvalid:    "Invalid OTP code",
	CodeOTPExpired:    "OTP code has expired",
	CodeOTPNotFound:   "OTP code not found",
	CodeOTPSendFailed: "Failed to send OTP",

	// Mail
	CodeMailSendFailed:       "Failed to send email",
	CodeMailConfigMissing:    "Mail configuration is missing",
	CodeMailConnectionFailed: "Failed to connect to mail service",
}

// GetMsg retrieves the message for a given error code
func GetMsg(code int) string {
	if m, ok := msg[code]; ok {
		return m
	}
	return "Unknown error"
}
