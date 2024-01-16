NAME = zk-socket-server
IMAGE_NAME = zk-socket-server
IMAGE_NAME_MIGRATION_SUFFIX = -migration
IMAGE_VERSION = latest

LOCATION ?= us-west1
PROJECT_ID ?= zerok-dev
REPOSITORY ?= zk-client

export GO111MODULE=on
export BUILDER_NAME=multi-platform-builder
export GOPRIVATE=github.com/zerok-ai/zk-utils-go

sync:
	go get -v ./...

build: sync
	go build -v -o bin/$(NAME) cmd/main.go

run: build
	go run cmd/main.go -c ./config/config.yaml

docker-build: sync
	$(GOOS) $(ARCH) go build -v -o bin/$(NAME) cmd/main.go
	docker build --no-cache $(DockerFile) -t $(IMAGE_PREFIX)$(IMAGE_NAME)$(IMAGE_NAME_SUFFIX):$(IMAGE_VERSION) .

docker-push:
	docker push $(IMAGE_PREFIX)$(IMAGE_NAME)$(IMAGE_NAME_SUFFIX):$(IMAGE_VERSION)


# ------- GKE ------------

# build app image
docker-build-gke: GOOS := GOOS=linux
docker-build-gke: ARCH := GOARCH=amd64
docker-build-gke: IMAGE_PREFIX := $(LOCATION)-docker.pkg.dev/$(PROJECT_ID)/$(REPOSITORY)/
docker-build-gke: DockerFile := -f Dockerfile
docker-build-gke: docker-build

# push app image
docker-push-gke: IMAGE_PREFIX := $(LOCATION)-docker.pkg.dev/$(PROJECT_ID)/$(REPOSITORY)/
docker-push-gke: docker-push

# build and push
docker-build-push-gke: docker-build-gke docker-push-gke
