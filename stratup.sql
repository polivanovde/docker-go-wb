create database pgdb;
create user uzer;
grant all on database pgdb to uzer;
create table if not exists messages
(
    order_uid string not null,
    message   jsonb
);
create unique index messages_order_uid_uindex
    on messages (order_uid);
-- для nats-str