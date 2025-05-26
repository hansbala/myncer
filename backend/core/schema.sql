CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    -- Serialized User proto.
    data BYTEA NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now()
);
