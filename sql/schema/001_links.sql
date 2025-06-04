-- +goose Up
CREATE TABLE links (id UUID PRIMARY KEY,
                        created_at TIMESTAMP NOT NULL,
                        updated_at TIMESTAMP NOT NULL,
                        link_name TEXT NOT NULL UNIQUE,
                        link_url TEXT NOT NULL);

-- +goose Down
DROP TABLE links;