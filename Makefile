.PHONY: all run test vendor docker build clean swag fmt

SHELL = /bin/bash

# preset constants and variables
PWD = $(shell pwd)
DEPS := $(wildcard *.go)

# color printing
CE = "\033[0m"
CRED = "\033[31m"
CGREEN = "\033[32m"
CYELLOW = "\033[33m"

# program detection
SWAG_EXISTS := $(shell which swag 2>/dev/null | grep -c "/swag" )
GOCILINT_EXISTS := $(shell which golangci-lint 2>/dev/null | grep -c "/golangci-lint" )
DOCKER_CHECK := $(shell sysctl net.inet.tcp.sack 2>/dev/null | grep -c '0')
GOFMT_CMD := goimports
GOFMT_EXIST := $(shell command -v $(GOFMT_CMD))
DIFF_FILES_CMD := git diff --name-status | grep -E "^[^D]" | grep -E "\.go$$" | grep -v "vendor/*" | grep -v "docs/*" | tr -d "[A-Z\t]"
CHANGE_FILES := $(shell $(DIFF_FILES_CMD))
GOMOD := $(shell go list -m)

run:
	@echo -e ${CGREEN} "make run.........................................................."
	swag init
	swag fmt
	go run main.go

swag:
	@echo -e ${CGREEN} "make swag........................................................."
	@if [ ${SWAG_EXISTS} -eq 0 ]; then  \
		export PATH=$$PATH:$$GOPATH/bin; \
		echo -e ${CYELLOW} "1. 若安装过但提示无法找到swag，尝试 source ~/.bashrc 或 ~/.zshrc"; \
		echo -e ${CYELLOW} "2. 若确实未安装过，请执行 make install-swag"; \
	else \
		echo "swag formatting........................................................."; \
		swag fmt; \
		echo -e ${CGREEN} "swag generating........................................................."; \
		if [ $$(swag i -g main.go 2>&1 | tee /dev/stderr | grep -c "cannot find type definition") -gt 0 ]; then \
			swag i --pd --parseInternal --parseDepth 1 -g main.go;\
		fi; \
		echo -e ${CGREEN} "swag finished........................................................."; \
    fi \

swag-fmt:
	@echo -e ${CGREEN} "make swag fmt........................................................."
	@swag fmt

install-swag:
	@if [ ${SWAG_EXISTS} -eq 0 ]; then  \
		GO111MODULE="on" go get github.com/swaggo/gin-swagger; \
		GO111MODULE="on" go install github.com/swaggo/swag/cmd/swag@latest; \
	else \
		echo "你已经安装过了"; \
	fi \

lint:
	@echo -e ${CGREEN} "make lint........................................................."
	@if [ ${GOCILINT_EXISTS} -eq 0 ]; then  \
		echo "未检测到golangci-lint........................................................."; \
		echo "> 1. 若安装过但提示无法找到 golangci-lint，尝试 source ~/.bashrc 或 ~/.zshrc"; \
		echo "> 2. 若确实未安装过，请执行 make install-lint 安装"; \
		echo "> 关于golangci-lint更多信息，请查阅 https://golangci-lint.run/usage/quick-start"; \
	else \
		golangci-lint run; \
	fi

fmt: $(CHANGE_FILES)
	@if [ -n "$(GOFMT_EXIST)" -a -n "$^" ]; then \
		$(GOFMT_CMD) -w -l -local $(GOMOD) $^; \
	fi
