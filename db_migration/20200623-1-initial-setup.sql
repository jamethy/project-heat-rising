create schema if not exists prh;

---- create app user separately first
-- create schema if not exists prh;
-- create user prhapp with password '<password here>';
-- grant connect on database projectrisingheat to prhapp;
-- grant usage on schema prh to prhapp;

-- alter default privileges in schema prh grant select on tables to prhapp;
-- alter default privileges in schema prh grant insert on tables to prhapp;

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
    id            serial primary key not null,
    created_at    timestamptz        not null default now(),
    timestamp     timestamp          not null,
    Provider      text,
    Temperature   float(8),
    FeelsLike     float(8),
    Pressure      float(8),
    Humidity      float(8),
    WindSpeed     float(8),
    WindDirection float(8),
    Clouds        float(8)
);

create index weather_timestamp on prh.weather (timestamp);

-- daily data
create table prh.daily_data
(
    id         serial primary key not null,
    created_at timestamptz        not null default now(),
    date       date               not null,
    Sunrise    timestamptz,
    Sunset     timestamptz
);

create index daily_data_timestamp on prh.daily_data (date);
