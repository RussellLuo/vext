package vext_test

import (
	"fmt"

	"github.com/RussellLuo/validating-ext"
	v "github.com/RussellLuo/validating/v2"
)

func ExampleIP() {
	zeroOrIP := vext.ZeroOr(vext.IP())

	value := "127.0.0."
	if err := v.Validate(v.Value(&value, zeroOrIP)); err != nil {
		fmt.Printf("err: %v\n", err)
	}

	// Output:
	// err: INVALID(invalid IP)
}
