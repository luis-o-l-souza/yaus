.PHONY: build
build:
	go build -o ./build/yaus

dev: 
	make build && ./build/yaus
