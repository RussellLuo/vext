package vext_test

import (
	"fmt"
	"time"

	v "github.com/RussellLuo/validating/v2"
	"github.com/RussellLuo/vext"
)

func ExampleTime() {
	zeroOrTime := vext.ZeroOr(vext.Time(time.RFC3339))

	value := "2006-01-02T15:04:05" // missing timezone
	if err := v.Validate(v.Value(&value, zeroOrTime)); err != nil {
		fmt.Printf("err: %v\n", err)
	}

	// Output:
	// err: INVALID(invalid time)
}
