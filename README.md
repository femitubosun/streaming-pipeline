# Distributed Streaming Pipeline

A high-throughput event streaming pipeline simulating real-time financial transaction processing. Polyglot microservices communicating via Kafka, with an HTTP API for querying pipeline state. Hosted locally via Docker Compose.

## Tech Stack

🐳 Infra: Docker Compose + Pulumi (IaC)

💬 Messaging: Kafka (Redpanda)

🤖 Services: Go, Rust, TypeScript

☎️ Comms: Kafka (inter-service), gRPC (API → Observability internal)

## Architecture

### Communication Patterns

- **Kafka** — All inter-service communication. Producer → Processor → (optional consumers). Temporal decoupling, fire-and-forget, 500k+ events/sec.
- **gRPC** — Internal only. API Layer queries Observability service for metrics/traces.
- **HTTP** — External API surface. Clients interact with the pipeline via REST.

### Why not gRPC everywhere?

The data flow is purely event-driven. The Producer doesn't know the Processor exists. The Processor doesn't know the API exists. Adding gRPC between services would introduce temporal coupling where none is needed. gRPC is reserved for request/response patterns: API querying Observability for dashboard data.

## Services

1. **Event Producer (Rust)** — Generates synthetic financial transactions at high volume. Configurable rate/burst. Publishes to `raw-events` topic.

2. **Stream Processor (Go)** — Consumes raw events, enriches and validates them (fraud signals, balance checks). Publishes processed events to `processed-events` topic.

3. **API Layer (TypeScript/Hono)** — Exposes HTTP REST API for querying pipeline state: throughput, error rates, flagged transactions. Internally queries Observability via gRPC.

4. **Observability (Go)** — Collects OpenTelemetry traces and metrics from all services via OTLP. Exposes gRPC query interface for the API layer.

## Data Flow

```
┌─────────────┐     Kafka      ┌─────────────┐     Kafka      ┌─────────────┐
│   Producer  │ ──raw-events──►│  Processor  │ ─processed-ev─►│   (topic)   │
│    (Rust)   │                │    (Go)     │                │             │
└─────────────┘                └─────────────┘                └─────────────┘
       │                              │
       │ OTLP traces                 │ OTLP traces
       ▼                              ▼
┌─────────────────────────────────────────────┐
│           Observability (Go)                 │
│  Collects traces/metrics from all services   │
│  Exposes gRPC query interface                │
└─────────────────────────────────────────────┘
       ▲
       │ gRPC queries
       │
┌─────────────────────────────────────────────┐
│           API Layer (TypeScript)             │
│  HTTP REST for external clients              │
│  gRPC client to Observability                │
└─────────────────────────────────────────────┘
```

## Module Structure

Each service is an independent module with its own dependency management:

```
.
├── producer/          # Rust crate
│   └── Cargo.toml
├── processor/         # Go module
│   └── go.mod
├── observability/     # Go module
│   └── go.mod
├── api/              # TypeScript (Bun)
│   └── package.json
├── proto/            # Shared protobuf contracts
│   └── go.mod
├── go.work           # Go workspace (local dev, gitignored)
└── docker-compose.yml
```

## Goals

- Sustain 500k–1M events/sec locally
- End-to-end trace visibility across services
- Everything provisioned via Docker Compose, torn down and reproduced in one command
- Clean separation: Kafka for events, gRPC for internal queries, HTTP for external API

## Out of Scope

- Auth
- Persistence layer
- Cloud deployment
- Service mesh
