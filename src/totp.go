package main

import (
	"time"
)

// time-based hotp
func GenerateTOTP(secret []byte, digits int, period int64) (string, error) {
	counter := time.Now().Unix() / period

	return GenerateHOTP(secret, uint64(counter), digits)
}
