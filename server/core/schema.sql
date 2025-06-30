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

CREATE TABLE IF NOT EXISTS sync_runs (
  run_id UUID PRIMARY KEY,
  sync_id UUID NOT NULL REFERENCES syncs(id) ON DELETE CASCADE,
  -- Source of truth: Serialized SyncRun proto.
  data BYTEA NOT NULL,
  created_at TIMESTAMPTZ DEFAULT now(),
  updated_at TIMESTAMPTZ DEFAULT now()
);

CREATE TABLE IF NOT EXISTS songs (
  -- Unique myncer song id.
  id UUID PRIMARY KEY,
  -- Source of truth: Serialized Song proto.
  data BYTEA NOT NULL,
  -- Useful fields for faster queries.
  -- Myncer datasource name.
  datasource VARCHAR(256) NOT NULL,
  -- The respective datasource's unique, stable id.
  datasourceSongId VARCHAR(256) NOT NULL,
  -- Metadata leveraging SQL for ACID compliance.
  created_at TIMESTAMPTZ DEFAULT now(),
  updated_at TIMESTAMPTZ DEFAULT now()
);
