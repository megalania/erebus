# Erebus

## Environment

1. Create an environment file within the "dep" directory:
    ```bash
    $ touch dep/.env
    ```

2. Copy and paste the following environment structure:
    ```dotenv
    # -------------------------------------------------------------------------------------------------------- General #
    TZ="< UNIX timezone >"

    # -------------------------------------------------------------------------------------------------------------- Nginx #
    NGINX_SERVER_NAME="< Domain name >"

    # ------------------------------------------------------------------------------------------------------- Postgres #
    POSTGRES_USER="< Postgres username >"
    POSTGRES_PASSWORD="< Postgres password >"
    POSTGRES_HOST="postgres"
    POSTGRES_PORT=5432
    POSTGRES_DB="erebus"
    ```

3. Edit the dynamic fields to fit your needs.
