# Swype API (mock social media api)

### Built With GoLang, Gin and GORM ORM

This is a Golang API project built with Gin framework and using GORM as the ORM.

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
git clone https://github.com/Cedar-81/swype-api
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

5. Run imgrations:

```bash
go run ./migrate/migrate.go
```

The API will be accessible at `http://localhost:<PORT>`.

## API Routes

The API provides the following routes:

### Handle Post Routes

- `GET /posts`: Get all posts.
  - Optional Query Parameters:
    - `userid`: Retrieve posts from a particular user.
    - `page`: Pagination - page number.
    - `limit`: Pagination - number of items per page.

- `GET /post/:id`: Get a specific post by ID.
- `POST /post`: Create a new post.
- `PATCH /post/:id`: Update a post.
- `DELETE /post/:id`: Delete a post.

### Handle Comment Routes

- `GET /comments`: Get all comments (requires authentication and superadmin access).
- `GET /comments/:postid`: Get comments by post ID.
  - Optional Query Parameters:
    - `page`: Pagination - page number.
    - `limit`: Pagination - number of items per page.

- `GET /comment/:id`: Get a specific comment by ID.
- `POST /comment/:postid`: Create a new comment.
  - Optional Query Parameters:
    - `parentid`: Comment ID to indicate the comment being replied to.
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
- `GET /search/:searchparam`: Search for users/posts.
  - Optional Query Parameter:
    - `type`: Filter the search results. Options: `video`, `image`, `user`, `content`.
- `GET /notifications`: Get user notifications (requires authentication).
- `GET /like/post/:postid`: Like/unlike a post (requires authentication).
- `GET /like/comment/:commentid`: Like/unlike a comment (requires authentication).

**Note:** Routes like `/createadmin` can only be

 accessed when logged in as a superuser. A GET request to `/comments` can only be accessed by a superuser and an admin.

Feel free to modify the routes and their corresponding handlers according to your specific requirements.

## License

This project is licensed under the [MIT License](LICENSE).

---

Replace `<repository-url>` with the actual URL of your repository.
