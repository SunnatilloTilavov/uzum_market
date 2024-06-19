package validator

import "regexp"


func PhoneValidator(phone string) bool {
	return regexp.MustCompile(`^\+998[0-9]{9}$`).MatchString(phone)
}

func MailValidator(mail string) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@gmail.com$`).MatchString(mail)
}