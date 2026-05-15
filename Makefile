.PHONY: help \
        dev dev/down dev/logs dev/clean \
        build build/backend build/worker build/cli build/frontend \
        lint lint/go lint/frontend \
        test test/backend test/worker test/frontend \
        generate generate/backend generate/frontend \
        clean

# ---------------------------------------------------------------------------
# Vydon developer Makefile.
# `make dev` brings the full local stack up; `make dev/down` tears it down.
# ---------------------------------------------------------------------------

DEV_COMPOSE_FILE = compose.dev.yml

help: ## Show this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nVydon Makefile targets:\n\n"} \
		/^[a-zA-Z_\/.\-]+:.*##/ { printf "  \033[36m%-22s\033[0m %s\n", $$1, $$2 }' \
		$(MAKEFILE_LIST)
	@printf "\nQuick start:  \033[36mmake dev\033[0m  then open http://localhost:3000\n\n"

# Local dev stack ----------------------------------------------------------

dev: ## Bring the full local stack up (db, temporal, backend, worker, app).
	docker compose -f $(DEV_COMPOSE_FILE) up --build -d
	@printf "\n\033[32mStack is up. Tail logs with: make dev/logs\033[0m\n"
	@printf "  app:      http://localhost:3000\n"
	@printf "  api:      http://localhost:8080\n"
	@printf "  temporal: http://localhost:8233\n"

dev/down: ## Stop the local stack and remove containers.
	docker compose -f $(DEV_COMPOSE_FILE) down

dev/logs: ## Stream logs from every container in the stack.
	docker compose -f $(DEV_COMPOSE_FILE) logs -f --tail=200

dev/clean: ## Stop the stack and wipe volumes (DESTRUCTIVE).
	docker compose -f $(DEV_COMPOSE_FILE) down --volumes --remove-orphans

# Build ---------------------------------------------------------------------

build: build/backend build/worker build/cli build/frontend ## Build every component.

build/backend: ## Build the backend binary.
	@cd backend && go build ./...

build/worker: ## Build the worker binary.
	@cd worker && go build ./...

build/cli: ## Build the cli binary.
	@cd cli && go build ./...

build/frontend: ## Install and build the frontend.
	@cd frontend && npm ci && npm run build

# Lint ---------------------------------------------------------------------

lint: lint/go lint/frontend ## Run all linters.

lint/go: ## Lint Go.
	@command -v golangci-lint >/dev/null || { echo "golangci-lint required"; exit 1; }
	golangci-lint run ./...

lint/frontend: ## Lint frontend.
	@cd frontend && npm run lint

# Test ---------------------------------------------------------------------

test: test/backend test/worker test/frontend ## Run all tests.

test/backend: ## Run backend tests.
	@cd backend && go test ./...

test/worker: ## Run worker tests.
	@cd worker && go test ./...

test/frontend: ## Run frontend tests.
	@cd frontend && npm test

# Code generation ---------------------------------------------------------

generate: generate/backend generate/frontend ## Regenerate proto + sqlc + mocks.

generate/backend: ## Regenerate backend code (proto, sqlc, mocks).
	@cd backend && make generate

generate/frontend: ## Regenerate frontend client code.
	@cd frontend && npm run generate

# Clean ---------------------------------------------------------------------

clean: ## Remove build artefacts.
	@cd backend && make clean 2>/dev/null || true
	@cd worker && make clean 2>/dev/null || true
	@cd cli && make clean 2>/dev/null || true
	rm -rf frontend/.next frontend/out frontend/apps/web/.next
