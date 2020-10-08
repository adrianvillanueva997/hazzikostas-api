# Multistage builder image
# Builder -> executable
# Builder stage
FROM golang:1.13-alpine as build-env
RUN apk add --no-cache gcc libc-dev

WORKDIR /build

COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .
RUN go build -o main .
# Executable stage
FROM alpine:latest
WORKDIR /app
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GIN_MODE=release
COPY --from=build-env /build/main .
COPY --from=build-env /build/.env .
EXPOSE 3000
ENTRYPOINT ./main