package main

import (
	"fmt"
	"log"

	"github.com/skip2/go-qrcode"
)

func GenerateTOTPQrCode(base32Secret string) {
	account := "fabricio.eike.s@gmail.com"
	issuer := "GolangTOTPStudy"
	counter := 0

	data := fmt.Sprintf(
		"otpauth://hotp/%s:%s?secret=%s&issuer=%s&counter=%d",
		issuer, account, base32Secret, issuer, counter,
	)

	err := qrcode.WriteFile(data, qrcode.Medium, 256, "totp_qr.png")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("QR Code gerado.")
}
