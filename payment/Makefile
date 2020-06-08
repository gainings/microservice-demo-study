DOCKERENV_PATH=/.dockerenv
GO_CONTAINER_NAME=api
ifneq ("$(wildcard $(DOCKERENV_PATH))","$(DOCKERENV_PATH)")
	DOCKER_COMPOSE_GO_CMD= docker-compose exec $(GO_CONTAINER_NAME)
endif

.PHONY: build
help: ## この文章を表示します。
	# http://postd.cc/auto-documented-makefile/
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

LDFLAGS = -ldflags "-X main.VERSION=${VERSION} -X main.COMMIT=${COMMIT} -X main.BRANCH=${BRANCH}"
build: generate ## webapiをbuildします
	GO111MODULE=on GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build ${LDFLAGS} -o dist/app ./cmd/api

generate: proto ## go generate を実行します
	$(DOCKER_COMPOSE_GO_CMD) go generate ./...

proto: ## protoファイルをコンパイルします
	protoc -I=proto --go_out=controller/grpc proto/payment.proto
