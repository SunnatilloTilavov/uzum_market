package config

const (
	ERR_INFORMATION     = "The server has received the request and is continuing the process"
	SUCCESS             = "The request was successful"
	ERR_REDIRECTION     = "You have been redirected and the completion of the request requires further action"
	ERR_BADREQUEST      = "Bad request"
	ERR_INTERNAL_SERVER = "While the request appears to be valid, the server could not complete the request"
	SmtpServer          = "smtp.gmail.com"
	SmtpPort            = "587"
	SmtpUsername        = "amirjonqodirov28@gmail.com"
	SmtpPassword        = "rabj vokw qbia qipw"
	
	CUSTOMER_TYPE = "customer"
	SELLER_TYPE = "seller"
	SYSTEM_TYPE = "system_user"
)

var SignedKey = []byte(`AtRdbumqoPjbcNjNhBgtmdAnRJyPQVXjwMPNYNbv`)
