-- Simple schema and seed for local testing
CREATE TABLE IF NOT EXISTS items (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

INSERT INTO items (name) VALUES ('sample item') ON CONFLICT DO NOTHING;
