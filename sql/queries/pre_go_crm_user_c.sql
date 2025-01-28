-- name: GetUserByEmailSQLC :one
select * from `pre_go_crm_user_c` where usr_email = ? limit 1;

-- name: UpdateUserStatusByUserId :exec
update `pre_go_crm_user_c`
set usr_status = $2,
    usr_udpated_at = $3
where usr_id = $1