FROM alpine:3.12
RUN apk --no-cache add ca-certificates
WORKDIR /the/workdir/path
RUN ls bin
COPY bin/go-outdoors /go/bin/go-outdoors


WORKDIR /go/bin/
# ENV PORT 8080
# EXPOSE 8080
# RUN go-outdoors binary
ENTRYPOINT ["/go/bin/go-outdoors"]
