-- +migrate Up
create table urls (
    hash text primary key
    original_url text not null
    created_at timestamp default now()
);

-- +migrate Down
drop table urls