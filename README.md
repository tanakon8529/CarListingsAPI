## CarListingsAPI

This repository contains the backend API for a used car listing application, Daveslist. The API is built using Golang, and utilizes GIN framework for the HTTP server, GORM for ORM, and supports Swagger for API documentation. Docker is used for containerization and Redis for caching.

### Table of Contents

- [Project Structure](#project-structure)
- [Tech Stack](#tech-stack)
- [Installation](#installation)
- [Running the Project](#running-the-project)
- [Running Tests](#running-tests)
- [API Documentation](#api-documentation)
- [License](#license)

### Project Structure

```plaintext
daveslist-emdpcv/
├── api/
│   ├── controllers/             # Contains all the controller files for handling HTTP requests
│   │   ├── auth.go
│   │   ├── category_controller.go
│   │   ├── health.go
│   │   ├── listing_controller.go
│   │   ├── private_message_controller.go
│   │   ├── reply_controller.go
│   │   ├── user_controller.go
│   ├── database/                # Database connection and initialization
│   │   └── database.go
│   ├── docs/                    # Swagger documentation files
│   │   ├── docs.go
│   │   ├── swagger.json
│   │   ├── swagger.yaml
│   ├── files/                   # Sample files for testing
│   │   ├── go1.jpg
│   │   ├── go2.jpeg
│   ├── middlewares/             # Middleware files for request handling
│   │   ├── error_logger.go
│   │   ├── token_auth.go
│   ├── models/                  # Data models
│   │   └── models.go
│   ├── routes/                  # Route definitions
│   │   └── routes.go
│   ├── services/                # Service layer for business logic
│   │   ├── gen_token.go
│   │   ├── redis.go
│   ├── settings/                # Configuration settings
│   │   └── configs.go
│   ├── tests/                   # Test files
│   │   └── api_test.go
│   └── main.go                  # Entry point of the application
├── .env                         # Environment variables
├── docker-compose.yml           # Docker Compose configuration
├── Dockerfile                   # Dockerfile for containerization
├── go.mod                       # Go module file
├── go.sum                       # Go module dependencies
├── README.md                    # Project documentation
├── run_tests.sh                 # Shell script to run tests
└── .git                         # Git version control directory
```

### Tech Stack

- **Golang**: The programming language used for the API.
- **GIN**: A web framework written in Go.
- **GORM**: ORM library for Golang.
- **SQLite**: Lightweight database used for local development.
- **Redis**: In-memory data structure store, used as a database, cache, and message broker.
- **Docker**: Containerization platform.
- **Swagger**: API documentation.

### Installation

1. **Clone the repository**

    ```sh
    git clone https://github.com/tanakon8529/CarListingsAPI.git
    ```

2. **Set up environment variables**

    Create a `.env` file in the root directory and add the required environment variables. Refer to the `.env.example` for guidance.

3. **Install dependencies**

    ```sh
    go mod download
    ```

4. **Build the Docker image**

    ```sh
    docker-compose build
    ```

### Running the Project

1. **Start the Docker containers**

    ```sh
    docker-compose up
    ```

2. **Run the application**

    The application will be available at `http://localhost:8080`.

3. **Access API documentation**

    Swagger documentation is available at `http://localhost:8080/swagger/index.html`.

### Running Tests

1. **Ensure the application is running** (if not already running via Docker)

    ```sh
    docker-compose up -d
    ```

2. **Run the tests**

    ```sh
    ./run_tests.sh
    ```

### API Documentation

The API is documented using Swagger. Once the application is running, you can access the API documentation at:

```
http://localhost:8080/swagger/index.html
```

### License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
