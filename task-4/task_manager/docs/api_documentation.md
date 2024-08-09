# Task Management API Documentation

Welcome to the documentation for the Task Management API! This API allows you to manage tasks in a simple and efficient way. The full postman documentation can be found [here](https://documenter.getpostman.com/view/27059858/2sA3s1pY2B). 

## Endpoints

### GET /tasks

Retrieve a list of all tasks.

### POST /tasks

Create a new task.

### GET /tasks/{id}

Retrieve a specific task by its ID.

### PUT /tasks

Update a specific task.

### DELETE /tasks/{id}

Delete a specific task by its ID.

## Request and Response Examples

### GET /tasks

**Request:**

```
GET /tasks
```

**Response:**

```
Status: 200 OK
Content-Type: application/json

[
    {
        "id": 1,
        "title": "Task 1",
        "description": "This is task 1",
        "deadline": "2020-01-01",
        "status": 0
    },
    {
        "id": 2,
        "title": "Task 2",
        "description": "This is task 2",
        "deadline": "2020-01-01",
        "status": 1
    }
]
```

### POST /tasks

**Request:**

```
POST /tasks
Content-Type: application/json

{
    "title": "New Task",
    "description": "This is a new task",
    "deadline": "2020-01-01",
    "status": 1
}
```

**Response:**

```
Status: 201 Created
Content-Type: application/json

{
    "id": 1,
    "title": "New Task",
    "description": "This is a new task",
    "deadline": "2020-01-01",
    "status": 1
}
```

### GET /tasks/{id}

**Request:**

```
GET /tasks/1
```

**Response:**

```
Status: 200 OK
Content-Type: application/json

{
    "id": 1,
    "title": "Task 1",
    "description": "This is task 1",
    "deadline": "2020-01-01",
    "status": 1
}
```

### PUT /tasks/{id}

**Request:**

```
PUT /tasks
Content-Type: application/json

{
    "id": 1,
    "title": "Updated Task",
    "description": "This task has been updated",
    "deadline": "2022-01-02"
    "status": 0
}
```

**Response:**

```
Status: 200 OK
Content-Type: application/json

{
    "id": 1,
    "title": "Updated Task",
    "description": "This task has been updated",
    "deadline": "2022-01-02"
    "status": 0
}
```

### DELETE /tasks/{id}

**Request:**

```
DELETE /tasks/1
```

**Response:**

```
Status: 204 No Content
```
