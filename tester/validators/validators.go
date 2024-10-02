package validators

import (
	"unicode"

	"github.com/go-playground/validator/v10"
)

func UsernameValidate(fl validator.FieldLevel) bool {
	// handling ->  not first not last == underScore
	username := fl.Field().String()
	const underScore uint8 = 95

	if username[0] == underScore || username[len(username)-1] == underScore {
		return false
	}
	for index, item := range username {
		// handling -> _ == 95 -> not together
		if item == int32(underScore) && username[index+1] == underScore {
			return false
		}
		// only digits and letter and _
		if !unicode.IsDigit(item) && !unicode.IsLetter(item) && item != int32(underScore) {
			return false
		}

	}

	return false
}
