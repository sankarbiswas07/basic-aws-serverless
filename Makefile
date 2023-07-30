GOARCH=amd64
GOOS=linux
BINARY_PATH=bin
CGO_ENABLED=0


.PHONY: deploy
deploy: clean build
	serverless deploy

.PHONY: build
build:
	CGO_ENABLED=$(CGO_ENABLED) GOARCH=$(GOARCH) GOOS=$(GOOS) go build -o $(BINARY_PATH)/basic-aws-serverless cmd/main.go

.PHONY: clean
clean:
	rm -rf $(BINARY_PATH)





