package actions

import (
	"go-api/db"
	"go-api/models"
)

func Insert(employee models.Employee) (id int64, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	sql := `INSERT INTO employees (name, occupation, salary) VALUES ($1,$2,$3) RETURNING id`

	err = conn.QueryRow(sql, employee.Name, employee.Occupation, employee.Salary).Scan(&id)

	return
}
