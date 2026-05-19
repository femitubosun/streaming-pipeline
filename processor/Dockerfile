FROM golang:1.25.5-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY cmd ./cmd
COPY internal ./internal
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o processor ./cmd

FROM gcr.io/distroless/static-debian12:nonroot

WORKDIR /app
COPY --from=builder /app/processor .

USER nonroot:nonroot

ENTRYPOINT ["./processor"]
