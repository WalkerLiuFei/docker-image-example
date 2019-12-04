############################
# STEP 1 build executable binary
############################
#
FROM golang:alpine AS builder
# Install git.
# Git is required for fetching the dependencies.
#RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/docker-image-example/echo-server
COPY . .

# Build the binary.
RUN go build -o /go/bin/echo-server
RUN chmod +x /go/bin/echo-server
############################
# STEP 2 build a small image
############################git
FROM scratch
# Copy our static executable.
COPY --from=builder /go/bin/echo-server /go/bin/echo-server
EXPOSE 8888/tcp
# Run the hello binary.
ENTRYPOINT ["/go/bin/echo-server"]