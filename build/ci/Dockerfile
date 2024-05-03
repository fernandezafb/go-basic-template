## Build Rest API ##

# Use latest golang base image
FROM golang:latest as build-env

ARG ENV_FILE=.env.example

# Set the current working directory & copy source to it inside the container
WORKDIR /app

# Download all dependencies defined in the go.mod & go.sum files
## Done separately to leverage caching layers
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY ${ENV_FILE} .env
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/package cmd/main.go

## Run Rest API ##

FROM alpine:latest
RUN apk --no-cache add ca-certificates

ARG VERSION="dev"

LABEL Name="go-basic-template" \
      Version="$VERSION" \
      Vendor="Bernia Labs" \
      Descrition="Go Basic Template API" \
      Maintainer="https://github.com/fernandezafb"

WORKDIR /run/app

# Copy pre-build binary file from the 1st stage
COPY --from=build-env /app/build/package .

# Expose 9450 port to the outside world
EXPOSE 9540

# Command to run the binary
CMD ["./main"]