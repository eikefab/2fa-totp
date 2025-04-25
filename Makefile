OUT_DIR=bin

build:
	go build -o $(OUT_DIR)/hmac.exe src/main.go

run:
	./bin/hmac