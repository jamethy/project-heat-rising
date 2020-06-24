create schema if not exists prh;

---- create app user separately first
-- create schema if not exists prh;
-- create user prhapp with password '<password here>';
-- create user prhreadonly with password '<password here>';
-- grant connect on database projectrisingheat to prhapp,prhreadonly;
-- grant usage on schema prh to prhapp,prhreadonly;

-- alter default privileges in schema prh grant select on tables to prhapp,prhreadonly;
-- alter default privileges in schema prh grant insert on tables to prhapp;
-- alter default privileges in schema prh grant usage,select on sequences to prhapp;

-- thermostat table
create table prh.thermostat
(
    id            serial primary key not null,
    created_at    timestamptz        not null default now(),
    timestamp     timestamp          not null,
    provider      text,
    thermostat_id text,
    target_cool   float(8),
    target_heat   float(8),
    actual_temp   float(8),
    humidity      float(8),
    is_heating    bool,
    is_cooling    bool
);

create index thermostat_timestamp on prh.thermostat (timestamp);

-- weather table
create table prh.weather
(
    id             serial primary key not null,
    created_at     timestamptz        not null default now(),
    timestamp      timestamp          not null,
    provider       text,
    temperature    float(8),
    feels_like     float(8),
    pressure       float(8),
    humidity       float(8),
    wind_speed     float(8),
    wind_direction float(8),
    clouds         float(8)
);

create index weather_timestamp on prh.weather (timestamp);

-- daily data
create table prh.daily_data
(
    id         serial primary key not null,
    created_at timestamptz        not null default now(),
    date       date               not null,
    sunrise    timestamptz,
    sunset     timestamptz
);

create index daily_data_timestamp on prh.daily_data (date);
