#!/bin/bash
set -e # Exit immediately if a command exits with a non-zero status.

# The 'psql' command is available because this script is run by the official postgres image.
# We use the '-v ON_ERROR_STOP=1' flag to ensure that the script will exit if any command fails.
# The script is executed as the 'postgres' user, which has superuser privileges.
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE USER $POSTGRES_DB_USER WITH PASSWORD '$POSTGRES_DB_PWD';

    GRANT ALL PRIVILEGES ON DATABASE $POSTGRES_DB TO $POSTGRES_DB_USER;

    GRANT USAGE, CREATE ON SCHEMA public TO $POSTGRES_DB_USER;
EOSQL
