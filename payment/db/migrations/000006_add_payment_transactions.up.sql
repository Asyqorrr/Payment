CREATE TABLE IF NOT EXISTS payment_transactions(
    patrx_no varchar(55),
    patrx_created_on date,
    patrx_debet decimal(18,2),
    patrx_credit decimal(18,2),
    patrx_notes varchar(125),
    patrx_acctno_from varchar(30),
    patrx_acctno_to varchar(30),
    patrx_patrx_ref varchar(55),
    patrx_traty_id int,
    constraint patrx_no_pk primary key (patrx_no),
	constraint patrx_no_uq unique (patrx_no),
    constraint patrx_acctno_from_fk foreign key (patrx_acctno_from) references user_accounts(usac_account_no),
    constraint patrx_acctno_to_fk foreign key (patrx_acctno_to) references user_accounts(usac_account_no),
    constraint patrx_traty_id_fk foreign key (patrx_traty_id) references transaction_type(traty_id)
);