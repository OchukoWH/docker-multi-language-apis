# Docker Multi-Language APIs

A teaching project demonstrating how to build and run applications in different programming languages without installing those languages or their runtimes on the host machine. Only Docker and the Docker CLI are required.

## What This Project Demonstrates

This project shows that Docker can provide complete runtime environments for applications written in any programming language. The host machine only needs Docker installed — no Node.js, Python, Go, npm, or pip is required. Each language's runtime comes from its official Docker image.

## Why Node.js, Python, and Go Don't Need to Be Installed on the Host

- **Docker images include everything needed**: Each official Docker image contains the language runtime, standard libraries, and all dependencies required to run the application.
- **Isolation**: Containers run in isolated environments with their own file systems, processes, and network interfaces.
- **Consistency**: The application runs the same way regardless of the host operating system, because the language runtime is bundled inside the Docker image.
- **Portability**: The same Docker images can run on any machine with Docker installed — Linux, macOS, or Windows.

## Docker Images Used

| Language | Official Image |
|----------|---------------|
| Node.js | `node:20-alpine` |
| Python | `python:3.12-slim` |
| Go | `golang:1.23-alpine` (build stage) + `alpine:3.20` (runtime stage) |

## Building Each Image Individually

```bash
# Node.js
cd docker-multi-language-apis/node-api
docker build -t node-hello-api .

# Python
cd docker-multi-language-apis/python-api
docker build -t python-hello-api .

# Go
cd docker-multi-language-apis/go-api
docker build -t go-hello-api .
```

## Running Each Container Individually

### Node.js

```bash
docker run -d --name node-hello -p 3000:3000 node-hello-api
```

### Python

```bash
docker run -d --name python-hello -p 8000:8000 python-hello-api
```

### Go

```bash
docker run -d --name go-hello -p 8080:8080 go-hello-api
```

All three containers can run simultaneously because they use different ports (3000, 8000, 8080).

## Testing Each API

### Node.js

```bash
curl http://localhost:3000/
# Returns: "Hello from the Node.js Docker container!"

curl http://localhost:3000/health
# Returns: "OK"
```

### Python

```bash
curl http://localhost:8000/
# Returns: "Hello from the Python Docker container!"

curl http://localhost:8000/health
# Returns: "OK"
```

### Go

```bash
curl http://localhost:8080/
# Returns: "OK"

curl http://localhost:8080/health
# Returns: "OK"
```

## Viewing Container Logs

```bash
docker logs node-hello
docker logs python-hello
docker logs go-hello
```

## Stopping and Removing Containers

```bash
docker stop node-hello python-hello go-hello
docker rm node-hello python-hello go-hello
```

## Removing the Images

```bash
docker rmi node-hello-api python-hello-api go-hello-api
```

## Verifying Containers Are Running

```bash
docker ps
```

Expected output shows three containers with the names `node-hello`, `python-hello`, and `go-hello`, all with their respective ports up.

## Running Everything Together (Full Demonstration)

```bash
# Build all images
cd docker-multi-language-apis/node-api && docker build -t node-hello-api .
cd docker-multi-language-apis/python-api && docker build -t python-hello-api .
cd docker-multi-language-apis/go-api && docker build -t go-hello-api .

# Run all containers simultaneously
docker run -d --name node-hello -p 3000:3000 node-hello-api
docker run -d --name python-hello -p 8000:8000 python-hello-api
docker run -d --name go-hello -p 8080:8080 go-hello-api

# Test all endpoints
curl http://localhost:3000/
curl http://localhost:3000/health
curl http://localhost:8000/
curl http://localhost:8000/health
curl http://localhost:8080/
curl http://localhost:8080/health

# View logs
docker logs node-hello
docker logs python-hello
docker logs go-hello

# Stop and clean up
docker stop node-hello python-hello go-hello
docker rm node-hello python-hello go-hello
docker rmi node-hello-api python-hello-api go-hello-api
```