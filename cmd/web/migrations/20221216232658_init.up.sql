CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    id uuid primary key,
    username varchar unique not null
);

--bun:split

CREATE TABLE tramites (
    id uuid primary key,
    ano_presupuestario int not null,
    link varchar not null,
    created_at timestamp not null default now(),
    updated_at timestamp not null default now(),
    categoria varchar not null,
    version int not null default 0,
    estado varchar not null default 'borrador',
    user_id uuid references users(id)
);

CREATE TABLE events (
    id uuid primary key default uuid_generate_v4(),
    kind varchar not null,
    payload jsonb not null,
    tramite_id uuid references tramites(id),
    created_at timestamp not null default now(),
    updated_at timestamp not null default now()
);

CREATE TABLE observations (
    id serial primary key,
    content varchar not null,
    tramite_id uuid references tramites(id)
);
