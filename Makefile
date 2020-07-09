generate-mock:
	@go get github.com/golang/mock/gomock
	@go install github.com/golang/mock/mockgen
	@~/go/bin/mockgen -source=core/product/interface.go -destination=application/product/mock/product_mock.go -package=mock

test:
	go test ./...

run:
	@docker-compose up -d	
	@go run cmd/main.go
