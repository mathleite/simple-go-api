<h1 align="center">Simple Go Api</h1>

## Setup

Just **build** docker with:
```bash
docker-compose up --build -d
```

## Server

To start:
* Enter inside `golang` container:
	```bash
	docker-compose exec golang bash
	```
* Build and run serve
	```go
	go build && ./app
	```

The server will be allocate on `http://localhost/api/books`

Config:
* Header: `Content-Type application/json`

## How it's works?

You can use:
- **POST** to create Books;
- **GET** to get Books and specific Book by **ID**;
- **DELETE** to delete Book by **ID**;
- **PUT** to update Book by **ID**

#### Example

**Create** Book:
```json
{
	"isbn": "432432",
	"title": "Andorinha",
	"author": {
		"firstname": "Carol",
		"lastname": "Williams"
	}
}
```