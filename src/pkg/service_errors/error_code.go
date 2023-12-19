package service_errors

const (
	// Token
	UnExpectedError = "Expected error"
	ClaimsNotFound  = "Claims not found"
	TokenRequired   = "token required"
	TokenExpired    = "token expired"
	TokenInvalid    = "token invalid"
	// OTP
	OtpExists   = "Otp exists"
	OtpUsed     = "Otp used"
	OtpNotValid = "Otp invalid"

	// User
	EmailExists       = "Email exists"
	UsernameExists    = "Username exists"
	UsernameNotExists = "Username not exists"
	PermissionDenied  = "Permission denied"

	// DB
	RecordNotFound = "record not found"
)
