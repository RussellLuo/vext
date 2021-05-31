package vext_test

import (
	"fmt"

	v "github.com/RussellLuo/validating/v2"
	"github.com/RussellLuo/vext"
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
