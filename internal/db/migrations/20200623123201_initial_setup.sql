-- +goose Up
create schema if not exists prh;

create table if not exists prh.thermostat
(
    id            serial primary key,
    created_at    timestamp not null default now(),
    timestamp     timestamp not null,
    provider      text,
    thermostat_id text,
    target_cool   float(8),
    target_heat   float(8),
    actual_temp   float(8),
    humidity      float(8),
    is_heating    bool,
    is_cooling    bool
);

create table if not exists prh.weather
(
    id             serial primary key,
    created_at     timestamp not null default now(),
    timestamp      timestamp not null,
    provider       text,
    temperature    float(8),
    feels_like     float(8),
    pressure       float(8),
    humidity       float(8),
    wind_speed     float(8),
    wind_direction float(8),
    clouds         float(8)
);

create table if not exists prh.daily_data
(
    id         serial primary key,
    created_at timestamp not null default now(),
    date       date      not null,
    sunrise    timestamp,
    sunset     timestamp
);

-- +goose Down
drop table if exists prh.daily_data;
drop table if exists prh.weather;
drop table if exists prh.thermostat;
