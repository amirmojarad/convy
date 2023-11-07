-- +goose Up
create table if not exists pr_chat_messages(
    id serial primary key,
    sender_id int not null,
    receiver_id int not null,
    message text not null,
    private_chat_id int not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    deleted_at timestamp default null,
    constraint fk_private_chat_id foreign key (private_chat_id) references private_chat(id) on delete cascade,
    constraint fk_sender_id foreign key (sender_id) references users(id) on delete no action,
    constraint fk_sender_id foreign key (receiver_id) references users(id) on delete no action
);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
