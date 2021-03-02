# STEP 1: Build executable binary
FROM golang:1.15.3-alpine3.12 as builder
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

RUN apk update && apk upgrade && \
  apk add --no-cache ca-certificates git openssh-client

RUN mkdir -p /api
WORKDIR /api
ADD . /api
RUN go get -u github.com/swaggo/swag/cmd/swag@v1.6.7
RUN swag init
RUN wget -O - -q https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s v1.28.3
RUN ./bin/golangci-lint run --timeout=30m ./...
RUN go mod download
RUN go test ./...
RUN go build -o api

# STEP 2: Build a small image start from scratch
FROM scratch

RUN apk update && apk upgrade && \
  apk add --no-cache ca-certificates tzdata && \
  rm -rf /var/cache/*

# Copy our static executable
COPY --from=builder /api/api .
COPY --from=builder /api/docs /docs

# Expose port 8000
EXPOSE 8000

# Command is run when container starts
ENTRYPOINT ["./api"]
