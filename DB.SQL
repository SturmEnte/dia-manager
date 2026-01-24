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

CREATE TABLE IF NOT EXISTS item_structures (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    attributes JSONB NOT NULL,
    general_information JSONB NOT NULL DEFAULT '[]'::jsonb,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_item_structures_user_id ON item_structures(user_id);
CREATE INDEX IF NOT EXISTS idx_item_structures_attributes_gin ON item_structures USING GIN (attributes);
CREATE INDEX IF NOT EXISTS idx_item_structures_general_info_gin ON item_structures USING GIN (general_information);

CREATE TABLE IF NOT EXISTS items (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    structure_id UUID NOT NULL REFERENCES item_structures(id) ON DELETE CASCADE,
    data JSONB NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_items_structure_id ON items(structure_id);
CREATE INDEX IF NOT EXISTS idx_items_data_gin ON items USING GIN (data);

-- Migration from 1.0.x to 1.1.x --
CREATE TABLE IF NOT EXISTS invalid_tokens (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    token TEXT UNIQUE NOT NULL,
    expires TIMESTAMPTZ NOT NULL
);

-- the old sessions table can be deletet, it's no longer used

ALTER TABLE catheters
ADD COLUMN change_reason smallint DEFAULT 0 NOT NULL;

-- Migration from 1.1.x to 1.2.x --
-- Add general_information column to item_structures if it doesn't exist
DO $$ 
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM information_schema.columns 
        WHERE table_name = 'item_structures' AND column_name = 'general_information'
    ) THEN
        ALTER TABLE item_structures 
        ADD COLUMN general_information JSONB NOT NULL DEFAULT '[]'::jsonb;
        
        CREATE INDEX IF NOT EXISTS idx_item_structures_general_info_gin 
        ON item_structures USING GIN (general_information);
    END IF;
END $$;