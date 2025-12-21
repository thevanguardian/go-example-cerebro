# Cerebro Sentinel

<p align="left">
  <img alt="Go Version" src="https://img.shields.io/badge/go-1.25-blue.svg" />
  <img alt="License" src="https://img.shields.io/badge/license-MIT-green.svg" />
  <img alt="Build Status" src="https://img.shields.io/badge/build-in%20progress-lightgrey.svg" />
  <img alt="Security" src="https://img.shields.io/badge/security-best%20practices-blueviolet.svg" />
</p>

Cerebro Sentinel is a Go-based educational project designed to demonstrate modern backend service development using secure defaults, clear architectural boundaries, and idiomatic Go practices.

The project uses a restrained Marvel X-Men theme to provide memorable context while maintaining professional, production-oriented design and implementation.

---

## Generative AI Disclosure

Portions of this project were developed with the assistance of generative AI tools, including the use of deliberate prompt engineering for documentation generation and basic code assistance.

All architectural decisions, implementation choices, review, validation, and final integration were performed by a human author. The resulting codebase reflects human engineering judgment and responsibility.

---

## Purpose

This project exists to:

- Demonstrate modern Go service design without heavy frameworks
- Emphasize security, correctness, and maintainability
- Provide a structured learning path for real-world Go development
- Serve as a reference for best practices in small-to-medium Go services

This is not intended to be a feature-complete product, but a well-reasoned, well-documented educational system.

---

## Project Status

**Active development**

The project is being built incrementally in defined phases, beginning with foundational setup and progressing through authentication, authorization, persistence, and observability.

---

## Requirements

- Go 1.25 or newer
- Docker (optional)
- Visual Studio Code with the Dev Containers extension (optional)

---

## Getting Started

### Clone the repository

```bash
git clone https://github.com/thevanguardian/go-example-cerebro
cd go-example-cerebro 
```
Install dependencies
```bash
go mod tidy
```

Run the service
```bash
make run
```
The service will start on:
http://localhost:8080


Health check:
```bash
curl http://localhost:8080/healthz
```
