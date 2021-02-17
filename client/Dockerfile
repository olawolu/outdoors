FROM alpine:3.12
RUN apk --no-cache add ca-certificates

COPY bin/go-outdoors /go/bin/go-outdoors
RUN cd go/bin && mkdir /client
ADD client/* /go/bin/client/
WORKDIR /go/bin/
# ENV PORT 8080
# EXPOSE 8080
# RUN go-outdoors binary
ENTRYPOINT ["/go/bin/go-outdoors"]
