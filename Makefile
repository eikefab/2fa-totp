OUT_DIR=bin

build:
	go build -o $(OUT_DIR)/hotp.exe src/main.go src/hotp.go src/totp.go src/qrcode.go

run:
	./bin/hotp