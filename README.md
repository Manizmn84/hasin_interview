# Hasin Interview Project

## Prerequisites
- Docker & Docker Compose installed on your system.

## Getting Started
1. Clone the project.
2. Run the following command in the `docker` folder to start the application:
   ```bash
   docker compose up --build
API Documentation (Testing)
The backend is running on http://localhost:8080. You can test the endpoints using the following commands in your terminal:

1. Create a Todo
Bash
curl -X POST http://localhost:8080/v1/todo/create \
     -H "Content-Type: application/json" \
     -d '{"title": "Task 1", "dsc": "Description", "np": 1.0}'
2. List All Todos
Bash
curl -X GET http://localhost:8080/v1/todo
3. Get Todo by ID
Bash
curl -X GET http://localhost:8080/v1/todo/1
4. Update Todo
Bash
curl -X PUT http://localhost:8080/v1/todo/update/1 \
     -H "Content-Type: application/json" \
     -d '{"title": "Updated Task", "dsc": "New Desc", "np": 2.0, "status": 1}'
5. Delete Todo
Bash
curl -X DELETE http://localhost:8080/v1/todo/delete/1



## Frontend Application
The frontend application is accessible through your web browser to provide a visual interface for managing your todos.

- **URL:** [http://localhost:3000](http://localhost:3000)