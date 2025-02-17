CREATE TABLE users
(
    id serial primary key,
    name varchar(256) not null,
    username varchar(256) not null unique,
    password_hash text not null
);

CREATE TABLE todo_lists 
(
    id serial primary key,
    title varchar(256) not null,
    description text,
    tags text[],
    done bool not null default false,
    created_at timestamp default now(),
    updated_at timestamp default now()
);

CREATE TABLE users_lists
(
    user_id int not null,
    list_id int not null,
    primary key (user_id, list_id),
    foreign key (user_id) references users(id) on delete cascade,
    foreign key (list_id) references todo_lists(id) on delete cascade
);