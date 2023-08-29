.PHONY: run

run:
	swag init
	swag fmt
	go run main.go