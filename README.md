# Sword Test

This project is a Golang application built using the Echo framework to serve a RESTful API. It utilizes Swagger for API documentation, NATS as a message broker for event-based communication, MySQL as the database, and Docker for containerization. Additionally, unit tests have been implemented for the repository and usecase layers.

## Rules

- This application has two types of users (Manager, Technician).
- The technician performs tasks and is only able to see, create or update his own performed tasks.
- The manager can see tasks from all the technicians, delete them, and should be notified when some tech performs a task.
- A task has a summary (max: 2500 characters) and a date when it was performed, the summary from the task can contain personal information.

## Prerequisites

Make sure you have the following installed:

- Docker
- Docker compose

## Installation

- Clone the repository:

   ```shell
   git clone https://github.com/netorissi/SwordTest.git
   ```

## API Usage

1. To start the application locally without Docker, run the following command:

  ```shell
   make run
   ```
   OR
   ```shell
   docker-compose up -d
   ```

   The server will be up and running on port 8080.

2. Once the server has started, you will see the following message in the terminal:

   ```
   Server started on :8080.
   ```

   This indicates that the API is ready to accept requests.

- Access the API documentation generated by Swagger by visiting [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html) in your browser.

- Authorize the user to access the APIs. The token should be in the format `Bearer example_token_tech_1`.
Use the following tokens for testing:
1. `Bearer example_token_tech_1` - User tech 1
2. `Bearer example_token_tech_2` - User tech 2
3. `Bearer example_token_manager` - User manager

![Swagger Docs](./assets/swagger.png)

## Using cURLs
> Only manager can list all tasks

```
curl -X 'GET' \
  'http://localhost:8080/api/v1/tasks' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer example_token_manager'
```

> Only tech can create task

```
curl -X 'POST' \
  'http://localhost:8080/api/v1/tasks' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer example_token_tech_1' \
  -H 'Content-Type: application/json' \
  -d '{
  "summary": "string"
}'
```

> Only tech can list my tasks

```
curl -X 'GET' \
  'http://localhost:8080/api/v1/tasks/me' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer example_token_tech_1'
```

> Only owner can update task

```
curl -X 'PUT' \
  'http://localhost:8080/api/v1/tasks/:id' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer example_token_tech_1' \
  -H 'Content-Type: application/json' \
  -d '{
  "summary": "string"
}'
```

> Only manager can delete task

```
curl -X 'DELETE' \
  'http://localhost:8080/api/v1/tasks/:id' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer example_token_manager'
```

> Only owner can complete task

```
curl -X 'PATCH' \
  'http://localhost:8080/api/v1/tasks/:id/complete' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer example_token_tech_1'
```

## Mocks

- Three test users were generated in the database for testing purposes.
1. User with ID 1 is a manager.
2. User with ID 2 is a tech user.
3. User with ID 3 is also a tech user.
