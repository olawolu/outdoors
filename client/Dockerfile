FROM alpine:3.12
RUN apk --no-cache add ca-certificates

WORKDIR /go/src/github.com/olawolu/outdoors

COPY bin/go-outdoors /go/bin/go-outdoors

# ENV PORT 8080
# EXPOSE 8080
# RUN go-outdoors binary
ENTRYPOINT ["/go/bin/go-outdoors"]
