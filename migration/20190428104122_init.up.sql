create table Teacher
(
    id   varchar(16) not null,
    name varchar(128)
);

create unique index Teacher_id_uindex
    on Teacher (id);

alter table Teacher
    add constraint Teacher_pk
        primary key (id);

create table Token
(
    id        bigserial    not null,
    tokenHash varchar(256) not null
);

create unique index Token_id_uindex
    on Token (id);

create unique index Token_tokenHash_uindex
    on Token (tokenHash);

alter table Token
    add constraint Token_pk
        primary key (id);
