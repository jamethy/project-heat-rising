-- +goose Up
create table if not exists prh.daily_data
(
    id              serial primary key,
    created_at      timestamp not null default now(),
    date            date      not null,
    sunrise         timestamp not null,
    sunset          timestamp not null,
    summary_date    timestamp not null,
    bed_time_temp   float     not null,
    fan_on          text      not null,
    temperature_max float     not null,
    temperature_avg float     not null,
    temperature_sum float     not null,
    feels_like_max  float     not null,
    feels_like_avg  float     not null,
    feels_like_sum  float     not null,
    uv_max          float     not null,
    uv_avg          float     not null,
    uv_sum          float     not null,
    rain_max        float     not null,
    rain_avg        float     not null,
    rain_sum        float     not null,
    cloud_max       float     not null,
    cloud_avg       float     not null,
    cloud_sum       float     not null
);


-- +goose Down
drop table if exists prh.daily_data;
