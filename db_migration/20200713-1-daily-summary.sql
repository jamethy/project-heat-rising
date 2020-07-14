alter table prh.daily_data
    add column summary_date    timestamp,
    add column bed_time_temp   float(8),
    add column fan_on          text,
    add column temperature_max float(8),
    add column temperature_avg float(8),
    add column temperature_sum float(8),
    add column feels_like_max  float(8),
    add column feels_like_avg  float(8),
    add column feels_like_sum  float(8),
    add column uv_max          float(8),
    add column uv_avg          float(8),
    add column uv_sum          float(8),
    add column rain_max        float(8),
    add column rain_avg        float(8),
    add column rain_sum        float(8),
    add column cloud_max       float(8),
    add column cloud_avg       float(8),
    add column cloud_sum       float(8);


grant update on prh.daily_data to prhapp;
