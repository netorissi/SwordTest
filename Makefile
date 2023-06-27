.PHONY: run
run: ## run it will instance server
	@echo Building api
	docker-compose up

.PHONY: stop
stop: ## stop instance server
	@echo Stoping api
	docker-compose down

.PHONY: mock
mock: ## generate all mocks interfaces
	rm -rf ./mocks/*
	go generate ./...

.PHONY: test
test: ## running unit tests with coverage
	go test -coverprofile=coverage.out ./...

.PHONY: test-html
test-html: ## generate html view unit tests with coverage
	go tool cover -html=coverage.out

.PHONY: swag
swag: ## generate docs swagger
	swag init --pd
