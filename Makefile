# Variables
APP_NAME=ipinfo
DOCKER_IMAGE=ipinfo

# Default to linux/amd64 unless overridden
TARGETOS ?= linux
TARGETARCH ?= amd64

.PHONY: help all build clean docker docker-run


help:
	@echo "Usage: make <target>"
	@echo ""
	@echo "Targets:"
	@echo "  build         Build the Go binary (static, cross-compiled)"
	@echo "  clean         Remove the binary"
	@echo "  docker        Build the Docker image"
	@echo "  docker-run    Run the Docker container"
	@echo "  help          Show this help message"
	@echo ""
	@echo "Environment variables:"
	@echo "  TARGETOS      Target OS for build (default: linux)"
	@echo "  TARGETARCH    Target architecture for build (default: amd64)"

all: build

build:
	GOOS=$(TARGETOS) GOARCH=$(TARGETARCH) CGO_ENABLED=0 go build -o $(APP_NAME) main.go

clean:
	rm -f $(APP_NAME)

docker:
	docker build --build-arg TARGETOS=$(TARGETOS) --build-arg TARGETARCH=$(TARGETARCH) -t $(DOCKER_IMAGE) .

docker-run:
	docker run -d -p 80:80 --name $(DOCKER_IMAGE) $(DOCKER_IMAGE)
