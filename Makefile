.DEFAULT_GOAL = help
.PHONY: clean ensure-build-dir

PROJECT_DIR = $(shell pwd)
PROJECT_BIN = $(PROJECT_DIR)/bin
GOLANGCI_LINT = $(PROJECT_BIN)/golangci-lint

BUILD_DIR = $(PROJECT_DIR)/build
PLATFORMS = linux/amd64 linux/arm64 darwin/arm64

##
## Build:
build: ## - Build all binaries for all platforms
build: build-server build-client

build-server: ## - Build server for all platforms
build-server: clean ensure-build-dir
	$(foreach platform,$(PLATFORMS),\
		$(eval OS := $(word 1,$(subst /, ,$(platform)))) \
		$(eval ARCH := $(word 2,$(subst /, ,$(platform)))) \
		CGO_ENABLED=0 GOOS=$(OS) GOARCH=$(ARCH) go build -o $(BUILD_DIR)/server-$(OS)-$(ARCH) ./cmd/server/main.go;)

build-client: ## - Build client for all platforms
build-client: clean ensure-build-dir
	$(foreach platform,$(PLATFORMS),\
		$(eval OS := $(word 1,$(subst /, ,$(platform)))) \
		$(eval ARCH := $(word 2,$(subst /, ,$(platform)))) \
		CGO_ENABLED=0 GOOS=$(OS) GOARCH=$(ARCH) go build -o $(BUILD_DIR)/client-$(OS)-$(ARCH) ./cmd/client/main.go;)

clean:
	rm -rf $(BUILD_DIR)

ensure-build-dir:
	mkdir -p $(BUILD_DIR)

##
## Run:
run-server: ## - Run server
	cd cmd/server && go run main.go

run-client: ## - Run client
	cd cmd/client && go run main.go

##
## Docker:
docker-build: docker-server-build docker-client-build ## - Build all docker images

docker-server-build: ## - Build server docker image
	docker build -t tz444/server -f docker/Dockerfile --build-arg APP=server --platform linux/arm64 .
	
docker-client-build: ## - Build client docker image
	docker build -t tz444/client -f docker/Dockerfile --build-arg APP=client --platform linux/arm64 .

##
## Other:
help: ## - Show help message
	@grep -F -h "##" $(MAKEFILE_LIST) | grep -F -v grep -F | sed -e 's/\\$$//' | awk 'BEGIN {FS = ":*[[:space:]]*##[[:space:]]*"}; \
	{ \
		if($$2 == "") \
			printf ""; \
		else if($$0 ~ /^#/) \
			printf "\n%s\n", $$2; \
		else if($$1 == "") \
			printf "     %-20s%s\n", "", $$2; \
		else \
			printf "\n    \033[34m%-20s\033[0m %s\n", $$1, $$2; \
	}'
	@echo "" # blank line at the end

