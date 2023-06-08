-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

create table if not exists users
(
    id   serial primary key,
    name varchar(20) unique not null,
    description varchar(255),
    email bytea unique not null,
    password bytea not null,
    photo bytea
);

create table if not exists musics
(
    id serial primary key,
    authorID int not null,
    title varchar(20) not null,
    author varchar(20) not null,
    music bytea not null,
    photo bytea not null,
    constraint authorID foreign key(authorID) references users(id),
    constraint author foreign key(author) references users(name)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
drop table events;
drop table users;
-- +goose StatementEnd
