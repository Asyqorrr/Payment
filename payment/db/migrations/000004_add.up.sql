CREATE TABLE IF NOT EXISTS users(
    user_id serial,
    user_password varchar(125),
    user_name varchar(35),
    user_email varchar(35),
    user_handphone varchar(15),
    user_created datetime,
    constraint user_id_pk primary key(user_id),
    constraint user_name_uq unique(user_name),
    constraint user_email_uq unique(user_email),
    constraint user_handphone_uq unique(user_handphone)
);