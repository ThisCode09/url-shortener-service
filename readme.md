# Go URL Shortener

A high-performance, lightweight URL shortener built from scratch in Go using standard libraries. Designed following Clean Architecture principles, it features modular layers, dependency injection, wildcard path routing, and a production-ready lifecycle with graceful shutdown.

---

## 🏗️ Project Architecture

The project strictly separates concerns into distinct layers, ensuring high testability, maintainability, and zero reliance on heavy third-party frameworks.

```text
url-shortener/
├── cmd/
│   └── server/
│       └── main.go          # Application entrypoint, wiring, and lifecycle management
├── internal/
│   ├── handlers/            # HTTP layer (request parsing, response rendering, status codes)
│   ├── shortener/           # Business logic layer (URL validation, short code generation)
│   └── storage/             # Data persistence layer (JSON-backed thread-safe storage)
├── .gitignore
└── go.mod