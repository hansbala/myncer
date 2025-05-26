# Database

This package is used for local development of myncer. In production, it uses hits a Cloud SQL
instance on GCP that I maintain.

In local development, it spins up a PostgreSQL database with auth environment setup in ENV files
in the docker image itself.
