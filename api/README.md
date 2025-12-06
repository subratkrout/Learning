Go API scaffold

- Dockerfile: builds the Go API (multi-stage)
- Endpoints:
	- `GET /health` - returns API and DB status
	- `GET /users` - list users (JSON array)
	- `POST /users` - create user with JSON {"name":"...","dob":"YYYY-MM-DD"}

Run with Docker Compose:
```
cp .env.example .env
docker-compose up --build api
```

The service reads `DATABASE_URL` or `POSTGRES_*` env vars to connect to Postgres. It uses GORM for migrations and DB access.
