FROM golang:1.15.5-alpine3.12
ARG version

# Install certificates and git
RUN apk add --update --no-cache ca-certificates git

# Create and use a directory where our project will be build
RUN mkdir -p /go/src/github.com/cvcio/covid-19-api/
WORKDIR /go/src/github.com/cvcio/covid-19-api/

# COPY go.mod and go.sum files to the workspace
COPY go.mod /go/src/github.com/cvcio/covid-19-api/
COPY go.sum /go/src/github.com/cvcio/covid-19-api/

# Get dependancies - will also be cached if we won't change mod/sum
#RUN go mod vendor

# COPY the source code
COPY cmd/ /go/src/github.com/cvcio/covid-19-api/cmd/
COPY models/ /go/src/github.com/cvcio/covid-19-api/models/
COPY pkg/ /go/src/github.com/cvcio/covid-19-api/pkg/
COPY vendor/ /go/src/github.com/cvcio/covid-19-api/vendor/

WORKDIR /go/src/github.com/cvcio/covid-19-api/cmd/api/
RUN GO111MODULE=on GOFLAGS=-mod=vendor CGO_ENABLED=0 GOOS=linux go build -v -a -installsuffix cgo -o api .

FROM alpine:3.10
RUN apk --no-cache add ca-certificates
WORKDIR /api/
COPY --from=0 /go/src/github.com/cvcio/covid-19-api/cmd/api/api .
ENTRYPOINT ["/api/api"]
