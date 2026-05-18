# go-crud

Minimum API server with authentication, user, product, and file upload features.

## API Documentation

[visit postman](https://documenter.getpostman.com/view/6507807/2sBXqRjcMs)

---

## Setup

### 1. Clone & install dependencies

```bash
go mod tidy
```

Or install a specific package:

```bash
go get <package-url>
```

### 2. Environment setup

Copy environment file:

```bash
cp .env.example .env
```

Then adjust values inside `.env`.

### 3. Run database seeder

```bash
go run cli/seed/main.go
```

### 4. Run development server

```bash
go run main.go
```

---

## Development

### Run with hot reload

Install Air:

```bash
go install github.com/air-verse/air@latest
```

Run:

```bash
air
```

Or if installed manually:

```bash
~/go/bin/air
```

---

## Project Guides

### Add config

1. Add key/value to `.env`
2. Register config in:

```txt
./config/config.go
```

### Add seeder

1. Create seeder file in:

```txt
./seeders
```

2. Register it in:

```txt
./cli/seed/main.go
```

---

## Available Commands

```bash
# install dependencies
go mod tidy

# run seeder
go run cli/seed/main.go

# run server
go run main.go

# hot reload
air
```