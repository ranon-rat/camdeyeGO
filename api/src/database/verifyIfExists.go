package database

func VerifyIfExists(token string) (id int, err error) {
	db := connect()
	defer db.Close()
	db.QueryRow("SELECT id FROM camera WHERE token = $1", EncryptSha256(token)).Scan(&id)
	return
}
