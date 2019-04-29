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