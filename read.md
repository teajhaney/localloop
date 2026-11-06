# LocalLoop

LocalLoop is a community marketplace for discovering and organizing local events. People can find nearby meetups, workshops, and pop-ups, while organizers can publish events and manage signups. The project is small enough to finish as an MVP, with room to grow.

## MVP

- Browse and search events; filter by category, date, and location.
- Sign up for events or cancel a registration.
- Let organizers view attendee lists.
- Provide an organizer dashboard for creating, editing, and canceling events.
- Support authentication with attendee and organizer roles.
- Show event details, capacity, and availability.

## Stack

- **Backend:** Go, Gin, PostgreSQL, and a REST API.
- **Frontend:** React, TypeScript, TanStack Router, and TanStack Query.
- **Possible later additions:** Redis caching, email reminders, maps, and live attendee updates.

## Core data

- Users
- Events
- Categories
- Event registrations

Key engineering concerns include handling event capacity safely, enforcing role-based access, filtering events, and keeping the UI in sync with API data.

## First milestone

Build a complete vertical slice across both apps: create an event, browse it, and register an attendee.
