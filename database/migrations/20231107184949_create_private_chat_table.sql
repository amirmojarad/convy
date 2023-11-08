-- +goose Up
create table if not exists private_chat(
    id serial primary key,
    first_user int not null,
    second_user int not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    deleted_at timestamp default null,
    constraint fk_first_user foreign key (first_user) references users(id) on delete no action,
    constraint fk_second_user foreign key (second_user) references users(id) on delete no action
);

-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
drop table private_chat;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
