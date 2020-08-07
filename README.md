# Roxprox ratelimit
* Implements an envoy ratelimit GRPC service
* Only in-memory state (github.com/coocood/freecache)
* Config retrieved in same way as roxprox

# Docker image
A docker image is available at Docker Hub: in4it/roxprox-ratelimit

# Commandline usage
* You'll need to either use local storage or s3 storage
```
Usage of ./ratelimit:
  -aws-region string
        AWS region
  -loglevel string
        log level (default "INFO")
  -storage-bucket string
        s3 storage bucket
  -storage-path string
        storage path
  -storage-type string
        storage type (default "local")
```


# Manual build 

```
protoc -I proto/ proto/cache.proto --go_out=plugins=grpc:proto/cache
make  # builds linux & darwin
```