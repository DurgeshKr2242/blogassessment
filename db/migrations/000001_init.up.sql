CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create blog_posts table
CREATE TABLE IF NOT EXISTS blog_posts (
    id          UUID            NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
    title       VARCHAR(255)    NOT NULL,
    description TEXT,
    body        TEXT,
    created_at  TIMESTAMP WITHOUT TIME ZONE     DEFAULT NOW(),
    updated_at  TIMESTAMP WITHOUT TIME ZONE     DEFAULT NOW()
);