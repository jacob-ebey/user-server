package validate

import (
	"github.com/asaskevich/govalidator"
)

func Email(e string) (bool, string) {
	valid := govalidator.IsEmail(e)

	if valid {
		return true, ""
	}

	return false, "The provided email address is not valid."
}
