FROM golang:1.23-alpine AS builder

WORKDIR /webtester

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o webtester ./webtester/webtester.go

FROM alpine:latest

WORKDIR /webtester

COPY --from=builder /webtester/webtester .

EXPOSE 3000

CMD ["./webtester"]
