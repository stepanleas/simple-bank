#!/bin/sh

set -e

echo "run db migration"
# Only source app.env if DB_SOURCE is not already set
if [ -z "$DB_SOURCE" ]; then
  source /app/app.env
fi

/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up

echo "start the application"
exec "$@"