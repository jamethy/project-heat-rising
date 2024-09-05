-- +goose Up
create table if not exists prh.upstairs
(
    id          serial primary key,
    created_at  timestamp not null default now(),
    timestamp   timestamp not null,
    provider    text,
    temperature float(8),
    pressure    float(8),
    humidity    float(8)
);

-- +goose Down
drop table if exists prh.upstairs;