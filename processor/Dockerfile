FROM golang:1.25.5-alpine AS builder

WORKDIR /app
COPY processor/go.mod processor/go.sum ./
COPY proto /proto
RUN go mod download

COPY processor/cmd ./cmd
COPY processor/internal ./internal
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o processor ./cmd

FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=builder /app/processor /processor
USER nonroot:nonroot

ENTRYPOINT ["/processor"]
