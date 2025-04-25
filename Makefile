OUT_DIR=bin

build:
	go build -o $(OUT_DIR)/hotp.exe src/main.go src/hotp.go

run:
	./bin/hotp