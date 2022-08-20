package models

import "go-api/db"

func Get(id int64) (employe Employee, err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}

	defer conn.Close()

	row := conn.QueryRow(`SELECT * FROM employees WHERE id=$id`, id)

	err = row.Scan(&employe.ID, &employe.Name, &employe.Occupation, &employe.Salary)

	return
}
