build_app:
	go build -o ./build/app  ./cmd/app/main/main.go 
clear:
	rm -rf ./build/app
rebuild:
	make clear
	make build_app