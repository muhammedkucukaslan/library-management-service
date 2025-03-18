<h1 align="center">Library Management Service API</h1>

<p align="center">
  A modern Go-based RESTful API for efficient library resource management
  <br />
  <br />
<img src="public/home.png" alt="Home" width="600">

</p>

## About The Project
This project simplifies library management with RESTful API. Built with Go, it includes user authentication, role-based access control, email notifications, and streamlined book borrowing and returning.

## Features

- **Role-Based Access Control (RBAC):** Fine-grained access control to manage user permissions effectively.
- **Custom Authentication System:** Secure user authentication with JWT-based token management from scratch.
- **Email Notifications:** User gets email while signing up and getting punished.
- **Book Borrowing and Returning:** Seamless workflows for managing book loans and returns.

### Built With

 ![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Fiber](https://img.shields.io/badge/Fiber-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-316192?style=for-the-badge&logo=postgresql&logoColor=white)
![JWT](https://img.shields.io/badge/JWT-000000?style=for-the-badge&logo=jsonwebtokens&logoColor=white)




## Getting Started

Follow these instructions to set up the project locally.

### Prerequisites

* Go 1.18+
* PostgreSQL 14+
* [Resend.com](https://resend.com/) API KEY 

You can prefer any database and email service provider. In order to do that,you just need to change `infrastructure` layer and `main.go` file.

### Installation

1. Clone the repository
   ```sh
   git clone https://github.com/muhammedkucukaslan/library-management-service.git
   ```

2. Navigate to the project directory
   ```sh
   cd library-management-service
   ```

3. Install dependencies
   ```sh
   go mod download
   ```

4. Create a `.env` file in the root directory with following configuration
   ```sh
   DATABASE_URL=postgres://user:password@localhost:5432/library_db?sslmode=disable
   JWT_SECRET_KEY=your_secret_key
   RESEND_API_KEY=your_resend_api_key
   EMAIL_SERVICE_DOMAIN=domain@example.com
    ```
5. Build and run the application (be sure PostgreSQL running)
   ```sh
   go build -o library-api
   ./library-api
   ```



## Database Schema

I probably change the schema in the future, but for now, it is as follows:

<img src="public/image.png" alt="Library Management Service" width="600">


## API Endpoints

### Base Endpoints
- `GET /healthcheck` - System health check

### Authentication
- `POST /api/signup` - Register a new user
- `POST /api/login` - Authenticate a user

### User Management
- `GET /api/users/me` - Get current user information
- `GET /api/users/:id` - Get specific user information
- `DELETE /api/users/:id` - Delete a user
- `PUT /api/users/:id` - Update user information
- `POST /api/users/:id/punish` - Add punishment to a user

### Author Management
- `POST /api/authors/` - Create a new author
- `DELETE /api/authors/:id` - Delete an author

### Book Management
- `POST /api/books/` - Create a new book
- `GET /api/books/` - List all books
- `GET /api/books/:id` - Get specific book information
- `DELETE /api/books/:id` - Delete a book

### Category Management
- `POST /api/categories/` - Create a new category
- `DELETE /api/categories/:id` - Delete a category
- `GET /api/categories/` - List all categories

### Loan Management
- `POST /api/loans/` - Borrow a book
- `GET /api/loans/` - List all loans
- `POST /api/loans/:id/return` - Return a borrowed book

> **Note**: All endpoints under `/api` require authentication except for `/api/signup` and `/api/login`

## Project Structure

See [file_tree.txt](https://github.com/muhammedkucukaslan/library-management-service/blob/main/file_tree.txt) for a detailed project structure.


## Issues or Feature Requests
Check out the [open issues](https://github.com/muhammedkucukaslan/library-management-service/issues) to see what others are working on or to find areas where you can contribute. If you have an idea for a new feature or spot a bug, feel free to [create a new issue](https://github.com/muhammedkucukaslan/library-management-service/issues/new/choose) to share your suggestions or report problems.