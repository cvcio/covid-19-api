REGISTRY=reg.plagiari.sm
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
	docker build -t $(REGISTRY)/$(PROJECT):$(VERSION) -f Dockerfile .

docker-push:
	docker push $(REGISTRY)/$(PROJECT):$(VERSION)

docker-release: linux vendor
	docker build -t $(REGISTRY)/$(PROJECT):$(TAG) -f Dockerfile .
	docker tag $(REGISTRY)/$(PROJECT):$(TAG) $(REGISTRY)/$(PROJECT):$(BRANCH) 
	docker push $(REGISTRY)/$(PROJECT):$(TAG)
	docker push $(REGISTRY)/$(PROJECT):$(BRANCH) 

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
	docker volume rm -f infoflow_data_mongo

clean-db: clean-mongo

prod:
	go mod vendor
	cp cmd/${APP}/Dockerfile.$(APP) .
	docker build -f Dockerfile.${APP} --rm -t ${APP}:$(TAG) .
	@chmod +x cmd/${APP}/deploy.sh
	NAME=${APP} REPO=$(REGISTRY) PROJECT=$(PROJECT) CIRCLE_SHA1=$(TAG) CIRCLE_BRANCH=$(BRANCH) cmd/${APP}/deploy.sh
	rm Dockerfile.${APP}

# This included makefile should define the 'custom' target rule which is called here.
include $(INCLUDE_MAKEFILE)

.PHONY: release
release: custom 
