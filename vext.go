package vext

import (
	"time"

	v "github.com/RussellLuo/validating/v2"
)

// ZeroOr is a composite validator factory used to create a validator, which will
// succeed if the field's value is zero, or if the given validator succeeds.
//
// ZeroOr will return the last error from the given validator if it fails.
func ZeroOr(validator v.Validator) v.Validator {
	zeroOrV := v.Any(
		v.Not(v.Nonzero()),
		validator,
	)

	return v.Func(func(field v.Field) v.Errors {
		errs := zeroOrV.Validate(field)
		if len(errs) == 0 {
			return nil
		}

		// Return the last error if fails.
		return errs[len(errs)-1:]
	})
}

// NewStringValidatorFactory creates a leaf validator factory, which will
// create a validator for validating a string value.
//
// The final validator will succeed if isValid returns true for a given string
// value. If it fails, the INVALID message is specified by msg.
func NewStringValidatorFactory(isValid func(string) bool, msg string) func() *v.MessageValidator {
	return func() (mv *v.MessageValidator) {
		mv = &v.MessageValidator{
			Message: msg,
			Validator: v.Func(func(field v.Field) v.Errors {
				switch t := field.ValuePtr.(type) {
				case *string:
					if !isValid(*t) {
						return v.NewErrors(field.Name, v.ErrInvalid, mv.Message)
					}
					return nil
				default:
					return v.NewErrors(field.Name, v.ErrUnsupported, "is unsupported")
				}
			}),
		}
		return
	}
}

// Time is a leaf validator factory used to create a validator, which will
// succeed when the field's value matches the given time format specified by layout.
func Time(layout string) (mv *v.MessageValidator) {
	isValid := func(value string) bool {
		_, err := time.Parse(layout, value)
		return err == nil
	}
	return NewStringValidatorFactory(isValid, "invalid time")()
}
