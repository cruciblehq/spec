# spec

Normative data models and wire contracts for Crucible.

## Purpose

This project defines the data models that multiple projects must agree on,
including schema, validation behavior, reference parsing, manifest and plan
semantics.

## Package Layout

- `agent`: Protocol for talking to the agent on each machine
- `manifest`: Resource manifest types
- `reference`: Resource reference grammar and version constraints
- `registry`: Registry contract types
- `trust`: Machine identity, enrollment, and mutual TLS
