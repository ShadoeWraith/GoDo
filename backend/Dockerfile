FROM golang:1.25.4-alpine AS builder

ENV CGO_ENABLED=0 GOOS=linux

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN go build -ldflags "-s -w" -o /app/backend-app .

FROM alpine:3.18

RUN apk update 
RUN apk add --no-cache ca-certificates

WORKDIR /root/

COPY --from=builder /app/backend-app .

EXPOSE 8080

CMD ["./backend-app"]