# Use the official lightweight PostgreSQL image
FROM postgres:15-alpine

# Set environment variables for default DB, user, and password
ENV POSTGRES_DB=myncer \
    POSTGRES_USER=devuser \
    POSTGRES_PASSWORD=devpass

# Optional: copy an initial schema if you want
# COPY ./schema.sql /docker-entrypoint-initdb.d/

# Expose the default PostgreSQL port
EXPOSE 5432
