# Distributed Streaming Pipeline — Project Spec

## Overview

A high-throughput event streaming pipeline simulating real-time financial transaction processing. Polyglot microservices communicating over gRPC, connected via Kafka, hosted locally on AWS-emulated infrastructure.

## Tech Stack

🚜 Infra: Floci (local AWS) + Pulumi (IaC)

💬 Messaging: MSK (Kafka/Redpanda)

🤖 Services: Go, Rust, TypeScript

☎️ Comms: gRPC + Protobuf

## Services

1. Event Producer (Rust) Generates synthetic financial transactions at high volume. Configurable rate/burst. Publishes to Kafka.

2. Stream Processor (Go) Consumes raw events, enriches and validates them (fraud signals, balance checks). Publishes processed events to a separate topic.

3. API Layer (TypeScript) Exposes a REST/gRPC API for querying aggregated pipeline state — throughput, error rates, flagged transactions.

4. Observability (Go) Collects OpenTelemetry traces and metrics across services. Exposes a simple dashboard endpoint.

## Data Flow

```
Rust Producer → [raw-events topic]
→ Go Processor → [processed-events topic]
→ TS API (query layer)
→ Observability sidecar (all services)
```


## Goals

- Sustain 500k–1M events/sec locally
- End-to-end trace visibility across services
- Everything provisioned via Pulumi, torn down and reproduced in one command


## Out of Scope

- Auth
- Persistence layer
- Real AWS deployment
