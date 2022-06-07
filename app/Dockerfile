ARG GO_VERSION=1.16.6

FROM golang:${GO_VERSION}-alpine AS builder

# Get dependencies directly
RUN go env -w GOPROXY=direct
# It's necessary to install git when building the app
RUN apk add --no-cache git
# Security certificates
RUN apk --no-cache add ca-certificates && update-ca-certificates

WORKDIR /src
# Copy both files to the main directory
COPY ./go.mod ./go.sum ./
# Download and install dependencies
RUN go mod download
# Dependencies for swagger
RUN go get -u github.com/swaggo/http-swagger

COPY ./ ./
# Disable the C++ compiler
RUN CGO_ENABLED=0 go build \
    -installsuffix 'static' \
    # Select the name
    -o /go-bootcamp


# This image will run the app
FROM scratch as runner

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY ./config/.env ./config/

COPY /files/characters.csv ./files/characters.csv

COPY ./docs/ ./docs/

COPY --from=builder /go-bootcamp /go-bootcamp


EXPOSE 5050
# Call the server
ENTRYPOINT ["/go-bootcamp"]