package vext

import (
	"github.com/asaskevich/govalidator"
)

var (
	IP    = NewStringValidatorFactory(govalidator.IsIP, "invalid IP")
	Email = NewStringValidatorFactory(govalidator.IsEmail, "invalid email")
)
