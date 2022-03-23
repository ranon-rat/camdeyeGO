package database

func GetCamerasDomain() (domains []string) {

	db := openDatabase()
	defer db.Close()
	rows, err := db.Query("SELECT domain FROM cameras")
	if err != nil {
		return
	}
	for rows.Next() {
		domain := ""
		rows.Scan(&domain)
		domains = append(domains, domain)
	}
	return
}
