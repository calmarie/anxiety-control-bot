CREATE TABLE IF NOT EXISTS states (
    id SERIAL PRIMARY KEY,
    state VARCHAR NOT NULL DEFAULT 'idle',
    stack JSONB NOT NULL DEFAULT '[]'::jsonb,
    data JSONB NOT NULL DEFAULT '{}'::jsonb,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);  

