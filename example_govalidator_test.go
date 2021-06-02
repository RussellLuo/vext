package vext_test

import (
	"fmt"

	v "github.com/RussellLuo/validating/v2"
	"github.com/RussellLuo/vext"
)

func Example_ip() {
	zeroOrIP := vext.ZeroOr(vext.IP())

	value := "127.0.0."
	if err := v.Validate(v.Value(&value, zeroOrIP)); err != nil {
		fmt.Printf("err: %v\n", err)
	}

	// Output:
	// err: INVALID(invalid IP)
}

func Example_email() {
	zeroOrEmail := vext.ZeroOr(vext.Email())

	value := "foo#example.com"
	if err := v.Validate(v.Value(&value, zeroOrEmail)); err != nil {
		fmt.Printf("err: %v\n", err)
	}

	// Output:
	// err: INVALID(invalid email)
}
