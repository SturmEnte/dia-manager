CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS invalid_tokens (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    token TEXT UNIQUE NOT NULL,
    expires TIMESTAMPTZ NOT NULL
);

CREATE TABLE IF NOT EXISTS catheters (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    started_at TIMESTAMPTZ NOT NULL,
    ended_at TIMESTAMPTZ DEFAULT NULL,
    change_reason smallint DEFAULT 0 NOT NULL,
    CHECK (change_reason >= 0)
    CHECK (ended_at IS NULL OR ended_at > started_at)
);

-- Migration from 1.0.x to 1.1.x --
CREATE TABLE IF NOT EXISTS invalid_tokens (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    token TEXT UNIQUE NOT NULL,
    expires TIMESTAMPTZ NOT NULL
);

-- the old sessions table can be deletet, it's no longer used

ALTER TABLE catheters
ADD COLUMN change_reason smallint DEFAULT 0 NOT NULL;