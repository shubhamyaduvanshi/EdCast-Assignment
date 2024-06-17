DOCKER_IMAGE_NAME = go-calc-db
DOCKER_TAG = v1.0
CONTAINER_NAME = go-calc-db-ct
DOCKERFILE = Dockerfile


all: build

install:
	dep ensure

build: 
	install build_linux

build_linux:
	mkdir -p ./build
	GOOS=linux go build -o ./build/go-calc main.go

# Build the Docker image
docker-build:
	docker build -t $(DOCKER_IMAGE_NAME):$(DOCKER_TAG) -f $(DOCKERFILE) .

# Run the Docker container
run:
	docker run --name $(CONTAINER_NAME) -d -p 8080:8080 $(DOCKER_IMAGE_NAME):$(DOCKER_TAG)

clean:
	rm -rf ./build

run_local:
	go run ./main.go

fmt:
	go fmt ./

