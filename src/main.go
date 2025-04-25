package main

import (
	"encoding/base32"
	"fmt"
	"log"
)

func main() {
	secret := []byte("dummy-shared-secret") // GenerateSecret()
	base32Secret := base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(secret)

	// if err != nil {
	// 	fmt.Print("Ocorreu um erro ao gerar o segredo:", err)

	// 	return
	// }

	code, err := GenerateHOTP(secret, 123456, 6)

	if err != nil {
		log.Fatal("Ocorreu um erro ao gerar o código HOTP:", err)

		return
	}

	fmt.Println("Código HOTP:", code)

	totp, err := GenerateTOTP(secret, 6, 30)

	if err != nil {
		log.Fatal("Ocorreu um erro ao gerar o código TOTP:", err)
	}

	fmt.Println("Código TOTP:", totp)

	GenerateTOTPQrCode(base32Secret)
}
