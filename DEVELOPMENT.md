# Developer Setup for Comic Handler

This project uses Docker for development with hot-reloading enabled for both the Angular frontend and Rust backend. It can also be run locally.

## Prerequisites

- [Docker and Docker Compose](https://docs.docker.com/compose/install/) (Recommended for full stack dev)
- [Node.js](https://nodejs.org/en/download/) (v25+ recommended for Angular 21) causes warnings but works, officially LTS is preferred.
- [Rust](https://www.rust-lang.org/tools/install) (for backend development)

## Development Workflow

### Option 1: Docker (Recommended)

This method ensures you have a consistent environment for both frontend and backend with hot-reloading.

#### Using VS Code Tasks

1. Open the project in VS Code
2. Press `Ctrl+Shift+P` (or `Cmd+Shift+P` on macOS)
3. Type "Tasks: Run Task" and select it
4. Choose "Start Development Environment"

**Available Tasks:**
- "Stop Development Environment" - Stop all running containers
- "Rebuild and Restart Backend" - Rebuild and restart just the backend
- "Rebuild and Restart Frontend" - Rebuild and restart just the frontend
- "View Backend Logs" - View the Rust backend logs
- "View Frontend Logs" - View the Angular frontend logs

#### Manual Docker Startup

```bash
docker compose -f docker-compose.dev.yml up --build
```

### Option 2: Local Development

If you prefer running services directly on your machine.

#### Frontend (Angular 21)

1.  Navigate to the directory:
    ```bash
    cd frontend
    ```
2.  Install dependencies:
    ```bash
    npm install
    ```
3.  Run the application:
    ```bash
    npm datetime
    npm start
    ```
    This runs `ng serve`.

#### Backend (Rust)

1.  Navigate to the directory:
    ```bash
    cd backend
    ```
2.  Run the server:
    ```bash
    cargo run
    ```
    Or use `cargo watch` for hot-reloading:
    ```bash
    cargo install cargo-watch
    cargo watch -x run
    ```

### Accessing the Applications

- **Frontend**: [http://localhost:4200](http://localhost:4200)
- **Backend API**: [http://localhost:9999](http://localhost:9999)

## Testing

### Frontend Tests

To run unit tests for the Angular frontend:

```bash
cd frontend
npm run test
```

> **Note**: If you are in an environment without a local Chrome installation (e.g., some containers), the project is configured to use Puppeteer to provide a Chromium binary automatically.

## Debugging

- **Frontend**: Use the "Launch Chrome against localhost" configuration in VS Code.
- **Backend**: View logs via Docker or standard stdout if running locally.

## Updating Dependencies

When updating versions or dependencies, you may need to clear artifacts:

```bash
# Clear frontend cache
rm -rf frontend/node_modules frontend/package-lock.json

# Clear backend cache
rm -rf backend/target

# Clear Docker images
docker image prune -f
```
