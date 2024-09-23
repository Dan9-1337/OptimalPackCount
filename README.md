# OptimalPackCount

## Overview

**OptimalPackCount** is a Golang-based application designed to calculate the optimal number of packs required to fulfill specific orders efficiently. The project is structured with clear separation of concerns, ensuring maintainability and scalability. This application can be used in various domains where packing optimization is crucial, such as logistics, e-commerce, and manufacturing.

## Features

- **Service-Oriented Architecture**: The project is divided into services, handlers, models, and UI components.
- **RESTful API**: Provides an interface to interact with the core logic of the application.
- **Docker Support**: The application is Dockerized for easy deployment and scalability.
- **Unit Testing**: Comprehensive unit tests ensure the robustness of the application.

## Getting Started

### Prerequisites

Before you begin, ensure you have the following installed:

- **Golang**: [Install Go](https://golang.org/doc/install) (version 1.16 or higher)
- **Docker**: [Install Docker](https://www.docker.com/get-started)

### Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/Dan9-1337/OptimalPackCount.git
    cd OptimalPackCount
    ```

2. Build the project:
    ```bash
    go build -o optimalpackcount
    ```

3. Run the application:
    ```bash
    ./optimalpackcount
    ```

4. Alternatively, run the application using Docker:
    ```bash
    docker build -t optimalpackcount .
    docker run -p 8080:8080 optimalpackcount
    ```

## API Documentation

### Calculate Packs

This endpoint calculates the optimal number of packs required to fulfill a given order based on the available package sizes.

- **Endpoint**: `/calculate-packs`
- **Method**: `POST`

#### Request Body
```json
{
  "orderedItems": 1000,
  "packageSizeList": [5000, 2000, 1000, 500, 250]
}
```

### Project Structure

`handler/`: Contains the HTTP handlers that manage incoming API requests and outgoing responses.

`models/`: Defines the data structures used throughout the application.

`service/`: Implements the core business logic for calculating the optimal pack count.

`tests/`: Contains unit tests to verify the functionality and reliability of the application.

`ui/`: Holds static assets for the user interface, if applicable.

### Running Tests
To run the unit tests provided with the project:

```bash
go test ./...
```

### Docker Deployment
If you prefer to deploy the application using Docker:

1. Build the Docker image:
```bash
docker build -t optimalpackcount .
```
2. Run the Docker container:
```bash
docker run -p 8080:8080 optimalpackcount
```
