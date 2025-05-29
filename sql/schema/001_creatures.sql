-- +goose Up
CREATE TABLE creatures (
    id UUID UNIQUE NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL,
    exp INT NOT NULL,
    book TEXT,
    page INT
);

-- +goose Down
DROP TABLE creatures;