[![SonarQube Pipeline](https://github.com/ArianeASA/vehicle-registration-manager/actions/workflows/sonar.yml/badge.svg)](https://github.com/ArianeASA/vehicle-registration-manager/actions/workflows/sonar.yml)
[![Pre-deployment Pipeline](https://github.com/ArianeASA/vehicle-registration-manager/actions/workflows/build.yml/badge.svg)](https://github.com/ArianeASA/vehicle-registration-manager/actions/workflows/build.yml)
# Vehicle API

This project is a sample server for managing vehicles, built with Go.

## Description

The Vehicle API allows you to manage vehicle information, including listing vehicles, adding new vehicles, updating existing vehicles, and deleting vehicles.

## Getting Started

### Prerequisites

- Go 1.23 or higher
- Make sure to have `go.mod` and `go.sum` files in your project directory

### Installing

1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/vehicle-api.git
    cd vehicle-api
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

### Running the Application

To run the application, use the following command:
```sh
go run main.go
```


The server will start on http://localhost:8080.  

### API Documentation
The API documentation is available in the docs/swagger.json file. You can use Swagger UI to visualize and interact with the API.  

```shell
swag init -g cmd/vehicle-registration-manager/main.go
```
link: http://localhost:8080/swagger/index.html

### Contributing
Fork the repository
- Create a new branch (git checkout -b feature/your-feature)
- Commit your changes (git commit -m 'Add some feature')
- Push to the branch (git push origin feature/your-feature)
- Open a Pull Request

### License
This project is licensed under the Apache 2.0 License - see the LICENSE file for details.