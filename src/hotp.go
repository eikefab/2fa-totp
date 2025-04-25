package main

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha1"
	"encoding/binary"
	"errors"
	"fmt"
	"math"
)

func GenerateSecret() ([]byte, error) {
	secret := make([]byte, 4)
	_, err := rand.Read(secret)

	if err != nil {
		return nil, err
	}

	return secret, nil
}

func GenerateHOTP(secret []byte, counter uint64, digits int) (string, error) {
	if digits < 6 {
		return "", errors.New("Quantidade de dígitos deve ser de no mínimo 6 (RFC 4226).")
	}

	counterBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(counterBytes, counter)

	hmacHash := hmac.New(sha1.New, secret)
	hmacHash.Write(counterBytes)
	hmacResult := hmacHash.Sum(nil)

	offset := hmacResult[len(hmacResult)-1] & 0x0F
	code := (int(hmacResult[offset])&0x7F)<<24 |
		(int(hmacResult[offset+1])&0xFF)<<16 |
		(int(hmacResult[offset+2])&0xFF)<<8 |
		(int(hmacResult[offset+3]) & 0xFF)

	modulo := int(math.Pow10(digits))
	hotp := code % modulo

	return fmt.Sprintf("%0*d", digits, hotp), nil
}
