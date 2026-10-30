# Go URL Shortener

A high-performance, lightweight URL shortener built from scratch in Go using standard libraries. Designed following Clean Architecture principles, it features modular layers, dependency injection, wildcard path routing, and a production-ready lifecycle with graceful shutdown.

## 🏗️ Project Architecture

The project separates concerns into distinct layers to support testability and maintainability, without relying on heavy third-party frameworks.

```text
url-shortener/
├── cmd/
│   └── server/
│       └── main.go          # Application entry point, wiring, and lifecycle management
├── internal/
│   ├── handlers/            # HTTP layer (request parsing, response rendering, status codes)
│   ├── shortener/           # Business logic (URL validation, short code generation)
│   └── storage/             # Data persistence (JSON-backed, thread-safe storage)
├── .gitignore
└── go.mod
```

### Layer Breakdown

1. **HTTP Handlers (`internal/handlers`)**
   - Uses Go 1.22+ enhanced `net/http` path wildcards (`{url...}`, `{code}`) for browser-based testing without external tools.
   - Is decoupled from storage and communicates exclusively through the service layer.

2. **Business Logic (`internal/shortener`)**
   - Contains the core rules for shortening URLs, enforcing uniqueness, and resolving short codes.

3. **Storage (`internal/storage`)**
   - Implements thread-safe JSON persistence, automatically serializing and saving state locally.

## 🚀 Key Features

- **Zero framework bloat:** Built entirely with Go’s native `net/http` standard library and a custom multiplexer.
- **Constructor injection:** Passes dependencies explicitly (`Storage` → `Service` → `Handlers`) to keep components decoupled.
- **Graceful shutdown:** Handles `SIGINT` (Ctrl+C) and `SIGTERM`, then drains active connections using a context timeout and `sync.WaitGroup`.
- **Browser-friendly testing:** Accepts long URLs through path wildcards (`/shorten/{url...}`) and redirects via short codes (`/{code}`).
- **Structured logging:** Uses Go’s native `slog` package for clean, parseable JSON or text logs.

## ⚙️ Getting Started

### Prerequisites

- Go 1.22 or higher

### Installation and Running

1. **Clone the repository:**

   ```bash
   git clone https://github.com/YOUR_USERNAME/url-shortener.git
   cd url-shortener
   ``` 