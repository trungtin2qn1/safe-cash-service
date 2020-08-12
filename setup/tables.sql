CREATE OR REPLACE FUNCTION next_id(OUT result bigint, seq text) AS $$
DECLARE
    our_epoch bigint := 13142200;
    seq_id bigint;
    now_millis bigint;
    shard_id int := 5;
BEGIN
    SELECT nextval(seq) % 1024 INTO seq_id;
    SELECT FLOOR(EXTRACT(EPOCH FROM clock_timestamp())) INTO now_millis;
    result := (now_millis - our_epoch)/1000 << 10;
    result := result | (shard_id <<5);
    result := result | (seq_id);
    
END;
    $$ LANGUAGE PLPGSQL;

--Generate new func insta5

--Create table api_keys
--Drop table api_keys
drop table if exists api_keys;
drop sequence if exists api_keys_id_seq;
CREATE SEQUENCE api_keys_id_seq;

CREATE TABLE api_keys (
    id bigint NOT NULL DEFAULT next_id('api_keys_id_seq') primary key,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp,
    value text unique
);

--Create table admin_users
--Drop table admin_users
drop table if exists admin_users;
drop sequence if exists admin_users_id_seq;
CREATE SEQUENCE admin_users_id_seq;

CREATE TABLE admin_users (
    id bigint NOT NULL DEFAULT next_id('admin_users_id_seq') primary key,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp,
    email text unique,
    first_name text,
    last_name text,
    password text,
    last_login timestamp
);

--Create table work_spaces
--Drop table work_spaces
drop table if exists work_spaces;
drop sequence if exists work_spaces_id_seq;
CREATE SEQUENCE work_spaces_id_seq;

CREATE TABLE work_spaces (
    id bigint NOT NULL DEFAULT next_id('work_spaces_id_seq') primary key,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp,
    name text unique,
    image text
);

--Create table stores
--Drop table stores
drop table if exists stores;
drop sequence if exists stores_id_seq;
CREATE SEQUENCE stores_id_seq;

CREATE TABLE stores (
    id bigint NOT NULL DEFAULT next_id('stores_id_seq') primary key,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp,
    name text unique,
    country text not null,
    state text not null,
    city text not null,
    district text,
    ward text,
    street text,
    image text,
    wp_id bigint,
    foreign key (wp_id) references work_spaces(id)
);

--Create table client_credentials
--Drop table client_credentials
drop table if exists client_credentials;
drop sequence if exists client_credentials_id_seq;
CREATE SEQUENCE client_credentials_id_seq;

CREATE TABLE client_credentials (
    id bigint NOT NULL DEFAULT next_id('client_credentials_id_seq') primary key,
    finger_print text unique,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp,
    store_id bigint,
    foreign key (store_id) references stores(id)
);

--Create table users
--Drop table users
drop table if exists users;
drop sequence if exists users_id_seq;
CREATE SEQUENCE users_id_seq;

CREATE TABLE users (
    id bigint NOT NULL DEFAULT next_id('users_id_seq') primary key,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp,
    email text unique,
    phone_number text,
    first_name text,
    last_name text,
    password text,
    username text,
    avatar text,
    role text,
    position int
);

--Create table store_junction_users
--Drop table store_junction_users
drop table if exists store_junction_users;
drop sequence if exists store_junction_users_id_seq;
CREATE SEQUENCE store_junction_users_id_seq;

CREATE TABLE store_junction_users (
    id bigint NOT NULL DEFAULT next_id('store_junction_users_id_seq') primary key,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp,
    role text,
    user_id bigint,
    store_id bigint,
    foreign key (user_id) references users(id),
    foreign key (store_id) references stores(id)
);

--Create table register_merchants
--Drop table register_merchants
drop table if exists register_merchants;
drop sequence if exists register_merchants_id_seq;
CREATE SEQUENCE register_merchants_id_seq;

CREATE TABLE register_merchants (
    id bigint NOT NULL DEFAULT next_id('register_merchants_id_seq') primary key,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp,
    content text,
    email text,
    name text,
    user_address text,
    store_name text,
    store_address text,
    status smallint,
    user_id bigint,
    foreign key (user_id) references users(id)
);

--Create table notification_tokens
--Drop table notification_tokens
drop table if exists notification_tokens;
drop sequence if exists notification_tokens_id_seq;
CREATE SEQUENCE notification_tokens_id_seq;

CREATE TABLE notification_tokens (
    id bigint NOT NULL DEFAULT next_id('notification_tokens_id_seq') primary key,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp,
    token text unique,
    value text unique,
    user_id bigint,
    foreign key (user_id) references users(id)
);

--Create table medias
--Drop table medias
drop table if exists medias;
drop sequence if exists medias_id_seq;
CREATE SEQUENCE medias_id_seq;

CREATE TABLE medias (
    id bigint NOT NULL DEFAULT next_id('medias_id_seq') primary key,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp,
    link text,
    type int
);

--Create table notifications
--Drop table notifications
drop table if exists notifications;
drop sequence if exists notifications_id_seq;
CREATE SEQUENCE notifications_id_seq;

CREATE TABLE notifications (
    id bigint NOT NULL DEFAULT next_id('notifications_id_seq') primary key,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp,
    title text,
    body text,
    is_read boolean DEFAULT false not null,
    user_id bigint,
    store_id bigint,
    foreign key (store_id) references stores(id),
    foreign key (user_id) references users(id)
);

--Create table unlocking_logs
--Drop table unlocking_logs
drop table if exists unlocking_logs;
drop sequence if exists unlocking_logs_id_seq;
CREATE SEQUENCE unlocking_logs_id_seq;

CREATE TABLE unlocking_logs (
    id bigint NOT NULL DEFAULT next_id('unlocking_logs_id_seq') primary key,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp,
    content text,
    method text,
    is_success boolean,
    user_id bigint,
    store_id bigint,
    foreign key (store_id) references stores(id),
    foreign key (user_id) references users(id)
);

--Create table user_login_logs
--Drop table user_login_logs
drop table if exists user_login_logs;
drop sequence if exists user_login_logs_id_seq;
CREATE SEQUENCE user_login_logs_id_seq;

CREATE TABLE user_login_logs (
    id bigint NOT NULL DEFAULT next_id('user_login_logs_id_seq') primary key,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp,
    user_id bigint,
    foreign key (user_id) references users(id)
);

--Create table store_medias
--Drop table store_medias
drop table if exists store_medias;
drop sequence if exists store_medias_id_seq;
CREATE SEQUENCE store_medias_id_seq;

CREATE TABLE store_medias (
    id bigint NOT NULL DEFAULT next_id('store_medias_id_seq') primary key,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp,
    name text unique,
    type text,
    thumbnail text,
    store_id bigint,
    user_id bigint,
    foreign key (user_id) references users(id),
    foreign key (store_id) references stores(id)
);

--Create table media_unlocking_logs
--Drop table media_unlocking_logs
drop table if exists media_unlocking_logs;
drop sequence if exists media_unlocking_logs_id_seq;
CREATE SEQUENCE media_unlocking_logs_id_seq;

CREATE TABLE media_unlocking_logs (
    id bigint NOT NULL DEFAULT next_id('media_unlocking_logs_id_seq') primary key,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp,
    store_media_id bigint,
    unlocking_log_id bigint,
    foreign key (store_media_id) references store_medias(id),
    foreign key (unlocking_log_id) references unlocking_logs(id)
);