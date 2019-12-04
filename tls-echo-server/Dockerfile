############################
# STEP 1 build executable binary
############################
#
FROM golang:alpine AS builder

# Create appuser.
RUN adduser -D -g '' appuser

WORKDIR $GOPATH/src/docker-image-example/echo-server
COPY . .

# Build the binary. cgo must be disable!
RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o   /go/bin/echo-server
RUN chmod +x /go/bin/echo-server
############################
# STEP 2 build a small image
############################git
FROM scratch
# Copy user from builder
COPY --from=builder /etc/passwd /etc/passwd

# Copy private key
COPY --from=builder $GOPATH/src/docker-image-example/echo-server/asset /go/bin/asset
# Copy our static executable.
COPY --from=builder /go/bin/echo-server /go/bin/echo-server
# Use unprivileged user
USER appuser

EXPOSE 8888/tcp
# Run the echo-server binary.
ENTRYPOINT ["/go/bin/echo-server"]