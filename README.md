# mytodo-app

A simple to-do list API built with Go and the Fiber web framework.

## Features

*   User registration and login
*   JWT-based authentication
*   CRUD operations for to-dos (Create, Read, Update, Delete)

## Technologies Used

*   [Go](https://golang.org/)
*   [Fiber](https://gofiber.io/)
*   [MySQL](https://www.mysql.com/)
*   [Docker](https://www.docker.com/)

## Getting Started

### Prerequisites

*   [Go](https://golang.org/doc/install)
*   [Docker](https://docs.docker.com/get-docker/)
*   [Docker Compose](https://docs.docker.com/compose/install/)

### Installation

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/rafli024/mytodo-app.git
    cd mytodo-app
    ```

2.  **Create a `.env` file:**
    Duplicate the `.env.example` file and rename it to `.env`. Then, fill in the required environment variables for your MySQL database connection.

3.  **Start the database:**
    Run the following command to start the MySQL database container:
    ```bash
    docker-compose up -d
    ```

4.  **Run the application:**
    ```bash
    go run main.go
    ```
    The application will be running on `http://localhost:8080` (or the port specified in your `.env` file).

### Build for Production

To build the application for production, run the following command. This will create a binary named `mytodo-app` that is ready for deployment.

```bash
GOOS=linux GOARCH=amd64 go build -o mytodo-app main.go
```

## API Endpoints

All endpoints are prefixed with `/v1`.

### Authentication

*   `POST /auth/register`: Register a new user.
*   `POST /auth/login`: Log in and receive a JWT token.

### Todos (Requires Authentication)

*   `GET /todos`: Get all to-dos for the authenticated user.
*   `POST /todos`: Create a new to-do.
*   `PUT /todos/:id`: Update a to-do by its ID.
*   `DELETE /todos/:id`: Delete a to-do by its ID.

## Project Structure

```
.
├── internal/         # Internal application logic
│   ├── config/       # Configuration management
│   ├── constant/     # Application constants
│   ├── contract/     # Go interfaces
│   ├── datasources/  # Database connections
│   ├── entities/     # Database models
│   ├── handler/      # HTTP handlers
│   ├── middlewares/  # Fiber middlewares
│   ├── model/        # Request/response models
│   ├── router/       # API router
│   └── service/      # Business logic
├── migrations/       # Database migration files
├── pkg/              # Reusable packages
├── public/           # Static files
└── test/             # Test files
```
