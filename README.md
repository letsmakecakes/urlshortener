# URL Shortener Service

A robust URL shortening service built with Go, featuring a RESTful API, MongoDB storage, and Docker support. This service allows you to create, manage, and track shortened URLs with comprehensive statistics.

## Features

- ✨ Create shortened URLs
- 🔄 Update existing URLs
- 📊 Track URL access statistics
- 🚀 RESTful API
- 🔒 Input validation
- 📝 Structured logging
- 🐳 Docker support
- 🧪 Comprehensive tests

## Tech Stack

- Go 1.21+
- MongoDB
- Docker & Docker Compose
- Zap Logger
- Gorilla Mux
- Testify (testing)

## Prerequisites

- Go 1.21 or higher
- Docker and Docker Compose
- MongoDB (if running locally)

## Getting Started

### Running with Docker

1. Clone the repository:
```bash
git clone https://github.com/letsmakecakes/urlshortener.git
cd urlshortener
```

2. Start the services:
```bash
docker-compose up -d
```

The following services will be available:
- URL Shortener API: http://localhost:8080
- MongoDB Express UI: http://localhost:8081 (admin/password)

### Running Locally

1. Set up environment variables:
```bash
cp .env.example .env
# Edit .env with your configuration
```

2. Install dependencies:
```bash
go mod download
```

3. Run the application:
```bash
go run cmd/api/main.go
```

## API Documentation

### Create Short URL

```http
POST /shorten
Content-Type: application/json

{
    "url": "https://example.com/some/long/url"
}
```

Response:
```json
{
    "id": "1",
    "url": "https://example.com/some/long/url",
    "shortCode": "abc123",
    "createdAt": "2024-01-02T12:00:00Z",
    "updatedAt": "2024-01-02T12:00:00Z"
}
```

### Get Original URL

```http
GET /shorten/{shortCode}
```

Response:
```json
{
    "id": "1",
    "url": "https://example.com/some/long/url",
    "shortCode": "abc123",
    "createdAt": "2024-01-02T12:00:00Z",
    "updatedAt": "2024-01-02T12:00:00Z"
}
```

### Update URL

```http
PUT /shorten/{shortCode}
Content-Type: application/json

{
    "url": "https://example.com/updated/url"
}
```

### Delete URL

```http
DELETE /shorten/{shortCode}
```

### Get URL Statistics

```http
GET /shorten/{shortCode}/stats
```

Response:
```json
{
    "id": "1",
    "url": "https://example.com/some/long/url",
    "shortCode": "abc123",
    "accessCount": 42,
    "createdAt": "2024-01-02T12:00:00Z",
    "updatedAt": "2024-01-02T12:00:00Z"
}
```

## Running Tests

### Unit Tests

```bash
go test ./internal/... -v
```

### Integration Tests

```bash
# Ensure the service is running
INTEGRATION_TEST=true go test ./tests/integration -v
```

## Project Structure

```
urlshortener/
├── cmd/
│   └── api/
│       └── main.go               # Application entry point
├── internal/
│   ├── api/                      # HTTP handlers and routes
│   ├── config/                   # Configuration management
│   ├── domain/                   # Domain models and interfaces
│   ├── pkg/                      # Internal packages
│   └── service/                  # Business logic
├── pkg/                          # Public packages
├── tests/                        # Integration tests
└── docker-compose.yml           # Docker composition
```

## Development

### Adding New Endpoints

1. Define the handler in `internal/api/handlers/url_handler.go`
2. Add the route in `internal/api/routes/routes.go`
3. Implement business logic in `internal/service/url_service.go`
4. Add tests for the new functionality

### Logging

The service uses Zap for structured logging. Logs are configured based on the environment:
- Development: Console-friendly, colored output
- Production: JSON format for better parsing

### Monitoring

Monitor the application using:
```bash
# View logs
docker-compose logs -f app

# Access MongoDB Express
open http://localhost:8081
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.