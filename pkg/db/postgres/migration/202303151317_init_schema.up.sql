create table chains
(
    id         NUMERIC primary key not null,
    name       varchar(255),
    symbol     varchar(255),
    rpc        json,
    updated_at TIMESTAMPTZ DEFAULT(now()),
    created_at TIMESTAMPTZ DEFAULT(now())
);