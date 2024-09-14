## Blogging MD

This is a simple blogging application written in Go.

### Features

- User Authentication
- Blog Posts
- Comments
- Pagination
- Search

### How to run

- Requirements
    - Docker
    - Docker Compose

- Run
  ```bash
  docker-compose up -d
  ```
  
### Database Schema
```mermaid
erDiagram
    USERS {
        int id PK
        string name
        string email
        string password_hash
        timestamp created_at
        timestamp updated_at
    }

    BLOG_POSTS {
        int id PK
        string title
        text content
        int author_id FK
        timestamp created_at
        timestamp updated_at
    }

    COMMENTS {
        int id PK
        int post_id FK
        string author_name
        text content
        timestamp created_at
    }

    USERS ||--o{ BLOG_POSTS : "has"
    BLOG_POSTS ||--o{ COMMENTS : "has"

```

### How to use

- Register
    - Endpoint: `/register`
    - Method: `POST`
    - Body
      ```json
      {
        "name": "John Doe",
        "email": "john@example.com",
        "password": "password"
      }
      ```
        - Response
          ```json
          {
            "code": 200,
            "status": "Successfully Registered",
            "data": {
              "id": "1",
              "name": "John Doe",
              "email": "john@example.com"
          }
          ```
- Login
- Endpoint: `/login`
- Method: `POST`
- Body
  ```json
  {
    "email": "john@example.com",
    "password": "password"
  }
  ```
    - Response
      ```json
      {
        "code": 200,
        "status": "Successfully Logged In",
        "data": {
          "token": "Token"
        }
      }
    ```
