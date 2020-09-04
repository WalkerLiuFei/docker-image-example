############################
# STEP 1 build executable binary
############################
#
FROM golang:alpine AS builder
# Install git.
# Git is required for fetching the dependencies.
#RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/docker-image-example/service-discovery-example
COPY . .

# Build the binary.
RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o   /go/bin/service-discovery-example
RUN chmod +x /go/bin/service-discovery-example
############################
# STEP 2 build a small image
############################git
FROM scratch
# Copy our static executable.
COPY --from=builder /go/bin/service-discovery-example /go/bin/service-discovery-example
EXPOSE 8888/tcp
# Run the echo-server binary.
ENTRYPOINT ["/go/bin/service-discovery-example"]