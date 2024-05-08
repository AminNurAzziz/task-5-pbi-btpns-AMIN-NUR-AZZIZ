# Fullstack Developer Project Based Internship Program

## Description
This is a simple REST API created as part of the Full Stack Developer final project. The API is built using the Go programming language and provides functionalities for user authentication, registration, logout, user list retrieval, user data update, and user data deletion. Additionally, users can access a list of profile photos, add new photos, edit profile photos, and delete profile photos.

## Features
- API Documentation available at [API Documentation](https://documenter.getpostman.com/view/30470341/2s9YsDja6b)

## Technology Used
- GO
- PostgreSQL

## Tools Used
- Gin Gonic Framework
- Gorm
- JWT Go
- Go Validator
- Docker for PostgreSQL

## Installation

Clone the repository to your local machine:

```bash
git clone https://github.com/AminNurAzziz/task-5-pbi-btpns-AMIN-NUR-AZZIZ.git
```

Change into the project directory:

```bash
cd task-5-pbi-btpns-AMIN-NUR-AZZIZ
```

Install the necessary dependencies:

```bash
go get .
```

Create a .env file and add your database configuration:
- **APP_PORT**: Port to be used by your application (eg: `8080`).
- **DB_HOST**: Your host base data (eg: `localhost` or IP address).
- **DB_PORT**: The port used by your database (eg: `5432` for PostgeSQL).
- **DB_USER**: Your base data username.
- **DB_PASSWORD**: Your basic data password.
- **DB_NAME**: Database name to be used by your application.
- **DB_SSL_MODE**: disable
- **JWT_SECRET**: The secret key used to generate and verify JWT tokens.

Make sure to populate these values with information appropriate to your development environment. All these configurations are very important to run the application properly.

Example `.env`:

```env
APP_PORT=8080
DB_HOST=localhost
DB_PORT=5407
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=photo
DB_SSL_MODE=disable
JWT_SECRET=L5b68gO3U2y4Lf05nd8mG6xRhzf9z0UIfj+vymVokd0=
```

Run the project:

```bash
go run main.go
```
