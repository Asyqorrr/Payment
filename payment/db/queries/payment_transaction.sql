-- name: GetAllTransactionInfo :many
SELECT patrx_no, patrx_created_on, patrx_debet, patrx_credit, patrx_notes, 
	patrx_acctno_from, patrx_acctno_to, t.traty_name, ua.usac_account_no, usac_balance
	FROM public.payment_transactions
	JOIN transaction_type as t ON t.traty_id = patrx_traty_id
	JOIN user_accounts as ua ON ua.usac_account_no = patrx_acctno_from;

-- name: GetTransactionFrom :many
SELECT patrx_no, patrx_created_on, patrx_debet, patrx_credit, patrx_notes, 
	patrx_acctno_from, patrx_acctno_to, t.traty_name, ua.usac_account_no, usac_balance
	FROM public.payment_transactions
	JOIN transaction_type as t ON t.traty_id = patrx_traty_id
	JOIN user_accounts as ua ON ua.usac_account_no = patrx_acctno_from
    WHERE patrx_acctno_from = $1;

-- name: GetTransactionTo :many
SELECT patrx_no, patrx_created_on, patrx_debet, patrx_credit, patrx_notes, 
	patrx_acctno_from, patrx_acctno_to, t.traty_name, ua.usac_account_no, usac_balance
	FROM public.payment_transactions
	JOIN transaction_type as t ON t.traty_id = patrx_traty_id
	JOIN user_accounts as ua ON ua.usac_account_no = patrx_acctno_to
    WHERE patrx_acctno_to = $1;

-- name: SetorTransaction :one
INSERT INTO payment_transactions(
	patrx_no, patrx_created_on, patrx_debet, patrx_acctno_from, patrx_acctno_to, patrx_traty_id)
	VALUES ( $1, $2, $3, $4, $5, 1)
    RETURNING *;

-- name: TransferSend :one
INSERT INTO payment_transactions(
	patrx_no, patrx_created_on, patrx_credit, patrx_notes, patrx_acctno_to, patrx_traty_id)
	VALUES ($1, $2, $3, $4, $5, $6)
    RETURNING *;

-- name: TransferReceive :one
INSERT INTO payment_transactions(
	patrx_no, patrx_created_on, patrx_debet, patrx_notes, patrx_acctno_from, patrx_traty_id, patrx_patrx_ref)
	VALUES ($1, $2, $3, $4, $5, $6, $7) 
    RETURNING *;

-- name: GetTransactionFromAndTo :many
SELECT patrx_no, patrx_created_on, patrx_debet, patrx_credit, patrx_notes, patrx_acctno_from, patrx_acctno_to, patrx_patrx_ref, patrx_traty_id
	FROM payment_transactions
WHERE patrx_acctno_from = $1 or patrx_acctno_to = $2;

-- name: Transfer :one
INSERT INTO payment_transactions(
	patrx_no, patrx_created_on, patrx_debet, patrx_credit, patrx_notes, patrx_acctno_from, patrx_acctno_to, patrx_patrx_ref, patrx_traty_id)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	RETURNING *;