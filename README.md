# TeenyTinyURL

A simple URL shortener for my domain **hmd.sh**.  
Using it to learn building APIs in go using the chi router with pgx and pgxpool for interacting with PostgreSQL and pooling. With a redis layer for caching.

This implementation won't allow external users to create shortened URLs, only the webmaster will have the ability.

## Planned Features
1. Implementing a fully featured chi routes with middleware for authentication and authorization
2. Adding an API token system to authenticate people and determine who can create shortened URLs with ability to revoke and create more
3. Adding a /healthz for monitoring the health of the service
4. Adding a redis caching layer to handle cached responses
5. Adding a PostgreSQL database with goose for migrations to hold all info about URLs, user, API tokens, and more.
6. Ability to add, remove, edit shortened URLs and tokens
7. Adding a docker file to setup the service easily with a docker compose file too
8. Will include the whole hosting apparatus with caddy and such
9. Learn and include a github workflow