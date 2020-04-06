CREATE OR REPLACE FUNCTION next_id(OUT result bigint, seq text) AS $$
DECLARE
    our_epoch bigint := 1314220021721;
    seq_id bigint;
    now_millis bigint;
    shard_id int := 5;
BEGIN
    SELECT nextval(seq) % 1024 INTO seq_id;
    SELECT FLOOR(EXTRACT(EPOCH FROM clock_timestamp())) INTO now_millis;
    result := (now_millis - our_epoch)*1000 << 23;
    result := result | (shard_id <<10);
    result := result | (seq_id);
    
END;
    $$ LANGUAGE PLPGSQL;


CREATE OR REPLACE FUNCTION convertTVkdau (x text) RETURNS text AS
$$
DECLARE
 cdau text; kdau text; r text;
BEGIN
 cdau = 'áàảãạâấầẩẫậăắằẳẵặđéèẻẽẹêếềểễệíìỉĩịóòỏõọôốồổỗộơớờởỡợúùủũụưứừửữựýỳỷỹỵÁÀẢÃẠÂẤẦẨẪẬĂẮẰẲẴẶĐÉÈẺẼẸÊẾỀỂỄỆÍÌỈĨỊÓÒỎÕỌÔỐỒỔỖỘƠỚỜỞỠỢÚÙỦŨỤƯỨỪỬỮỰÝỲỶỸỴ';
 kdau = 'aaaaaaaaaaaaaaaaadeeeeeeeeeeeiiiiiooooooooooooooooouuuuuuuuuuuyyyyyaaaaaaaaaaaaaaaaadeeeeeeeeeeeiiiiiooooooooooooooooouuuuuuuuuuuyyyyy';
 r = x;
 FOR i IN 0..length(cdau)
 LOOP
 r = replace(r, substr(cdau,i,1), substr(kdau,i,1));
 END LOOP;
 RETURN r;
END;
$$ LANGUAGE plpgsql;
--Generate new func insta5

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
    name text unique
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
    email text,
    phone_number text,
    first_name text,
    last_name text,
    password text,
    username text,
    position int,
    store_id bigint,
    foreign key (store_id) references stores(id)
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
    value text,
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