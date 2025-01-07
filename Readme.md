# Vehicle Registration Manager

[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=ArianeASA_vehicle-registration-manager&metric=coverage)](https://sonarcloud.io/summary/overall?id=ArianeASA_vehicle-registration-manager)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=ArianeASA_vehicle-registration-manager&metric=alert_status)](https://sonarcloud.io/summary/overall?id=ArianeASA_vehicle-registration-manager)
[![SonarQube Pipeline](https://github.com/ArianeASA/vehicle-registration-manager/actions/workflows/sonar.yml/badge.svg)](https://github.com/ArianeASA/vehicle-registration-manager/actions/workflows/sonar.yml)
[![Pre-deployment Pipeline](https://github.com/ArianeASA/vehicle-registration-manager/actions/workflows/build.yml/badge.svg)](https://github.com/ArianeASA/vehicle-registration-manager/actions/workflows/build.yml)
[![Migration apply](https://github.com/ArianeASA/vehicle-registration-manager/actions/workflows/migration-apply.yml/badge.svg?branch=develop)](https://github.com/ArianeASA/vehicle-registration-manager/actions/workflows/migration-apply.yml)


## Description

The Vehicle Registration Manager is a sample server application built with Go for managing vehicle data. It provides RESTful API endpoints for listing, searching, registering, and updating vehicles. The project includes database migration support using goose, and API documentation generated with Swagger. The application can be run locally or with Docker Compose for easy setup and deployment.


## Implementation

The project is implemented using the following technologies:
- **Go**: The main programming language used for the server application.
- **Goose**: A CLI tool for managing database migrations.
- **Swagger**: For generating API documentation.
- **Docker Compose**: For containerizing the application and its dependencies.
- **GitHub Actions**: For automated deployment and CI/CD pipelines.


## Getting Started

### Prerequisites

- **Go 1.23 or higher**: Ensure you have Go installed. You can download it from the [official Go website](https://golang.org/dl/).
- **Go Modules**: Make sure to have `go.mod` and `go.sum` files in your project directory. These files manage the dependencies for your Go project.
- **Goose CLI**: Install the `goose` CLI tool for managing database migrations. You follow the installation instructions [here](https://github.com/pressly/goose?tab=readme-ov-file#install).
- **Docker**: Ensure Docker is installed and running on your machine. Docker is used to containerize the application and its dependencies.
- **Docker Compose**: Install Docker Compose to manage multi-container Docker applications. You can follow the installation instructions [here](https://docs.docker.com/compose/install/).


### Installing

1. Clone the repository:
    ```sh
    git clone https://github.com/ArianeASA/vehicle-registration-manager.git
    cd vehicle-registration-manager
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

### Running the Application Locally

1. **Start the database with Docker Compose**:
    ```sh
    make docker-up-database
    ```
   The database is started in a Docker container using the `docker-compose.yml` file.
   Use database type postgres to run the database in a Docker container.
   
   Configurations:
      - POSTGRES_USER=yourusername
      - POSTGRES_PASSWORD=yourpassword
      - POSTGRES_DB=dealership
      - Schema: prod
   
   The database will start on http://localhost:5432.


2. **Check database migrations status**:
    ```sh
    make goose-status
    ```

3. **Run database migrations**:
    ```sh
    make goose-up
    ```
   All database migrations are stored in the `migrations` directory.


4. **Generate API documentation**:
    ```sh
    make doc
    ```
   The API documentation is available in the docs/swagger.json file. You can use Swagger UI to visualize and interact with the API.  
   The Swagger UI is available at http://localhost:8080/swagger/index.html.

   ![img.png](img.png)

5. **Run the application**:
    ```sh
    make run
    ```
   The server will start on http://localhost:8080.


6. **Rollback database migrations**:
    ```sh
    make goose-down
    ```
    
7. **Stop the application with Docker Compose**:
    ```sh
    make docker-down-all
    ```
   

8. **Clean up generated files**:
    ```sh
    make clean
    ```
   

9. **Show help message**:
    ```sh
    make help
    ```


### Testing

To run the tests, use the following command:
```sh
go test ./...
```


### License
This project is licensed under the Apache 2.0 License - see the LICENSE file for details.