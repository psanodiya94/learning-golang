# learning-golang
## Overview

This repository contains code and resources for learning the Go programming language. It includes various examples and exercises to help you get started with Go and improve your skills.

## Prerequisites

Before you begin, ensure you have the following installed:

- Go (latest version)
- A code editor (such as VSCode)

## Getting Started

To run the main Go program, use the following command:

```sh
go run main.go
```

## Running Tests

To run the tests for the project, you can use the following commands:

- Run all tests:
    ```sh
    go test
    ```

- Run tests with detailed output:
    ```sh
    go test -v
    ```

- Run tests and display coverage:
    ```sh
    go test -cover
    ```

- Run tests, generate a coverage profile, and display the coverage report in a web browser:
    ```sh
    go test -coverprofile=coverage.out && go tool cover -html=coverage.out
    ```
