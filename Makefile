mock:
	@go get github.com/golang/mock/gomock
	@go install github.com/golang/mock/mockgen
	@~/go/bin/mockgen -source=domain/product.go -destination=product/mock/product_mock.go -package=mock

test:
	go test ./...

run:
	@docker-compose up -d	
	@go run api/main.go
