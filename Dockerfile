# Step 1
FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o app
ENTRYPOINT [ "/app/app" ]