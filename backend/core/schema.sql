CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    -- Source of truth: Serialized User proto.
    data BYTEA NOT NULL,
    -- Used for queries.
    email VARCHAR(256) NOT NULL UNIQUE,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now()
);
