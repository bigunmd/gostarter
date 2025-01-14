CREATE TABLE IF NOT EXISTS hero (
    id TEXT NOT NULL,
    name TEXT NOT NULL,
    owner TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    PRIMARY KEY (id),
    UNIQUE (name, owner)
);
