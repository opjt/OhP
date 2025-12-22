-- users
CREATE TABLE
    users (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid_v7(),
        email TEXT NOT NULL UNIQUE,
        created_at TIMESTAMP NOT NULL DEFAULT now(),
        updated_at TIMESTAMP
    );