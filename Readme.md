# GoLang Quotes API Application

## Introduction

This application provides a simple RESTful API built with GoLang and the Echo framework. It offers two endpoints: one to fetch a list of quotes and another to get the quote of the day. The application is designed to be lightweight, efficient, and easy to use.

## Technology Stack

- **GoLang**: The backend of this application is developed using GoLang, a powerful programming language known for its performance and simplicity.
- **Echo Framework**: Echo is a fast and minimalist Go web framework that provides robust routing and middleware features for building web applications.
- **MongoDB**: MongoDB is used as the database to store and manage the quotes data. It offers flexibility and scalability, making it suitable for handling various types of data.
- **Docker**: The application is containerized using Docker, allowing for easy deployment and scaling across different environments.

## Getting Started

### Using Docker

1. Ensure you have Docker installed on your system. If not, you can download and install it from [here](https://docs.docker.com/get-docker/).
2. Clone this repository:
   ```sh
   git clone https://github.com/YashTripathi01/quotes-API-Go
   ```
3. Navigate to the project directory:
   ```sh
   cd quotes-API-Go
   ```
4. Build the docker image:
   ```sh
   docker build -t quotes-image .
   ```
5. Run the docker container:
   ```sh
   docker run -p 1323:1323 -e DATABASE_URL=mongodb://localhost:27017/quotes quotes-image
   ```

### Without Docker

1. Ensure you have Go installed on your system. If not, you can download and install it from [here](https://golang.org/dl/).
2. Clone this repository:
   ```sh
   git clone https://github.com/YashTripathi01/quotes-API-Go
   ```
3. Navigate to the project directory:
   ```sh
   cd quotes-API-Go
   ```
4. Copy the `.env.example` file to `.env` and update the values:
   ```sh
   cp .env.example .env
   ```
5. Install dependencies:
   ```sh
   go mod tidy
   ```
6. Build the application:
   ```sh
   go build
   ```
7. Run the application:
   ```sh
   ./quotes-API-Go
   ```

## Endpoints

### Get all quotes list

- URL: `/quotes`
- Method: GET
- Description: Retrieves a list of all quotes available in the system.

### Get the quote of the day

- URL: `/quotes/today`
- Method: GET
- Description: Retrieves the quote of the day.

## Usage

- Once the application is running, you can use any HTTP client like cURL, Postman, or browser to access the defined endpoints.

## Dependencies

[Echo](https://github.com/labstack/echo): Web framework for Go.

## Contributing

Contributions are welcome! If you find any issues or have suggestions, feel free to open an issue or create a pull request.
