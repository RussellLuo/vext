package vext_test

import (
	"fmt"
	"time"

	v "github.com/RussellLuo/validating/v3"
	"github.com/RussellLuo/vext"
)

func Example_time() {
	value := "2006-01-02T15:04:05" // missing timezone
	if err := v.Validate(v.Value(value, vext.Time(time.RFC3339))); err != nil {
		fmt.Printf("err: %v\n", err)
	}

	// Output:
	// err: INVALID(invalid time)
}

func Example_ip() {
	value := "127.0.0."
	if err := v.Validate(v.Value(value, vext.IP())); err != nil {
		fmt.Printf("err: %v\n", err)
	}

	// Output:
	// err: INVALID(invalid IP)
}

func Example_email() {
	value := "foo#example.com"
	if err := v.Validate(v.Value(value, vext.Email())); err != nil {
		fmt.Printf("err: %v\n", err)
	}

	// Output:
	// err: INVALID(invalid email)
}
