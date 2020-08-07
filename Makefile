GOARCH = amd64
RATELIMIT_BINARY = ratelimit
RATELIMITCTL_BINARY = ratelimitctl

build-darwin: build-ratelimit-darwin build-ratelimitctl-darwin

build-ratelimit-darwin:
		GOOS=darwin GOARCH=${GOARCH} go build ${LDFLAGS} -o ${RATELIMIT_BINARY}-darwin-${GOARCH} cmd/ratelimit/main.go 

build-ratelimit-linux:
		GOOS=linux GOARCH=${GOARCH} go build ${LDFLAGS} -o ${RATELIMIT_BINARY}-linux-${GOARCH} cmd/ratelimit/main.go 

build-ratelimitctl-darwin:
		GOOS=darwin GOARCH=${GOARCH} go build ${LDFLAGS} -o ${RATELIMITCTL_BINARY}-darwin-${GOARCH} cmd/ratelimitctl/main.go 

build-ratelimitctl-linux:
		GOOS=linux GOARCH=${GOARCH} go build ${LDFLAGS} -o ${RATELIMITCTL_BINARY}-linux-${GOARCH} cmd/ratelimitctl/main.go 

