package vext

import (
	"time"

	v "github.com/RussellLuo/validating/v3"
	"github.com/asaskevich/govalidator"
)

// Time is a leaf validator factory used to create a validator, which will
// succeed when the field's value matches the given time format specified by layout.
func Time(layout string) (mv *v.MessageValidator) {
	isValid := func(value string) bool {
		_, err := time.Parse(layout, value)
		return err == nil
	}
	return v.Is(isValid).Msg("invalid time")
}

func IP() *v.MessageValidator {
	return v.Is(govalidator.IsIP).Msg("invalid IP")
}

func Email() *v.MessageValidator {
	return v.Is(govalidator.IsEmail).Msg("invalid email")
}
