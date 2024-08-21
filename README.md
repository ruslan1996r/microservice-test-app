# How to start
- `docker-compose up --build` - use this command in the root of the application to start all services at once

# Testing

### Gateway API
#### `POST http://localhost:8002/api/v1/users/sign_in`
- **Description**: Returns the token. Use this token to get a list of Users

#### `GET http://localhost:8002/api/v1/users`
- **Description**: Get all users.
- **Headers**: `Authorization: <your_token>` - token header is required.

#### `GET http://localhost:8002/api/v1/books`
- **Description**: Get all books.

### Auth API
#### `POST http://localhost:8001/api/v1/token`
- **Description**: Returns the token.

#### `POST http://localhost:8001/api/v1/check_token`
- **Description**: Get all books. Validates the token. Returns the `true` string on success.
- **Request**:
  ```json
  {
    "token": "string"
  }

### Resource API
#### `GET http://localhost:8000/api/v1/users`
- **Description**: Get all users.

#### `GET http://localhost:8000/api/v1/books`
- **Description**: Get all books.