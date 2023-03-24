-- name: CreateKafkaDeadLetter :one
INSERT INTO kafka_dead_letter_queues (
    id,
    topic,
    value
) VALUES (
    $1, $2, $3
) RETURNING *;