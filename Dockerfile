FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o echo-server .

FROM gcr.io/distroless/static:nonroot
COPY --from=builder /app/echo-server /echo-server
EXPOSE 8080
ENTRYPOINT ["/echo-server"]
