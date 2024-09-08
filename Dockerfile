#
# Build go project
#
FROM golang:1.19-alpine as go-builder

WORKDIR /roxprox-ratelimit

COPY . .

RUN apk add -u -t build-tools curl git && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ratelimit cmd/ratelimit/main.go && \
    apk del build-tools && \
    rm -rf /var/cache/apk/*

#
# Runtime container
#
FROM alpine:3.19.4  

WORKDIR /app

RUN apk --no-cache add ca-certificates bash curl

COPY --from=go-builder /roxprox-ratelimit/ratelimit .

ENTRYPOINT ["./ratelimit"]
