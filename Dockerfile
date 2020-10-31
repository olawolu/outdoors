FROM golang:1.15-alpine AS builder

# Install git for the dependecies
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

WORKDIR /go/src/github.com/olawolu/outdoors
COPY go.mod go.sum ./

# Fetch dependencies
RUN go mod download
RUN go mod verify

COPY . .

# Build the binary
WORKDIR /go/src/github.com/olawolu/outdoors/cmd/outdoors
RUN CGO_ENABLED=0 go build -o /go/bin/outdoorapi

# build a small image
FROM alpine:3.12
RUN apk --no-cache add ca-certificates

# copy the static executable
COPY --from=builder /go/bin/outdoorapi /go/bin/outdoorapi
# COPY --from=builder /go/src/github.com/olawolu/outdoors/.env . 
EXPOSE 8080
# RUN outdoorapi binary
ENTRYPOINT ["/go/bin/outdoorapi" ]
