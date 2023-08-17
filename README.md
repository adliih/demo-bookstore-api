# demo-bookstore-api

Welcome to the Bookstore API project! This repository contains a RESTful API built using Go that allows you to manage a collection of books. Whether you're a book enthusiast or a developer looking to learn Go, this project provides a practical example of how to design and implement a web API.

## Table of Contents

- [demo-bookstore-api](#demo-bookstore-api)
  - [Table of Contents](#table-of-contents)
  - [Features](#features)
  - [Getting Started](#getting-started)
    - [Prerequisites](#prerequisites)
  - [Usage Examples](#usage-examples)
  - [Contributing](#contributing)
  - [License](#license)

## Features

- Create, Read, Update, and Delete (CRUD) operations for managing books.
- Error handling and validation for robust data processing.
- SQLite database integration for persistent storage.
- RESTful API design with clear and intuitive endpoints.

## Getting Started

### Prerequisites

- Go (Golang) installed. [Download Go](https://golang.org/dl/)
- Git installed. [Download Git](https://git-scm.com/downloads)

<!-- ### Running the API -->

<!-- ## API Documentation -->

## Usage Examples

1. Retrieve all books:

```sh
curl http://localhost:8080/api/v1/books
```

1. Create a new book:

```sh
curl -X POST http://localhost:8080/api/v1/books -d '{"title": "Sample Book", "author": "John Doe", "genre": "Fiction"}'
```

1. Update a book

```sh
curl -X PUT http://localhost:8080/api/v1/books/{bookID} -d '{"title": "Updated Book Title"}'
```

1. Delete a book

```sh
curl -X DELETE http://localhost:8080/api/v1/books/{bookID}
```

## Contributing

Contributions are welcome! If you find a bug or have an enhancement in mind, feel free to open an issue or submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).
