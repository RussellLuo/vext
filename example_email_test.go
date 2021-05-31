package vext_test

import (
	"fmt"

	"github.com/RussellLuo/validating-ext"
	v "github.com/RussellLuo/validating/v2"
)

func ExampleEmail() {
	zeroOrEmail := vext.ZeroOr(vext.Email())

	value := "foo#example.com"
	if err := v.Validate(v.Value(&value, zeroOrEmail)); err != nil {
		fmt.Printf("err: %v\n", err)
	}

	// Output:
	// err: INVALID(invalid email)
}
