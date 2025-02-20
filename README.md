# go-forge-api

[![Go Report Card](https://goreportcard.com/badge/github.com/bambutcha/go-forge-api)](https://goreportcard.com/report/github.com/bambutcha/go-forge-api) [![License](https://img.shields.io/github/license/bambutcha/go-forge-api)](https://github.com/bambutcha/go-forge-api/blob/master/LICENSE)

<br>

A powerful and flexible toolkit for building production-ready REST APIs in Go.

  

- [Features](#features)

- [Quick Start](#quick-start)

- [Contributing](#contributing)

- [License](#license)

- [Acknowledgments](#acknowledgments)

---

## Features

- 🚀 **Production-Ready**: Built with best practices and real-world usage in mind

- 🔒 **Secure**: Built-in authentication and middleware support

- 📝 **Clean Architecture**: Following SOLID principles and clean architecture patterns

- 🔌 **Extensible**: Easy to add new features and customize existing ones

- 📊 **Database Support**: PostgreSQL support with migrations included

- 🔍 **Logging**: Structured logging with logrus

- ⚡ **Performance**: Optimized for high performance

- 🧪 **Testing**: Comprehensive testing utilities included

---

## Quick Start

### Prerequisites

- Go 1.24

- PostgreSQL

- Make (optional, for using Makefile commands)

### Installation

```bash

# Clone the repository

git clone https://github.com/bambutcha/go-forge-api.git

# Navigate to the project

cd go-forge-api

# Install dependencies

go mod download

```

### Running

```bash

# Using make

make run

# Or directly

go run cmd/api-server/main.go

```

---

## Project Structure

```
go-forge-api/
├── cmd/                   # Application entry points
├── configs/               # Configuration files
├── internal/              # Private application code
│   ├── app/               # Application core
│   │   ├── api-server/    # HTTP server implementation
│   │   ├── model/         # Domain models
│   │   └── store/         # Data access layer
└── migrations/            # Database migrations
```

---

## API Endpoints

| Method | Path              | Description           |
| ------ | ----------------- | --------------------- |
| POST   | `/users`          | Create a new user     |
| POST   | `/sessions`       | Create a new session  |
| GET    | `/private/whoami` | Get current user info |

---

## Development

### Running Tests

`# Run all tests make test # Run specific tests go test ./internal/app/...`

### Database Migrations

`# Create a new migration make migrate-create name=add_users_table  # Apply migrations make migrate-up  # Rollback migrations make migrate-down`

---

## Contributing

We welcome contributions! Please see our [Contributing Guidelines](https://github.com/BinaryBenefactors/obscura-project/blob/master/docs/CONTRIBUTING.md) for details.

### Development Process

2. Fork the repository
3. Create your feature branch (`git checkout -b feature/amazing-feature`)
4. Commit your changes (`git commit -m 'Add some amazing feature'`)
5. Push to the branch (`git push origin feature/amazing-feature`)
6. Open a Pull Request

---

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

## Acknowledgments

- [Gorilla Mux](https://github.com/gorilla/mux) for routing
- [Logrus](https://github.com/sirupsen/logrus) for logging
- [PostgreSQL](https://www.postgresql.org/) for database
