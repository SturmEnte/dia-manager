CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS sessions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    token TEXT UNIQUE NOT NULL,
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    created TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    expires TIMESTAMPTZ NOT NULL
    CHECK (expires > created)
);

CREATE TABLE IF NOT EXISTS catheters (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    started TIMESTAMPTZ NOT NULL,
    stopped TIMESTAMPTZ DEFAULT NULL,
    CHECK (stopped IS NULL OR stopped > started),
);