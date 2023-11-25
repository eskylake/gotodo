# GoTodo

Very simple Todo API using Go Fiber.

## Installation

```bash
docker compose -f docker-compose.dev.yml up -d
```

## Usage

```bash
mv .env.example .env
```
Open `.env` and make changes.

```bash
docker exec -t go-todo-api air
```

Serves on [localhost:{GOTODO_API_HTTP_PORT}](http://localhost:4000/api)

For example [GET Todos](http://localhost:4000/api/todos)

## Build

```bash
docker build -t gotodoapi -f Dockerfile.prod .
```

## License

[MIT](https://choosealicense.com/licenses/mit/)