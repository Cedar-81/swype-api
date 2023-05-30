# Swype API (mock social media api)

### Built With Go Lang Gin API with GORM ORM

This is a Go lang API project built with Gin framework and using GORM as the ORM.

## Requirements

To run this API locally, you need to set the following environment variables in a `.env` file:

- `PORT`: Port number for the localhost.
- `DB_URL`: Database URL for connecting to the database.
- `SECRET`: Secret key for password hashing.
- `SUPER_EMAIL`: Superadmin email address.
- `SUPER_PASS`: Superadmin password.

## Installation

1. Clone the repository:

```bash
git clone <repository-url>
```

2. Install the dependencies:

```bash
go mod download
```

3. Set up the environment variables by creating a `.env` file and providing the required values.

4. Run the API:

```bash
go run main.go
```

The API will be accessible at `http://localhost:<PORT>`.

## API Routes

The API provides the following routes:

### Handle Post Routes

- `GET /posts`: Get all posts.
- `GET /post/:id`: Get a specific post by ID.
- `POST /post`: Create a new post.
- `PATCH /post/:id`: Update a post.
- `DELETE /post/:id`: Delete a post.

### Handle Comment Routes

- `GET /comments`: Get all comments (requires authentication and admin access).
- `GET /comments/:postid`: Get comments by post ID.
- `GET /comment/:id`: Get a specific comment by ID.
- `POST /comment/:postid`: Create a new comment.
- `PATCH /comment/:id`: Update a comment.
- `DELETE /comment/:id`: Delete a comment.

### Handle Authentication Routes

- `POST /signup`: User signup.
- `POST /login`: User login.
- `GET /logout`: User logout.
- `POST /createadmin`: Create an admin user (requires authentication and superadmin access).
- `GET /validate`: Validate user authentication.
- `GET /resetpassword`: Reset user password.
- `PATCH /updatepassword`: Update user password.

### Handle Users Routes

- `GET /users`: Get all users (requires authentication).
- `GET /user`: Get logged-in user (requires authentication).
- `PATCH /user`: Update user information (requires authentication).
- `GET /user/:username`: Get a specific user by username (requires authentication).
- `DELETE /user`: Delete user (requires authentication).

### Handle Special Routes

- `PATCH /follow/:tofollow`: Follow/unfollow a user (requires authentication).
- `GET /feed`: Get user's feed posts (requires authentication).
- `GET /search/:searchparam`: Search for users/posts (requires authentication).
- `GET /notifications`: Get user notifications (requires authentication).
- `GET /like/post/:postid`: Like/unlike a post (requires authentication).
- `GET /like/comment/:commentid`: Like/unlike a comment (requires authentication).

Feel free to modify the routes and their corresponding handlers according to your specific requirements.

## License

This project is licensed under the [MIT License](LICENSE).

---
