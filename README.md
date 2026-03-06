# beta-workplace

A production-ready Go backend project scaffolded with [go-backend](https://github.com/namf2001/go-backend-template).

## Features

- 🏗️ Clean Architecture (Handler → Controller → Repository)
- 🔧 Chi HTTP Router with middleware
- 🔐 JWT Authentication
- 🌐 Google OAuth2
- 🗄️ Postgres Database with connection pooling
- 📦 SQL Migrations
- 📝 Swagger API Documentation
- 📊 Prometheus Metrics
- 🐳 Docker & Docker Compose
- 🚀 GitHub Actions CI/CD
- ♻️ Graceful Shutdown
- 🔥 Air Hot-reload for development

## Getting Started

### Prerequisites

- Go 1.24+
- Postgres
- Docker & Docker Compose

### Setup

```bash
# Copy environment file
cp .env.example .env.dev

# Start database
docker-compose up -d


# Run migrations
make migrate-up

# Run development server
make dev
```

## Project Structure

```
.
├── cmd/server/          # Application entry point
├── config/              # Configuration management
├── internal/
│   ├── controller/      # Business logic
│   ├── handler/         # HTTP handlers & middleware
│   ├── model/           # Domain models
│   ├── pkg/             # Internal shared packages
│   └── repository/      # Data access layer
├── migrations/          # SQL migration files
├── Dockerfile
├── docker-compose.yml
├── Makefile
└── .env.example
```

## Available Commands

```bash
make help              # Show all available commands
make dev               # Run with Air hot-reload
make build             # Build binary
make test              # Run tests
make swagger           # Generate Swagger docs
make docker-build      # Build Docker image
make docker-up         # Start containers
make docker-down       # Stop containers
```

## License

MIT
