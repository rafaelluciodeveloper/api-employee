package models

import "go-api/db"

func GetAll() (employees []Employee, err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}

	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM employees`)
	if err != nil {
		return
	}

	for rows.Next() {
		var employee Employee

		err = rows.Scan(&employee.ID, &employee.Name, &employee.Occupation, &employee.Salary)
		if err != nil {
			continue
		}

		employees = append(employees, employee)
	}

	return
}
