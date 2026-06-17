# Hasin Interview Project

## Prerequisites

- Docker & Docker Compose installed on your system.

## Getting Started

1. Clone the project.
2. Before starting, install frontend dependencies:
   ```bash
   cd ../hasin_interview_frontend && npm install && cd -
   ```
3. Run the following command in the `docker` folder to start the application:
   ```bash
   cd docker && docker compose --env-file ../.env up --build
   ```

## API Documentation

The backend is running on `http://localhost:8080`. You can test the endpoints using the following commands:

### 1. Create a Todo

```bash
curl -X POST http://localhost:8080/v1/todo/create \
     -H "Content-Type: application/json" \
     -d '{"title": "Task 1", "dsc": "Description", "np": 1.0}'
```

### 2. List All Todos

```bash
curl -X GET http://localhost:8080/v1/todo
```

### 3. Get Todo by ID

```bash
curl -X GET http://localhost:8080/v1/todo/1
```

### 4. Update Todo

```bash
curl -X PUT http://localhost:8080/v1/todo/update/1 \
     -H "Content-Type: application/json" \
     -d '{"title": "Updated Task", "dsc": "New Desc", "np": 2.0, "status": 1}'
```

### 5. Delete Todo

```bash
curl -X DELETE http://localhost:8080/v1/todo/delete/1
```

## Frontend Application

The frontend application is accessible through your web browser to provide a visual interface for managing your todos.

- **URL:** [http://localhost:3000](http://localhost:3000)
