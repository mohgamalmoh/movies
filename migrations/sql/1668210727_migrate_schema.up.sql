BEGIN;

create table movies
(
    id       serial
        primary key,
    name     varchar(200),
    genre    varchar(255),
    year     varchar(255),
    overview text
);

create table movies_sync_status
(
    id   integer,
    name varchar
);

create table users
(
    id    bigserial
    primary key,
    name  varchar(255),
    email varchar(255)
    );

create table users_movies
(
    id     bigserial
        primary key,
    "user" bigint,
    movie  bigint
);

COMMIT;