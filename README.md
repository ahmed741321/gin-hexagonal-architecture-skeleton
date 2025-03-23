# gin-hexagonal-architecture-skeleton

This project implements a hexagonal architecture in Go, providing a clean separation between the application core and external dependencies.

## Project Structure

```
gin-hexagonal-architecture-skeleton
├── cmd
│   └── server
│       └── main.go          # Entry point of the application
├── internal
│   ├── app
│   │   ├── app.go           # Application logic
│   │   └── app_test.go      # Unit tests for application logic
│   ├── domain
│   │   └── models.go        # Domain models
│   ├── ports
│   │   ├── http
│   │   │   └── handler.go    # HTTP handlers
│   │   └── repository
│   │       └── repository.go  # Repository interfaces
│   └── services
│       └── service.go       # Service layer logic
├── pkg
│   └── config
│       └── config.go        # Application configuration
├── go.mod                    # Module definition
└── README.md                 # Project documentation
```

## Getting Started

To get started with the project, follow these steps:

1. **Clone the repository:**
   ```
   git clone https://github.com/yourusername/gin-hexagonal-architecture-skeleton.git
   cd gin-hexagonal-architecture-skeleton
   ```

2. **Install dependencies:**
   ```
   go mod tidy
   ```

3. **Run the application:**
   ```
   go run cmd/server/main.go
   ```

4. **Run tests:**
   ```
   go test ./...
   ```

## Usage

Once the application is running, you can interact with it through the defined HTTP endpoints. Refer to the documentation in `internal/ports/http/handler.go` for available routes and their usage.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any enhancements or bug fixes.

## License

This project is licensed under the MIT License. See the LICENSE file for details.# gin-hexagonal-architecture-skeleton
