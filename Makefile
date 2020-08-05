GOARCH = amd64
RATELIMIT_BINARY = ratelimit

build-darwin: build-ratelimit-darwin

build-ratelimit-darwin:
		GOOS=darwin GOARCH=${GOARCH} go build ${LDFLAGS} -o ${RATELIMIT_BINARY}-darwin-${GOARCH} cmd/ratelimit/main.go 

build-ratelimit-linux:
		GOOS=linux GOARCH=${GOARCH} go build ${LDFLAGS} -o ${RATELIMIT_BINARY}-linux-${GOARCH} cmd/ratelimit/main.go 

