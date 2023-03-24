CREATE TABLE kafka_dead_letter_queues(
    id NUMERIC PRIMARY KEY,
    topic VARCHAR NOT NULL,
    value TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT(now()),
    updated_at TIMESTAMPTZ DEFAULT(now())
)