-- +goose Up
alter table prh.weather add column uv_index float;
update prh.weather set uv_index = -1 where true;
alter table prh.weather alter column uv_index set not null;

alter table prh.weather add column rain_level float;
update prh.weather set rain_level = -1 where true;
alter table prh.weather alter column rain_level set not null;

alter table prh.weather add column weather_description text;
update prh.weather set weather_description = '' where true;
alter table prh.weather alter column weather_description set not null;

-- +goose Down
alter table prh.weather drop column uv_index;
alter table prh.weather drop column rain_level;
alter table prh.weather drop column weather_description;
