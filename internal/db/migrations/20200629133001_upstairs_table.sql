-- +goose Up
create table if not exists prh.upstairs
(
    id          serial primary key,
    created_at  timestamp not null default now(),
    timestamp   timestamp not null,
    provider    text not null,
    temperature float not null,
    pressure    float not null,
    humidity    float not null
);

-- +goose Down
drop table if exists prh.upstairs;