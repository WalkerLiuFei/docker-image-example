############################
# STEP 1 build executable binary
############################
#
FROM golang:alpine AS builder
# Install git.
# Git is required for fetching the dependencies.
#RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/mypackage/myapp/
COPY . .

# Build the binary.
RUN go build GOARCH=amd64 go build -ldflags="-w -s"  -o /go/bin/simple-server
############################
# STEP 2 build a small image
############################
FROM scratch
# Copy our static executable.
COPY --from=builder /go/bin/simple-server /go/bin/simple-server
# Run the hello binary.
ENTRYPOINT ["/go/bin/simple-server"]