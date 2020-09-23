
deps:
	@echo "\nInstalling dependencies"
	@go get ./...

run: 
	@echo "\nServing app"
	@go run main.go