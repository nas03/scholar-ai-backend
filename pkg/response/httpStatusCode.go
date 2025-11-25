package response

const (
	// Success codes
	CodeSuccess = 200

	// User related codes
	CodeRegisterInternalError = 2001
	CodeUserAlreadyExists     = 2002
	CodeUserNotFound          = 2003
	CodeFailedGetUser         = 2004
	CodeFailedUpdateUser      = 2005
	CodeInvalidInput          = 2006
	CodeInvalidOTP            = 2007
	CodeOTPExpired            = 2008
	CodeEmailNotVerified      = 2009
	CodePhoneNotVerified      = 2010
	CodeInvalidEmail          = 2011
	CodeInvalidUsername       = 2012
	CodeEmptyPassword         = 2013
	CodeOTPNotFound           = 2014

	// Mail related codes
	CodeMailConfigMissing    = 3001
	CodeMailUsernameMissing  = 3002
	CodeMailPasswordMissing  = 3003
	CodeMailClientCreation   = 3004
	CodeMailConnectionFailed = 3005
	CodeMailSendFailed       = 3006
)

// Error messages mapping (following fidecwalletserver pattern)
var msg = map[int]string{
	CodeSuccess: "Success",

	// User related messages
	CodeRegisterInternalError: "Internal server error occurred during registration",
	CodeUserAlreadyExists:     "User already exists with this email or username",
	CodeUserNotFound:          "User not found",
	CodeFailedGetUser:         "Failed to retrieve user information",
	CodeFailedUpdateUser:      "Failed to update user information",
	CodeInvalidInput:          "Invalid input parameters provided",
	CodeOTPNotFound:           "Failed to retrieve otp",
	CodeInvalidOTP:            "Invalid OTP provided",
	CodeOTPExpired:            "OTP has expired",
	CodeEmailNotVerified:      "Email address not verified",
	CodePhoneNotVerified:      "Phone number not verified",
	CodeInvalidEmail:          "Invalid email format",
	CodeInvalidUsername:       "Invalid username format",
	CodeEmptyPassword:         "Password cannot be empty",

	// Mail related messages
	CodeMailConfigMissing:    "Mail configuration is missing",
	CodeMailUsernameMissing:  "Mail username is missing",
	CodeMailPasswordMissing:  "Mail password is missing",
	CodeMailClientCreation:   "Failed to create mail client",
	CodeMailConnectionFailed: "Failed to connect to mail service",
	CodeMailSendFailed:       "Failed to send email",
}
