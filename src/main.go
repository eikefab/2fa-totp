package main

import "fmt"

func main() {
	secret, err := GenerateSecret()

	if err != nil {
		fmt.Print("Ocorreu um erro ao gerar o segredo:", err)

		return
	}

	code, err := GenerateHOTP(secret, 123456, 6)

	if err != nil {
		fmt.Println("Ocorreu um erro ao gerar o código:", err)

		return
	}

	fmt.Println("Código HOTP:", code)
}
