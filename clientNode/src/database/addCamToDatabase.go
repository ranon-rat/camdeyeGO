package database

import "time"

func AddURL(url string) (err error) {
	db := openDatabase()
	defer db.Close()
	// in case that the domain is repeat
	_, err = db.Exec("insert into cameras values (domain,  last_time_up ) value(?1,?2)", url, time.Now().Unix())
	return
}
