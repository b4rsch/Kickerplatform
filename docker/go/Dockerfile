FROM golang:1.17.2 as builder

RUN mkdir /src
WORKDIR /src

COPY ./ ./
RUN go mod download
RUN GOFLAGS=-mod=mod go generate ./...
RUN CGO_ENABLED=0 GOOS=linux go build -o /src/server -installsuffix cgo -ldflags '-w'

FROM alpine:latest

RUN mkdir /app
WORKDIR /app
COPY --from=builder /src/server /app

ENTRYPOINT ["/app/server"]
