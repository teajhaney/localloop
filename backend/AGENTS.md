# Backend contribution rules

## Stack

- Use Go with Gin for the REST API.
- Use PostgreSQL on Neon, accessed with `pgx` and raw, parameterized SQL. Do not add an ORM unless the project direction is explicitly changed.
- Keep database schema changes in explicit, versioned SQL migrations. Do not make ad hoc schema or data changes as part of routine development.

## Feature organization

- Organize application code by domain inside `internal/`.
- Give each feature its own package directory. For example, authentication belongs in `internal/auth/` and user functionality belongs in `internal/users/`.
- Keep a feature's service, HTTP handler, request/response types, and route registration together in that feature directory (these can be separate files within it).
- Feature packages own their domain behavior. Avoid broad `handlers/`, `services/`, or `routes/` directories that scatter one feature across the project.
- Keep genuinely shared infrastructure, such as application configuration and database connection setup, in dedicated shared packages. Avoid putting feature-specific logic there.
- Register feature routes from the API composition/startup layer; keep route definitions close to the feature they serve.

## API and data access

- Expose REST endpoints through Gin and keep HTTP concerns in handlers.
- Put business rules in feature services so they are not coupled to Gin request objects.
- Use parameterized SQL for all database values. Never build SQL by interpolating user input.
- Enforce authorization in backend code; frontend role checks are not a security boundary.
- Handle event capacity and registration changes atomically in PostgreSQL so concurrent requests cannot exceed capacity.
- Return consistent, intentional HTTP status codes and error responses. Do not expose credentials, connection strings, or internal database details to clients.

## Configuration and secrets

- Read runtime configuration from environment variables. Keep real credentials in ignored local env files or the deployment secret store.
- Never commit secrets or copy secret values into documentation, logs, or examples.
- Keep `.env.example` limited to variable names and safe placeholder values.

## Scope and verification

- Keep optional services such as Redis, email, maps, and live updates out of the MVP until requested.
- When asked to verify a backend change, run the relevant Go formatting and checks and report what ran.
