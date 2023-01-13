#!/bin/sh
set -e

source /app/app.env

echo "Run DB migration"
/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up

echo "Start the app"
exec "$@"