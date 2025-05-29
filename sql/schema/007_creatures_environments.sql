-- +goose Up
CREATE TABLE creatures_environments (
    creature_id UUID NOT NULL REFERENCES  creatures(id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    environment_id UUID NOT NULL REFERENCES  environments(id) ON DELETE CASCADE,
    UNIQUE(creature_id, environment_id)
);

-- +goose Down
DROP TABLE creatures_environments;