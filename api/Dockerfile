FROM alpine:3.12
RUN apk --no-cache add ca-certificates

COPY bin/outdoorapi /go/bin/outdoorapi

WORKDIR /go/bin/
# ENV PORT 8080
# EXPOSE 8080
# RUN outdoorapi binary
ENTRYPOINT ["/go/bin/outdoorapi"]
