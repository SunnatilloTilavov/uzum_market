CREATE TABLE IF NOT EXISTS auth (
    gmail varchar UNIQUE,
    password varchar
);

ALTER TABLE auth ADD COLUMN created_at timestamp default NOW();
ALTER TABLE auth ADD COLUMN updated_at timestamp;
ALTER TABLE auth ADD COLUMN deleted_at timestamp;
ALTER TABLE auth RENAME TO customer;

CREATE TABLE IF NOT EXISTS seller(
    gmail varchar UNIQUE,
    password varchar,
    created_at timestamp default NOW(),
    updated_at timestamp,
    deleted_at timestamp
);

CREATE TABLE IF NOT EXISTS system_user(
    gmail varchar UNIQUE,
    password varchar,
    created_at timestamp default NOW(),
    updated_at timestamp,
    deleted_at timestamp
);