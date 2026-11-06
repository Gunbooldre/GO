include_env = set -a; [ -f .env ] && . ./.env; set +a;

.PHONY: run

run:
	@$(include_env) go run ./cmd/web
