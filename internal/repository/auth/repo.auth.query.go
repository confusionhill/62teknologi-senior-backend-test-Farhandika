package auth

const (
	queryInsertNewUser      = "INSERT INTO users(email, password, full_name) VALUES ($1,$2,$3) RETURNING id, access;"
	queryGetUserInfoByEmail = "SELECT id, email, password, full_name, access FROM users WHERE email=$1 LIMIT 1;"
)
