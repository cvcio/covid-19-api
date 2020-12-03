PROJECT=covid-19-api
TAG:=$(shell git rev-parse HEAD)
BRANCH:=$(shell git rev-parse --abbrev-ref HEAD)

ifeq (,$(VERSION))
VERSION=latest
endif

keys:
	openssl genpkey -algorithm RSA -out private.pem -pkeyopt rsa_keygen_bits:2048

tools:
	go get github.com/oxequa/realize
	go get github.com/golangci/golangci-lint

run:
	realize start

run-api:
	realize start -n covid-19

test:
	go test -v ./...

.PHONY: linux
linux: GOOS := linux
linux: GOARCH := amd64

.PHONY: vendor
vendor: 
	go mod vendor

.PHONY: docker
docker: linux vendor
	docker build -t $(PROJECT):$(VERSION) -f Dockerfile .

db-start:
	docker-compose up -d

db-logs:
	docker-compose logs -f

db: db-start db-logs

db-stop:
	docker-compose stop

lint:
	golangci-lint run -e vendor

clean-mongo:
	docker-compose stop mongo
	docker-compose rm -f -v mongo
	docker volume rm -f data_mongo

clean-db: clean-mongo

# This included makefile should define the 'custom' target rule which is called here.
include $(INCLUDE_MAKEFILE)

.PHONY: release
release: custom 
