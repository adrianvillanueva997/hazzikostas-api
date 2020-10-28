# Multistage builder image
# Builder -> executable

# Builder stage
# Go build
FROM golang:1.13-alpine as build-go
RUN apk update && apk add --no-cache gcc libc-dev
WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o main .

# Javascript build
FROM node:14.15.0-alpine3.12 as build-js
RUN apk --no-cache update && apk add make g++ && rm -rf
WORKDIR /build_js
COPY client .
RUN npm install && npm run build

# Production stage
FROM alpine:latest
WORKDIR /app
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GIN_MODE=release
COPY --from=build-go /build/main .
COPY --from=build-js /build_js/build ./client/build
COPY .env .
EXPOSE 5000
ENTRYPOINT ./main