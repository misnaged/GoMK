APP:=mk2.exe
COMMON_PATH	?= $(shell pwd)
APP_ENTRY_POINT:=./main.go
BUILD_OUT_DIR:=./
GOOS	:= windows
GOARCH	:= amd64


build:
	env CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc GOOS=$(GOOS) GOARCH=$(GOARCH) go build -buildmode=c-shared -o main.exe main.go
