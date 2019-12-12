package main

import (
	"fmt"

	_ "github.com/lib/pq"
)

func getRiver(riverID int) (River, error) {
	//Retrieve
	res := River{}

	var id int
	var name string
	var city string
	var level int
	var date string

	err := db.QueryRow(`SELECT id, name, city, level, date FROM rivers where id = $1`, riverID).Scan(&id, &name, &city, &level, &date)
	if err == nil {
		res = River{ID: id, Name: name, City: city, Level: level, Date: date}
	}

	return res, err
}

func allRivers() ([]River, error) {
	//Retrieve
	rivers := []River{}

	rows, err := db.Query(`SELECT id, name, city, level, date FROM rivers order by id`)
	defer rows.Close()
	if err == nil {
		for rows.Next() {
			var id int
			var name string
			var city string
			var level int
			var date string

			err = rows.Scan(&id, &name, &city, &level, &date)
			if err == nil {
				currentRiver := River{ID: id, Name: name, City: city, Level: level, Date: date}
				rivers = append(rivers, currentRiver)
			} else {
				return rivers, err
			}
		}
	} else {
		return rivers, err
	}

	return rivers, err
}

func insertRiver(name, city string, level int, date string) (int, error) {
	//Create
	var riverID int
	err := db.QueryRow(`INSERT INTO rivers(name, city, level, date) VALUES($1, $2, $3, $4) RETURNING id`, name, city, level, date).Scan(&riverID)

	if err != nil {
		return 0, err
	}

	fmt.Printf("Last inserted ID: %v\n", riverID)
	return riverID, err
}

func updateRiver(id int, name, city string, level int, date string) (int, error) {
	//Create
	res, err := db.Exec(`UPDATE rivers set name=$1, city=$2, level=$3, date=$4 where id=$5 RETURNING id`, name, city, level, date, id)
	if err != nil {
		return 0, err
	}

	rowsUpdated, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(rowsUpdated), err
}

func removeRiver(riverID int) (int, error) {
	//Delete
	res, err := db.Exec(`delete from rivers where id = $1`, riverID)
	if err != nil {
		return 0, err
	}

	rowsDeleted, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(rowsDeleted), nil
}
