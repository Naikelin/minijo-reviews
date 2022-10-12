# Step 1
FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git
WORKDIR /build
COPY . .
RUN go mod tidy
RUN go build -o app main.go

# Step 2
FROM frolvlad/alpine-glibc
COPY --from=builder /build/app .
ENTRYPOINT [ "./app" ]