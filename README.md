# Cloud Assignment

Why this approach is smarter (if done properly)

You are building in layers:

Layer 1 → HTTP server
Layer 2 → Docker
Layer 3 → AWS EC2 or ECS
Layer 4 → CI/CD (GitHub Actions)
Layer 5 → Reverse proxy + domain
Layer 6 → Then database

This way you learn:

Networking

Ports

Security groups

Deployment

Infrastructure thinking

That’s real backend engineering.

## Commands

```bash
go run cmd/server/main.go

docker build -t cloud_assignment .
docker run -p 8080:8080 cloud_assignment.

docker ps
docker stop <container ID>
```
