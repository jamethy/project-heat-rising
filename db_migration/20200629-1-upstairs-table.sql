-- upstairs table
create table prh.upstairs
(
    id             serial primary key not null,
    created_at     timestamptz        not null default now(),
    timestamp      timestamp          not null,
    provider       text,
    temperature    float(8),
    pressure       float(8),
    humidity       float(8)
);

create index upstairs_timestamp on prh.upstairs (timestamp);
