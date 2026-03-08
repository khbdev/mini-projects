# Mini Redis (Go)

A small Redis-like server built with **Go**, using **TCP** and an in-memory **map**.

This project is a simple experiment to understand how key-value databases like Redis work.

## Run with Docker

```bash
docker pull gopherkhb/mini-redis:latest
docker run -d -p 8084:8084 --name mini-redis gopherkhb/mini-redis:latest
```

## Test

Open another terminal and connect with:

```bash
nc localhost 8084
```

Example:

```
SET name gopher


GET name gopher


DEL name gopher



```

This project is built for **learning and experimentation**.
