GOOS ?= windows
GOARCH ?= amd64

EXE_SUFFIX ?=

ifeq ($(GOOS),windows)
EXE_SUFFIX = .exe
endif

compile: 
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o bin/output$(EXE_SUFFIX)

run: compile
	./bin/output.exe
