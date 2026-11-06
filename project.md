# LocalLoop Project

## Overview

LocalLoop is a local event discovery and RSVP product. The repository currently contains the UI design references and an initial Go API foundation. The planned experience covers browsing events, viewing event details, free RSVPs, account settings, and organizer event management. Paid ticketing and payment flows are out of scope for the current product direction.

## Current stack

- **Frontend:** React 19, TypeScript, Vite, TanStack Router/Start, TanStack Query, React Hook Form, and Zod (`frontend/`).
- **Backend:** Go, Gin, and `pgx` (`backend/`).
- **Database:** Neon hosted PostgreSQL, accessed with raw SQL through `pgx`; no ORM is configured.
- **Design references:** `stitch_localloop_risograph_design_system/`.

## Backend setup

The API reads `DATABASE_URL` from the process environment or a local `.env` file. Keep the Neon connection string private. The repository ignores `.env` and `backend/.env`; use the checked-in `backend/.env.example` as a template and place your real connection URL in `backend/.env`.

From the repository root:

```sh
cp backend/.env.example backend/.env
# Replace DATABASE_URL in backend/.env with the Neon connection string.
cd backend
go run ./cmd/api
```

The API listens on port `8080` by default. `GET /health` returns the API and database health status. The database pool can be tuned with `DATABASE_MAX_CONNS`, `DATABASE_MIN_CONNS`, `DATABASE_MAX_CONN_LIFETIME_SECONDS`, and `DATABASE_MAX_CONN_IDLE_SECONDS`. `DATABASE_HEALTH_TIMEOUT_SECONDS` controls the health check ping timeout.

Neon connection URLs should use TLS (typically `sslmode=require`). Use the pooled connection string from Neon for application traffic when appropriate. The current startup health check opens and pings a connection; it does not create tables or write records.

## Database approach

Use PostgreSQL directly through `pgx` and parameterized SQL. Keep schema changes explicit and versioned as SQL migrations when the application schema is defined. No schema, seed data, or database mutations are part of the current setup.

## Frontend setup

```sh
cd frontend
npm install
npm run dev
```

Available frontend scripts are `dev`, `build`, and `preview`.

## Current implementation status

- The backend loads configuration, connects to PostgreSQL, and exposes a database-aware health endpoint.
- Frontend app source files and API routes have not yet been added to the repository.
- No application tables or migrations have been created.
