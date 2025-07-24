# Astrocartics API

This is a Go-based RESTful API for accessing EVE Online static data, including regions, constellations, systems, and stargates. It features a clean, three-tier architecture (Controller, Service, DBA) and provides interactive API documentation via Swagger UI.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- [Go](https://golang.org/doc/install) (version 1.22 or later)
- A running [PostgreSQL](https://www.postgresql.org/download/) database instance
- Access to a terminal or command prompt

### Installation

1.  **Clone the repository:**
    ```sh
    git clone https://github.com/astrocartics-xyz/Astrocartics-API.git
    cd Astrocartics-API
    ```

2.  **Install Go dependencies:**
    This command will download the necessary libraries (like `chi` for routing and `swag` for documentation) and update your `go.sum` file.
    ```sh
    go mod tidy
    ```

3.  **Install the Swagger `swag` tool:**
    This command installs the command-line tool used to generate the API documentation from your code annotations.
    ```sh
    go install github.com/swaggo/swag/cmd/swag@latest
    ```
    > **Note:** If your shell cannot find the `swag` command after installation, you may need to add Go's binary directory to your system's `PATH`. You can do this by adding the following line to your shell's configuration file (e.g., `.zshrc`, `.bash_profile`):
    > `export PATH=$PATH:$(go env GOPATH)/bin`
    > Remember to restart your terminal or run `source ~/.zshrc` to apply the changes.

## Configuration

The API is configured using environment variables, which can be placed in a `.env` file in the root of the project directory.

1.  **Create a `.env` file:**
    You can copy the example below to create your own `.env` file.
    ```sh
    touch .env
    ```

2.  **Populate the `.env` file:**
    Open the `.env` file and add the following content, replacing the placeholder values with your actual database credentials.

    ```env
    # ---------------------------------
    # PostgreSQL Database Configuration
    # ---------------------------------
    # The hostname or IP address of your PostgreSQL server.
    DATABASE_URL=postgres://username:password@host:port/db_name

    # ---------------------
    # API Server Configuration
    # ---------------------
    # The port for the API server to run on.
    PORT=8080
    ```

## Usage

Follow these steps to generate the documentation and run the API server.

### 1. Generate API Documentation

Before running the server for the first time, you need to generate the Swagger documentation files.

Run the following `swag` command from the project root:
```sh
swag init -g cmd/api/main.go
```
This command parses the annotations in your Go code and creates a `docs` directory containing the `swagger.json` and other necessary files.

### 2. Run the API Server

Start the server by running the `main.go` file:
```sh
go run cmd/api/main.go
```
You should see a log message indicating that the server has started:
```
2025/07/24 16:00:00 Server starting on port 8080...
```

### 3. Accessing the API

Your API is now running and accessible.

- **API Base URL:** `http://localhost:8080/api/v1`
- **Swagger UI:** `http://localhost:8080/swagger/index.html`

You can use a tool like `curl` or Postman to interact with the API endpoints, or simply open the Swagger UI in your web browser to explore and test them interactively.
