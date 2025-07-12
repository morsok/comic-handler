# Developer Setup for Comic Handler

This project uses Docker for development with hot-reloading enabled for both the Angular frontend and Rust backend.

## Prerequisites

- Docker and Docker Compose installed
- VS Code (for integrated development tasks)

## Development Workflow

### Using VS Code Tasks

1. Open the project in VS Code
2. Press `Ctrl+Shift+P` (or `Cmd+Shift+P` on macOS)
3. Type "Tasks: Run Task" and select it
4. Choose "Start Development Environment" to start both frontend and backend with hot-reloading

Other available tasks:

- "Stop Development Environment" - Stop all running containers
- "Rebuild and Restart Backend" - Rebuild and restart just the backend
- "Rebuild and Restart Frontend" - Rebuild and restart just the frontend
- "View Backend Logs" - View the Rust backend logs
- "View Frontend Logs" - View the Angular frontend logs

### Manual Startup

To start the development environment manually:

```bash
docker compose -f docker-compose.dev.yml up --build
```

### Accessing the Applications

- Frontend: [http://localhost:4200](http://localhost:4200)
- Backend API: [http://localhost:9999](http://localhost:9999)

## Hot Reloading

- **Frontend**: Changes to Angular files will automatically trigger a reload
- **Backend**: Changes to Rust files will automatically rebuild and restart the server

## Debugging

For debugging the frontend, you can use the "Launch Chrome against localhost" configuration in VS Code's Run and Debug panel.

For backend debugging, you can view the logs using the "View Backend Logs" task.

## Updating to Latest Versions

When updating to latest versions, clear the caches first:

```bash
# Stop containers
docker compose -f docker-compose.dev.yml down

# Clear frontend cache
rm -rf frontend/node_modules frontend/package-lock.json

# Clear backend cache  
rm -rf backend/target

# Clear Docker images (optional)
docker image prune -f
```

Then start the development environment to rebuild with latest versions.
