FROM rust:1.95-alpine AS builder
RUN apk add --no-cache bash cmake make gcc g++ openssl-dev pkgconfig musl-dev zlib-dev
WORKDIR /app
COPY Cargo.lock Cargo.toml ./
RUN mkdir -p src && echo 'fn main() {}' > src/main.rs
RUN cargo build --release
COPY src/ src/
RUN cargo build --locked --release

FROM alpine:latest
RUN apk add --no-cache ca-certificates
COPY --from=builder /app/target/release/producer /bin/server
CMD ["/bin/server"]
