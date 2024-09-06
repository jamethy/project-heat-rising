-- +goose Up
create schema if not exists prh;

create table if not exists prh.thermostat
(
    id            serial primary key,
    created_at    timestamp not null default now(),
    timestamp     timestamp not null,
    provider      text      not null,
    thermostat_id text      not null,
    target_cool   float     not null,
    target_heat   float     not null,
    actual_temp   float     not null,
    humidity      float     not null,
    is_heating    bool      not null,
    is_cooling    bool      not null
);

create table if not exists prh.weather
(
    id             serial primary key,
    created_at     timestamp not null default now(),
    timestamp      timestamp not null,
    provider       text not null,
    temperature    float not null,
    feels_like     float not null,
    pressure       float not null,
    humidity       float not null,
    wind_speed     float not null,
    wind_direction float not null,
    clouds         float not null
);

-- +goose Down
drop table if exists prh.weather;
drop table if exists prh.thermostat;
