CREATE TABLE IF NOT EXISTS business_entity (
    buty_id SERIAL,
    buty_name VARCHAR(15),
    buty_code VARCHAR(5),
    buty_enty_id int,
    constraint buty_id_pk primary key(buty_id),
    constraint buty_name_uq unique(buty_name),
    constraint buty_code_uq unique(buty_code),
    constraint buty_enty_id_fk foreign key (buty_enty_id) references entity_type(entity_id)
);