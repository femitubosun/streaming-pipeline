FROM golang:1.25.5-alpine AS builder

WORKDIR /app
COPY observability/go.mod observability/go.sum ./
COPY proto /proto
RUN go mod download

COPY observability/cmd ./cmd
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o observability ./cmd

FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=builder /app/observability /observability

USER nonroot:nonroot

ENTRYPOINT ["/observability"]
