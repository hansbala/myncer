CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    -- Source of truth: Serialized User proto.
    data BYTEA NOT NULL,
    -- Used for queries.
    email VARCHAR(256) NOT NULL UNIQUE,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now()
);

CREATE TABLE IF NOT EXISTS datasource_tokens (
  id UUID PRIMARY KEY,
  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  datasource VARCHAR(64) NOT NULL,
  -- Source of truth: Serialized OAuthToken proto.
  data BYTEA NOT NULL,
  created_at TIMESTAMPTZ DEFAULT now(),
  updated_at TIMESTAMPTZ DEFAULT now(),
  UNIQUE (user_id, datasource)
);

CREATE TABLE IF NOT EXISTS syncs (
  id UUID PRIMARY KEY,
  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  -- Source of truth: Serialized Sync proto.
  data BYTEA NOT NULL,
  created_at TIMESTAMPTZ DEFAULT now(),
  updated_at TIMESTAMPTZ DEFAULT now()
);
