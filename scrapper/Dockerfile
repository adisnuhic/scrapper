# Accept the Go version for the image to be set as a build argument.
# Default to Go 1.14
ARG GO_VERSION=1.14

FROM golang:${GO_VERSION}-alpine AS builder

# Install the Certificate-Authority certificates for the app to be able to make
# calls to HTTPS endpoints.
# Git is required for fetching the dependencies.
RUN apk add --no-cache ca-certificates git 

# Unit tests
#RUN CGO_ENABLED=0 go test -v

# Set the working directory outside $GOPATH to enable the support for modules.
WORKDIR /tmp/app

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
#COPY go.sum .

# Download modules
RUN go mod download

# Import the code
COPY . .

# Build the app
RUN go build -o ./out/scrapper .

# Stage 2 built from stage 1 
# Running container
FROM alpine

WORKDIR /app

# Import the Certificate-Authority certificates for enabling HTTPS.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy executable file
COPY --from=builder /tmp/app/out/scrapper .

# Perform any further action as an unprivileged user.
USER nobody:nobody

# Run the binary program
CMD ["/app/scrapper"]