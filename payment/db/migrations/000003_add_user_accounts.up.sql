CREATE TABLE IF NOT EXISTS user_accounts(
    usac_id serial,
    usac_account_no varchar(30),
    usac_balance decimal(18,2),
    usac_created_on date,
    usac_buty_id int,
    usac_user_id int,
    constraint usac_id_pk primary key(usac_id),
    constraint usac_account_no_uq unique (usac_account_no),
    constraint usac_buty_id_fk foreign key(usac_buty_id) references business_entity(buty_id),
    constraint usac_user_id_fk foreign key(usac_user_id) references users(user_id)
);