# Makefile

# Update this with your actual database connection string.
DB_URL=postgres://SYS:aegleadmin@localhost:5432/blogassessment?sslmode=disable

# Directory where migration files are stored.
MIGRATION_DIR=./db/migrations

# Command alias for migrate.
MIGRATE_CMD=migrate -path $(MIGRATION_DIR) -database "$(DB_URL)"

.PHONY: migrate-up migrate-down migrate-new

# Run all "up" migrations.
migrate-up:
	$(MIGRATE_CMD) up

# Rollback (run one "down" migration).
migrate-down:
	$(MIGRATE_CMD) down

# Create a new migration file.
# Usage: make migrate-new name=add_new_feature
migrate-new:
	migrate create -ext sql -dir $(MIGRATION_DIR) $(name)