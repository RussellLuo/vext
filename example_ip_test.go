package vext_test

import (
	"fmt"

	"github.com/RussellLuo/validating-ext"
	v "github.com/RussellLuo/validating/v2"
)

func Example_IP() {
	zeroOrIP := vext.ZeroOr(vext.IP())

	ip := "127.0.0."
	if err := v.Validate(v.Value(&ip, zeroOrIP)); err != nil {
		fmt.Printf("err: %v\n", err)
	}

	// Output:
	// err: INVALID(invalid IP)
}
