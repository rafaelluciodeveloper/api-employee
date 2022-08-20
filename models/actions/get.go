package actions

import (
	"go-api/db"
	"go-api/models"
)

func Get(id int64) (employee models.Employee, err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}

	defer conn.Close()

	row := conn.QueryRow(`SELECT * FROM employees WHERE id=$1`, id)

	err = row.Scan(&employee.ID, &employee.Name, &employee.Occupation, &employee.Salary)

	return
}
