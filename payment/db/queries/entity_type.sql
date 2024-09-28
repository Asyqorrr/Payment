-- name: FindEntityById :one
SELECT entity_id, entity_name
	FROM entity_type WHERE entity_id=$1;

-- name: FindAllEntity :many
SELECT entity_id, entity_name
	FROM entity_type;

-- name: CreateEntity :one
INSERT INTO entity_type(
	 entity_name)
	VALUES ($1) RETURNING *;

-- name: UpdateEntity :one
UPDATE entity_type
	SET entity_name=$2
	WHERE entity_id=$1    
    RETURNING *;

-- name: DeleteEntity :exec
DELETE FROM entity_type
	WHERE entity_id=$1
    RETURNING *;