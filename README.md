# Cryptopals Crypto Challenges in Go

A structured implementation and cryptanalysis suite for the Cryptopals Crypto Challenges, written in Go. The goal of this project is to build reusable cryptographic primitives and verify attacks through rigorous, automated test suites.

---

## Project Architecture

The repository strictly separates reusable cryptographic components from specific challenge implementations to ensure scalability and maintainability:

*   **`pkg/cipher/`**: Core primitives and block/stream cipher operations, operating purely on `[]byte`.
*   **`pkg/analysis/`**: Statistical tools, frequency metrics, and cryptanalysis utilities.
*   **`testdata/`**: Static testing files and corpus data (e.g., frequency maps, encrypted payloads) recognized natively by the Go toolchain.
*   **`set1/`, `set2/`, etc.**: End-to-end challenge validations and adversarial attack simulations orchestrated as Go tests.

---

## Technical Design Principles

*   **Zero String Abuse:** Cryptographic pipelines work on raw bytes (`[]byte`). Encodings (Hex, Base64) are handled strictly at I/O boundaries.
*   **Memory Optimization:** Explicit preallocation of buffers (`make([]byte, length)`) to eliminate Garbage Collector thrashing and dynamic reallocation.
*   **Idiomatic Error Propagation:** Graceful error handling (`error`) over execution panics.
*   **Headless Analysis:** Separation of calculation from standard output (`stdout`), allowing clean benchmarking and parallel test execution.

---

## Getting Started

### Prerequisites
*   Go 1.25.0
*   Linux / WSL environment recommended

### Running the Test Suites

Execute all challenge solutions:
```bash
go test -v ./set...
