# Cloud Assignment REST API

This project is a REST API for managing tasks. It provides endpoints to create, read, update, and delete tasks, as well as a health check endpoint.

## Project Structure

```bash
cloud_assignment
├── cmd
│   └── server
│       └── main.go          # Entry point of the application
├── internal
│   ├── handlers
│   │   └── handlers.go      # Handler functions for API endpoints
│   ├── middleware
│   │   └── middleware.go     # Middleware functions for the API
│   └── models
│       └── models.go        # Data structures used in the application
├── pkg
│   └── utils
│       └── utils.go         # Utility functions for the application
├── configs
│   └── config.go            # Configuration settings for the application
├── go.mod                    # Module dependencies and Go version
├── go.sum                    # Checksums for module dependencies
└── README.md                 # Project documentation
```

## API Endpoints

- `POST /tasks` - Create a new task
- `GET /tasks` - List all tasks
- `GET /tasks/{id}` - Get a specific task by ID
- `PUT /tasks/{id}` - Update a specific task by ID
- `DELETE /tasks/{id}` - Delete a specific task by ID
- `GET /health` - Health check endpoint

## Setup Instructions

1. Clone the repository:

   ```bash
   git clone <repository-url>
   cd cloud_assignment
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Run the server:

   ```bash
   go run cmd/server/main.go
   ```

4. Access the API at `http://localhost:8080`.
