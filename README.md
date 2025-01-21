# Event Management System

This is a learning-focused toy project that implements an event management system using Go and the Gin web framework.

## Features

- User signup and login with JWT-based authentication and authorization
- CRUD operations for events
- Event registration and cancellation
- PostgreSQL database integration with migrations using `golang-migrate`

## Prerequisites

- Go (version 1.19 or later)
- Docker
- golang-migrate

## Getting started

### 1. Clone the repo

```bash
git clone https://github.com/dsabljic/event-management
cd event-management
```

### 2. Set up environment variables

Create a `.env` file in the project root with the following content:

```env
ADDR=:8080
DB_ADDR=postgres://admin:admin@localhost:54332/event-management?sslmode=disable
```

### 3. Start the PostgreSQL DB

Use Docker Compose to start the PostgreSQL service:

```bash
docker-compose up --build
```

### 4. Apply database migrations

Run migrations to set up the database schema:

```bash
make migrate-up
```

### 5. Run the application

Start the server with:

```bash
go run main.go
```

or optionally:

```bash
air
```

The API will be available at `http://localhost:8080`.

## API endpoints

### Public endpoints

- `POST /signup` - Create a new user
- `POST /login` - Login with valid credentials

### Authenticated endpoints

- `POST /events` - Create a new event
- `GET /events` - Retrieve all events
- `GET /events/:id` - Get details of a specific event
- `PUT /events/:id` - Update an event
- `DELETE /events/:id` - Delete an event
- `POST /events/:id/register` - Register for an event
- `DELETE /events/:id/register` - Cancel registration
