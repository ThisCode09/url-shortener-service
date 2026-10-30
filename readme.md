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
```text

## 🏗️ Layer Breakdown

The project strictly separates concerns into distinct layers, ensuring high testability, maintainability, and zero reliance on heavy third-party frameworks.

1. **HTTP Handlers (`internal/handlers`)**: 
   - Uses Go 1.22+ enhanced `net/http` path wildcards (`{url...}`, `{code}`) for seamless browser-based testing without external tool requirements.
   - Decoupled from storage; communicates exclusively via the service layer.
2. **Business Logic (`internal/shortener`)**: 
   - Core service containing rules for shortening URLs, enforcing uniqueness, and mapping resolution.
3. **Storage (`internal/storage`)**: 
   - Implements thread-safe JSON persistence, automatically serializing and saving state locally.

---

## 🚀 Key Features

* **Zero Framework Bloat**: Built entirely using Go's native `net/http` standard library and custom multiplexer routing.
* **Constructor Injection**: Strict dependency passing (`Storage` $\rightarrow$ `Service` $\rightarrow$ `Handlers`) ensuring clean decoupling.
* **Graceful Shutdown**: Intercepts `SIGINT` (Ctrl+C) and `SIGTERM` signals, executing an orderly context-timeout drain of active connections using `sync.WaitGroup`.
* **Browser-Friendly Testing**: Designed to accept long URLs directly via path wildcards (`/shorten/{url...}`) and redirect instantly via short codes (`/{code}`).
* **Structured Logging**: Utilizes Go's native `slog` package for clean, parseable JSON/text logging.

---

## ⚙️ Getting Started

### Prerequisites
* Go 1.22 or higher installed on your machine.

### Installation & Running

1. **Clone the repository:**
   ```bash
   git clone [https://github.com/YOUR_USERNAME/url-shortener.git](https://github.com/YOUR_USERNAME/url-shortener.git)
   cd url-shortener