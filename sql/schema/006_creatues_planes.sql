-- +goose Up
CREATE TABLE creatures_planes (
    creature_id UUID NOT NULL REFERENCES  creatures(id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    plane_id UUID NOT NULL REFERENCES  planes(id) ON DELETE CASCADE,
    UNIQUE(creature_id, plane_id)
);

-- +goose Down
DROP TABLE creatures_planes;