-- name: GetAllUserAccount :many
SELECT ua.usac_id, usac_account_no, usac_balance, usac_created_on, u.user_name, be.buty_name
FROM user_accounts as ua
join users as u on ua.usac_user_id = u.user_id
join business_entity as be on ua.usac_buty_id = be.buty_id;

-- name: FindUserAccountByAccno :one
SELECT ua.usac_id, usac_account_no, usac_balance, usac_created_on, u.user_name, be.buty_name
FROM user_accounts as ua
join users as u on ua.usac_user_id = u.user_id
join business_entity as be on ua.usac_buty_id = be.buty_id
where ua.usac_account_no = $1;

-- name: CreateUserAccount :one
INSERT INTO user_accounts(
	usac_account_no, usac_balance, usac_created_on, usac_buty_id, usac_user_id)
	VALUES (
	$1, $2, $3, $4, $5)
    RETURNING *;

-- name: DeleteUserAccount :exec
DELETE FROM user_accounts
    WHERE usac_user_id = $1
    RETURNING *;

-- name: UpdateBalance :one
UPDATE user_accounts
	SET  usac_balance=$2
	WHERE usac_account_no=$1
	RETURNING *;