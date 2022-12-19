CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    id uuid primary key,
    username varchar unique not null
);

--bun:split

CREATE TABLE tramites (
    id uuid primary key,
    anoPresupuestario int not null,
    link varchar not null,
    created_at timestamp not null default now(),
    updated_at timestamp not null default now(),
    categoria varchar not null,
    user_id uuid references users(id)
);
