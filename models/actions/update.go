package actions

import (
	"go-api/db"
	"go-api/models"
)

func Update(id string, employe models.Employee) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	res, err := conn.Exec(`UPDATE employees  SET name=$1,occupation=$2,salary=$2 WHERE id=$5`, employe.Name, employe.Occupation, employe.Salary, id)

	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}
