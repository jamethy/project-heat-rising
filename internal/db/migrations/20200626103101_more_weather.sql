-- +goose Up
alter table prh.weather add column uv_index float(8);
alter table prh.weather add column rain_level float(8);
alter table prh.weather add column weather_description text;

-- +goose Down
alter table prh.weather drop column uv_index;
alter table prh.weather drop column rain_level;
alter table prh.weather drop column weather_description;
