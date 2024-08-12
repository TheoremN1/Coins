create table users
(
    Id uuid primary key,
    Name text not null,
    Surname text not null,
    Email text not null unique,
    Password text not null
)