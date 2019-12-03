############################
# STEP 1 build executable binary
############################
#
FROM golang:alpine AS builder
# Install git.
# Git is required for fetching the dependencies.
#RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/
COPY . .

# Build the binary.
RUN go build   -ldflags="-w -s"  -o /go/bin/simple-server
############################
# STEP 2 build a small image
############################git
FROM scratch
# Copy our static executable.
COPY --from=builder /go/bin/simple-server /go/bin/simple-server
# make simple-server excutable
RUN ["chmod", "+x", "/go/bin/simple-server"]
# Run the hello binary.
ENTRYPOINT ["/go/bin/simple-server"]