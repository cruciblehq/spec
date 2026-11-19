# spec

Normative data models and wire contracts for Crucible.

## Purpose

This project defines what Crucible programs must agree on: protocols, contracts,
and formats. This includes schemas, validation behavior, reference parsing, and
// the semantics of manifests and plans.

## Package Layout

- `agent`: Protocol between a client and the agent on each machine
- `manifest`: Resource manifests, deployment plans, and deployment state
- `reference`: Resource reference grammar and version constraints
- `registry`: Registry API contract and types
- `trust`: Machine identity, enrollment, revocation, and mutual TLS
