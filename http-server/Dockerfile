############################
# STEP 1 build executable binary
############################
#
FROM golang:alpine AS builder
# Install git.
# Git is required for fetching the dependencies.
#RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/docker-image-example/http-server
COPY . .

# Build the binary.
RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o   /go/bin/http-server
RUN chmod +x /go/bin/http-server
############################
# STEP 2 build a small image
############################git
FROM scratch
# Copy our static executable.
COPY --from=builder /go/bin/http-server /go/bin/http-server
EXPOSE 8888/tcp
# Run the http-server binary.
ENTRYPOINT ["/go/bin/http-server"]