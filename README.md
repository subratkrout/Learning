# Learning

This repository contains a minimal scaffolding for a 3-tier application:

- `api/` - Flask-based API service (Dockerized)
- `ui/` - Static UI served by `nginx` (Dockerized)
- `db/` - Postgres initialization SQL

Quick start (requires Docker & Docker Compose):

```bash
cp .env.example .env
docker-compose up --build
```

- UI: http://localhost:8080
- API health: http://localhost:8000/health

Files added as scaffolding: `docker-compose.yml`, `.env.example`, `api/`, `ui/`, `db/init.sql`.

Next steps: implement API routes, UI pages, and database migrations as needed.

This branch now contains a Go-based API and a Next.js + Tailwind UI.

Run the full stack with Docker Compose:

```bash
cp .env.example .env
docker-compose up --build
```

Notes:
- API: `api` service runs on container port 8000 and exposes `/health` and `/items`.
- UI: `ui` service is a Next.js app which connects to the API via internal Docker network.
- DB: `db/init.sql` seeds a simple `items` table used by the API.

Shadcn UI: the `ui/` app includes Tailwind and is ready for `shadcn/ui` integration. To add shadcn components, run `npx shadcn-ui@latest init` inside `ui/` and follow the prompts.