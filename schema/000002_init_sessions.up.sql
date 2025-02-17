CREATE TABLE sessions
(
    id serial primary key,
    user_id int not null,
    created_at timestamp not null default now(),
    expires_at timestamp not null,
    foreign key (user_id) references users(id) on delete cascade
);