set dotenv-load
DB_MIGRATION_FILE:="./database/migrations/schema.sql"

test:
    go test -v ./...

db_generate:
    sqlc generate

# up database
db_up:
    docker compose up -d

db_schema: db_up
    @echo $DATABASE_URL
    atlas schema inspect -u $DATABASE_URL

#
# dev-url: dev database is required by atras. It's used temporaly and locally for sql validation.
# https://atlasgo.io/concepts/dev-database
db_migrate: db_up
    atlas schema apply \
        --url $DATABASE_URL \
        --dev-url "docker://postgres/16/dev" \
        --to file://{{DB_MIGRATION_FILE}}

# run application
r: db_up
    go run .
