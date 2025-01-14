-- name: InsertHero :exec
INSERT INTO hero (
    id,
    name,
    owner,
    created_at,
    updated_at
) VALUES (
    @id,
    @name,
    @owner,
    @created_at,
    @updated_at
);
