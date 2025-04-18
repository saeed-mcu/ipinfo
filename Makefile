# Variables
APP_NAME=ipinfo
DOCKER_IMAGE=ipinfo

# Default to linux/amd64 unless overridden
TARGETOS ?= linux
TARGETARCH ?= amd64

.PHONY: all build clean docker docker-run

all: build

build:
	GOOS=$(TARGETOS) GOARCH=$(TARGETARCH) CGO_ENABLED=0 go build -o $(APP_NAME) main.go

clean:
	rm -f $(APP_NAME)

docker:
	docker build --build-arg TARGETOS=$(TARGETOS) --build-arg TARGETARCH=$(TARGETARCH) -t $(DOCKER_IMAGE) .

docker-run:
	docker run -d -p 80:80 --name $(DOCKER_IMAGE) $(DOCKER_IMAGE)
