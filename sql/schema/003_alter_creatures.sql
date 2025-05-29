-- +goose Up
ALTER TABLE creatures
ADD COLUMN type_id UUID NOT NULL references types(id) ON DELETE RESTRICT; 

-- +goose Down
ALTER TABLE creatures
DROP COLUMN type_id;