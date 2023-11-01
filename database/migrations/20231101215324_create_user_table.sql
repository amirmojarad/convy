-- +goose Up
create table users(
    id serial primary key,
    username varchar(255) not null unique,
    email varchar(255) not null unique,
    hashed_password varchar(255) not null,
    first_name varchar(255) default '',
    last_name varchar(255) default '',
    last_login timestamp default null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    deleted_at timestamp default null
);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
