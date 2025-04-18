# Simple Go Web Server using Gin

This is a simple web server built with Go and the Gin framework.

## Functionality

The `server.go` file sets up a basic HTTP server using Gin:

*   It initializes a default Gin router.
*   It defines a single route handler for `GET /`.
*   When a GET request is made to the root path (`/`), the server responds with the plain text string "Hello, World!" and an HTTP status code of 200 (OK).
*   The server listens and serves on `0.0.0.0:8080` by default.

## Development Setup

**Prerequisites:**

*   **Go:** Ensure you have Go installed on your system. You can download it from [https://go.dev/dl/](https://go.dev/dl/).

**Steps:**

1.  **Clone the repository:**
    ```bash
    git clone <repository-url>
    cd <repository-directory>
    ```
2.  **Install dependencies:** The necessary dependencies (Gin and Testify) are listed in the `go.mod` file. You can download them using:
    ```bash
    go mod download
    ```
    Alternatively, they will be automatically downloaded when you build or run the project.

## Build

To build the executable binary:

```bash
go build
```

This will create an executable file (e.g., `mymodule` or `server` depending on your module name and OS) in the project directory.

## Run

To run the server directly without building:

```bash
go run server.go
```

The server will start and listen on port 8080. You can access it by navigating to `http://localhost:8080` in your web browser or using a tool like `curl`:

```bash
curl http://localhost:8080
```

You should see the output: `Hello, World!`

## Test

To run the unit tests defined in `server_test.go`:

```bash
go test
```

The output will indicate whether the tests passed or failed.
