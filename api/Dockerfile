FROM alpine:3.12
RUN apk --no-cache add ca-certificates

COPY /bin/outdoor-api /go/bin/outdoor-api

WORKDIR /go/bin/
# ENV PORT 8080
# EXPOSE 8080
# RUN outdoorapi binary
ENTRYPOINT ["/go/bin/outdoor-api"]
