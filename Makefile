mock:
	@go get github.com/golang/mock/gomock
	@go install github.com/golang/mock/mockgen
	@$(GOPATH)/bin/mockgen -source=domain/product.go -destination=product/mocks/product_mock.go -package=mocks

test:
	go test ./...

run:
	@docker-compose up -d	
	@go run api/main.go
