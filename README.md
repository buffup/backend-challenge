# Buff Up Backend Challenge

As part of the recruitment process we would like to complete the following backend challenge.
We use this backend challenge to see how you go about solving a real-world problem that we have faced at Buff Up.

For more details please review [https://sportbuff.notion.site/Backend-Challenge-0344091de0b44e8bb5b39200f132e85f](Sport Buff Backend Challenge)

Please take the time to think about the code and structure of your work.

**Please try not to spend more than 5 hours on this task**, as this would not be respectful of your time.

Please host your solution on Git so that we can clone it.
However **please do not submit work as a PR as these are publicly visible**

## Task

In the `api/` directory you will find a basic `gRPC` service definition for a points and leaderboards service.
This service contains RPCs for creating and retrieving a `leaderboard`. It also includes an endpoint for awarding points.

There is also a `internal/database` directory which includes some basic boilerplate for connecting to postgres and the
necessary models but feel free to change this structure if you want.
We have used [https://github.com/golang-migrate/migrate](golang-migrate/migrate) to handle migrations.

We need you to:

- Provide a `Go` implementation of the `gRPC` service from the `api/` directory of this repo.
- Implement a `Postgres` based store for the `gRPC` service
- Provide adequate test coverage for this simple service

## How to impress us

There are a few optional tasks you can complete if you really want to show off.

1. Adding Observability
    - Here are a few examples of what you could add:
    - Structured logging
    - Tracing
    - Metrics

2. Adding Configuration and Secrets management

3. Think about the limitations of your solution and how you could improve in the future
    - Document these in either the `README.md` or in comments alongside the code
