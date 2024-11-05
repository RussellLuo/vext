package vext

import (
	"fmt"
	"sync"
	"time"

	v "github.com/RussellLuo/validating/v3"
	"github.com/asaskevich/govalidator"
	"github.com/th0th/disposableemail"
	"golang.org/x/exp/slices"
)

var disposableEmail disposableemail.Service

var once sync.Once

// ASCII is a leaf validator factory used to create a validator, which will
// succeed when the field's value contains only ASCII chars.
func ASCII() *v.MessageValidator {
	return v.Is(govalidator.IsASCII).Msg("invalid ASCII")
}

// Alpha is a leaf validator factory used to create a validator, which will
// succeed when the field's value contains only ASCII letters (a-zA-Z).
func Alpha() *v.MessageValidator {
	return v.Is(govalidator.IsAlpha).Msg("invalid alpha")
}

// Alphanumeric is a leaf validator factory used to create a validator, which will
// succeed when the field's value contains only letters and numbers.
func Alphanumeric() *v.MessageValidator {
	return v.Is(govalidator.IsAlphanumeric).Msg("invalid alphanumeric")
}

// Base64 is a leaf validator factory used to create a validator, which will
// succeed when the field's value is a base64 encoded string.
func Base64() *v.MessageValidator {
	return v.Is(govalidator.IsBase64).Msg("invalid base64")
}

// CIDR is a leaf validator factory used to create a validator, which will
// succeed when the field's value is an valid CIDR notation (IPV4 & IPV6).
func CIDR() *v.MessageValidator {
	return v.Is(govalidator.IsCIDR).Msg("invalid CIDR")
}

// CreditCard is a leaf validator factory used to create a validator, which will
// succeed when the field's value is a credit card.
func CreditCard() *v.MessageValidator {
	return v.Is(govalidator.IsCreditCard).Msg("invalid credit card")
}

// DNSName is a leaf validator factory used to create a validator, which will
// succeed when the field's value is a DNS name.
func DNSName() *v.MessageValidator {
	return v.Is(govalidator.IsDNSName).Msg("invalid DNS name")
}

// DataURI is a leaf validator factory used to create a validator, which will
// succeed when the field's value is a base64 encoded data URI such as an image.
func DataURI() *v.MessageValidator {
	return v.Is(govalidator.IsDataURI).Msg("invalid data URI")
}

// DialString is a leaf validator factory used to create a validator, which will
// succeed when the field's value is a dial string.
func DialString() *v.MessageValidator {
	return v.Is(govalidator.IsDialString).Msg("invalid dial string")
}

// Email is a leaf validator factory used to create a validator, which will
// succeed when the field's value is an email.
func Email() *v.MessageValidator {
	return v.Is(govalidator.IsEmail).Msg("invalid email")
}

// EmailNonDisposable is a leaf validator factory used to create a validator, which will
// succeed when the field's value is a non-disposable email address.
// It uses the `is-email-disposable` package to check if the email domain is disposable.
func EmailNonDisposable() *v.MessageValidator {
	once.Do(func() {
		disposableEmail2, err := disposableemail.New()
		if err != nil {
			panic(err)
		}

		disposableEmail = disposableEmail2
	})

	messageValidator := v.MessageValidator{
		Message: "is disposable e-mail address",
	}

	messageValidator.Validator = v.Func(func(field *v.Field) v.Errors {
		value, ok := field.Value.(string)
		if !ok {
			return v.NewUnsupportedErrors("EmailNonDisposable", field, "")
		}

		checkResult := disposableEmail.Check(value)
		if checkResult.IsDisposable {
			return v.NewInvalidErrors(field, messageValidator.Message)
		}

		return nil
	})

	return &messageValidator
}

// Hash is a leaf validator factory used to create a validator, which will
// succeed when the field's value is a hash of type algorithm.
//
// Supported algorithms:
// - md4
// - md5
// - sha1
// - sha256
// - sha384
// - sha512
// - ripemd128
// - ripemd160
// - tiger128
// - tiger160
// - tiger192
// - crc32
// - crc32b
func Hash(algorithm string) *v.MessageValidator {
	algorithms := []string{
		"md4", "md5",
		"sha1", "sha256", "sha384", "sha512",
		"ripemd128", "ripemd160",
		"tiger128", "tiger160", "tiger192",
		"crc32", "crc32b",
	}
	if !slices.Contains(algorithms, algorithm) {
		panic(fmt.Errorf("unsupported hash algorithm: %s", algorithm))
	}

	isValid := func(value string) bool {
		return govalidator.IsHash(value, algorithm)
	}
	return v.Is(isValid).Msg(fmt.Sprintf("invalid %s hash", algorithm))
}

// HexNumber is a leaf validator factory used to create a validator, which will
// succeed when the field's value is a hexadecimal number.
func HexNumber() *v.MessageValidator {
	return v.Is(govalidator.IsHexadecimal).Msg("invalid hexadecimal number")
}

// HexColor is a leaf validator factory used to create a validator, which will
// succeed when the field's value is a hexadecimal color.
func HexColor() *v.MessageValidator {
	return v.Is(govalidator.IsHexcolor).Msg("invalid hexadecimal color")
}

// IP is a leaf validator factory used to create a validator, which will
// succeed when the field's value is an IPv4 or IPv6 string.
func IP() *v.MessageValidator {
	return v.Is(govalidator.IsIP).Msg("invalid IP")
}

// IPv4 is a leaf validator factory used to create a validator, which will
// succeed when the field's value is an IPv4 string.
func IPv4() *v.MessageValidator {
	return v.Is(govalidator.IsIPv4).Msg("invalid IPv4")
}

// IPv6 is a leaf validator factory used to create a validator, which will
// succeed when the field's value is an IPv6 string.
func IPv6() *v.MessageValidator {
	return v.Is(govalidator.IsIPv6).Msg("invalid IPv6")
}

// ISBN is a leaf validator factory used to create a validator, which will
// succeed when the field's value is an ISBN (version 10 or 13) string.
func ISBN(version int) *v.MessageValidator {
	versions := []int{10, 13}
	if !slices.Contains(versions, version) {
		panic(fmt.Errorf("unsupported ISBN version: %d", version))
	}

	isValid := func(value string) bool {
		return govalidator.IsISBN(value, version)
	}
	return v.Is(isValid).Msg(fmt.Sprintf("invalid ISBN %d", version))
}

// ISO3166Alpha2 is a leaf validator factory used to create a validator, which will
// succeed when the field's value is a two-letter country code.
func ISO3166Alpha2() *v.MessageValidator {
	return v.Is(govalidator.IsISO3166Alpha2).Msg("invalid ISO3166 Alpha-2")
}

// ISO3166Alpha3 is a leaf validator factory used to create a validator, which will
// succeed when the field's value is a three-letter country code.
func ISO3166Alpha3() *v.MessageValidator {
	return v.Is(govalidator.IsISO3166Alpha3).Msg("invalid ISO3166 Alpha-3")
}

// ISO4217 is a leaf validator factory used to create a validator, which will
// succeed when the field's value is an ISO currency code.
func ISO4217() *v.MessageValidator {
	return v.Is(govalidator.IsISO4217).Msg("invalid ISO4217")
}

// ISO693Alpha2 is a leaf validator factory used to create a validator, which will
// succeed when the field's value is a two-letter language code.
func ISO693Alpha2() *v.MessageValidator {
	return v.Is(govalidator.IsISO693Alpha2).Msg("invalid ISO693 Alpha-2")
}

// ISO693Alpha3 is a leaf validator factory used to create a validator, which will
// succeed when the field's value is a three-letter language code.
func ISO693Alpha3b() *v.MessageValidator {
	return v.Is(govalidator.IsISO693Alpha3b).Msg("invalid ISO693 Alpha-3b")
}

// Latitude is a leaf validator factory used to create a validator, which will
// succeed when the field's value is a latitude.
func Latitude() *v.MessageValidator {
	return v.Is(govalidator.IsLatitude).Msg("invalid latitude")
}

// Longitude is a leaf validator factory used to create a validator, which will
// succeed when the field's value is a longitude.
func Longitude() *v.MessageValidator {
	return v.Is(govalidator.IsLongitude).Msg("invalid longitude")
}

// MAC is a leaf validator factory used to create a validator, which will
// succeed when the field's value is a MAC address.
func MAC() *v.MessageValidator {
	return v.Is(govalidator.IsMAC).Msg("invalid MAC")
}

// MagnetURI is a leaf validator factory used to create a validator, which will
// succeed when the field's value is a magnet URI.
func MagnetURI() *v.MessageValidator {
	return v.Is(govalidator.IsMagnetURI).Msg("invalid magnet URI")
}

// MongoID is a leaf validator factory used to create a validator, which will
// succeed when the field's value is a hex-encoded representation of a MongoDB ObjectId.
func MongoID() *v.MessageValidator {
	return v.Is(govalidator.IsMongoID).Msg("invalid MongoDB ID")
}

// RGBColor is a leaf validator factory used to create a validator, which will
// succeed when the field's value is a RGB color in form rgb(RRR, GGG, BBB).
func RGBColor() *v.MessageValidator {
	return v.Is(govalidator.IsRGBcolor).Msg("invalid RGB color")
}

// SSN is a leaf validator factory used to create a validator, which will
// succeed when the field's value is a U.S. Social Security Number.
func SSN() *v.MessageValidator {
	return v.Is(govalidator.IsSSN).Msg("invalid SSN")
}

// Time is a leaf validator factory used to create a validator, which will
// succeed when the field's value matches the given time format specified by layout.
func Time(layout string) (mv *v.MessageValidator) {
	isValid := func(value string) bool {
		_, err := time.Parse(layout, value)
		return err == nil
	}
	return v.Is(isValid).Msg("invalid time")
}

// URL is a leaf validator factory used to create a validator, which will
// succeed when the field's value is an URL.
func URL() *v.MessageValidator {
	return v.Is(govalidator.IsURL).Msg("invalid URL")
}

// UUID is a leaf validator factory used to create a validator, which will
// succeed when the field's value is an UUID (version 3, 4 or 5) string.
func UUID() *v.MessageValidator {
	return v.Is(govalidator.IsUUID).Msg("invalid UUID")
}

// UUID is a leaf validator factory used to create a validator, which will
// succeed when the field's value is an UUID version 3 string.
func UUIDv3() *v.MessageValidator {
	return v.Is(govalidator.IsUUIDv3).Msg("invalid UUIDv3")
}

// UUID is a leaf validator factory used to create a validator, which will
// succeed when the field's value is an UUID version 4 string.
func UUIDv4() *v.MessageValidator {
	return v.Is(govalidator.IsUUIDv4).Msg("invalid UUIDv4")
}

// UUID is a leaf validator factory used to create a validator, which will
// succeed when the field's value is an UUID version 5 string.
func UUIDv5() *v.MessageValidator {
	return v.Is(govalidator.IsUUIDv5).Msg("invalid UUIDv5")
}
