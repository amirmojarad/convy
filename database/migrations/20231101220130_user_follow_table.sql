-- +goose Up
create table user_follow (
    id serial primary key,
    follower integer not null,
    following integer not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    deleted_at timestamp default null,
    constraint fk_follower foreign key (follower) references users(id) on delete cascade,
    constraint fk_following foreign key (following) references users(id) on delete cascade
);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
drop table user_follow;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
