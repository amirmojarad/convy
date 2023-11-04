-- +goose Up
create table if not exists tokens (
    id serial primary key,
    user_id int not null unique,
    hashed_tokens varchar(255),
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    deleted_at timestamp default null,
    constraint fk_tokens foreign key (user_id) references users(id) on delete cascade
);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
