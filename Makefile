# ---------------------------------------------------------------------------------------------- Variables declaration #

# ----------------------------------------------------------------------------------------------------- Docker-Compose #
PROFILE = "development"

# ----------------------------------------------------------------------------------------------------------------- Go #
CGO_ENABLED = 0
GOARCH      = "amd64"
GOOS        = "linux"
VERSION     = `git describe --tags --abbrev=0`
TIMESTAMP   = `date +"%a %d %b %Y %H:%M:%S"`
FLAGS       = "-X 'main.version=$(VERSION)' -X 'main.timestamp=$(TIMESTAMP)'"

# --------------------------------------------------------------------------------------------------- Makefile targets #

# ------------------------------------------------------------------------------------------------------------- Docker #
.PHONY: docker/build
docker/build:
	docker build \
		   --build-arg CGO_ENABLED=$(CGO_ENABLED) \
		   --build-arg GOARCH=$(GOARCH) \
		   --build-arg GOOS=$(GOOS) \
           --build-arg VERSION="$(VERSION)" \
           --build-arg TIMESTAMP="$(TIMESTAMP)" \
           -f dep/Dockerfile \
           -t erebus:"$(VERSION)" \
           .

# ----------------------------------------------------------------------------------------------------- Docker-Compose #
.PHONY: docker-compose/down
docker-compose/down:
	docker-compose -f dep/docker-compose.yaml --profile $(PROFILE) down --volumes

.PHONY: docker-compose/log
docker-compose/log:
	docker-compose -f dep/docker-compose.yaml logs -f --tail 10

.PHONY: docker-compose/pull
docker-compose/pull:
	docker-compose -f dep/docker-compose.yaml --profile $(PROFILE) pull

.PHONY: docker-compose/up
docker-compose/up:
	docker-compose -f dep/docker-compose.yaml --profile $(PROFILE) up -d

# ----------------------------------------------------------------------------------------------------------------- Go #
.PHONY: go/build
go/build:
	CGO_ENABLED=$(CGO_ENABLED) GOARCH=$(GOARCH) GOOS=$(GOOS) go build -o art/$(GOOS)/erebus -a -ldflags=$(FLAGS) -v .

.PHONY: go/run
go/run:
	go run -ldflags=$(FLAGS) -v .

# ----------------------------------------------------------------------------------------------------------- protobuf #
.PHONY: protobuf/generate
protobuf/generate:
	protoc --go_out=. \
		   --go_opt=paths=source_relative \
		   --go-grpc_out=. \
		   --go-grpc_opt=paths=source_relative \
		   pkg/proto/*/*.proto
