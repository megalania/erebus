CREATE TABLE IF NOT EXISTS agent (
    id         UUID      PRIMARY KEY DEFAULT uuid_generate_v4 (),
    created_at TIMESTAMP NOT NULL    DEFAULT NOW (),
    updated_at TIMESTAMP NOT NULL    DEFAULT NOW ()
);
