# Frontend contribution rules

## Stack

- Use React and TypeScript with TanStack Router for routing and TanStack Query for server state and API data.
- Use the existing project dependencies and patterns before adding another library.

## Organization and responsibilities

- Group UI code by product feature or route so event discovery, event details, registrations, authentication, and organizer tools remain easy to locate.
- Keep reusable, feature-neutral UI components separate from feature-specific components. Avoid putting application behavior into generic shared components.
- Keep API calls and server data interactions in clear feature-oriented modules and use TanStack Query for fetching, caching, and updating server state.
- Use URL search parameters for shareable event browse filters such as category, date, and location.
- Keep form input and API data validated with the existing TypeScript and Zod tools where appropriate.

## User experience and data behavior

- Keep the interface responsive and accessible, including keyboard navigation, semantic controls, labels, and useful loading, empty, and error states.
- Treat the backend as the source of truth for authentication, attendee/organizer permissions, capacity, and registration status. UI checks improve the experience but do not replace server authorization.
- Keep UI data in sync after mutations by updating or invalidating the relevant TanStack Query data.
- Follow the LocalLoop design references in `../stitch_localloop_risograph_design_system/` when implementing visual design.
- Keep the MVP focused on discovering events, free registration, and organizer event management. Defer optional maps, email reminders, live updates, and other additions until requested.

## Configuration and verification

- Keep environment-specific values out of source control. Do not expose backend secrets in browser code; only intentionally public frontend configuration may use `VITE_` variables.
- Prefer the existing npm scripts and patterns. When asked to verify a frontend change, run the relevant checks and report what ran.
