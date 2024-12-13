package test

import (
	"database/sql"
	"fmt"
	"os"
	"testing"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func SeedDb(db *sql.DB, t *testing.T) error {
	//Ask Nick for better Path initilization
	data, err := os.ReadFile("C:/Users/User/Desktop/Projects/RealWorld/realworld-golang/app/test/seed.sql")

	// Is recovering good?
	if err != nil {
		panic("Cannot read seed.sql file!")
	}

	rows, err := db.Query(string(data))

	fmt.Println(rows, err)
	if err != nil {
		return err
	}

	for rows.Next() {
		var usr int
		if err := rows.Scan(&usr); err != nil {
			return err
		}
		fmt.Println("heeheh")

		fmt.Println(usr)
	}
	if err = rows.Err(); err != nil {
		return err
	}

	defer rows.Close()

	return nil
}
