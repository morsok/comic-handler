# Docker Compose Command Reference

## IMPORTANT: Always use the new Docker Compose syntax

### ✅ CORRECT - Use `docker compose` (with space)

```bash
docker compose up
docker compose down
docker compose build
docker compose -f docker-compose.dev.yml up --build
docker compose -f docker-compose.dev.yml restart backend
docker compose -f docker-compose.dev.yml logs -f frontend
```

### ❌ INCORRECT - Don't use `docker-compose` (with hyphen)

```bash
# These are DEPRECATED and should NOT be used:
docker-compose up
docker-compose down
docker-compose build
```

## Why the change

- `docker-compose` was the old standalone tool
- `docker compose` is the new plugin integrated into Docker CLI (Docker Compose V2)
- The old `docker-compose` command is deprecated and may not be available on newer systems

## Remember

- Always use `docker compose` (space) in scripts
- Always use `docker compose` (space) in documentation
- Always use `docker compose` (space) in VS Code tasks
- This is the modern, supported way to use Docker Compose

## Quick Check

If you see `docker-compose` anywhere in the codebase, it should be changed to `docker compose`!
