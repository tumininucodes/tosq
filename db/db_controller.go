package db

import (
	"database/sql"
	f "fmt"
	_ "github.com/go-sql-driver/mysql"
)

func OpenDB() *sql.DB {

	db, error := sql.Open("mysql", "root:alade2001@tcp(localhost:3306)/testdb")
	if error != nil {
		f.Println("error validating sql.Open arguments")
		panic(error.Error())
	} else {
		f.Println("Successfully opened")
	}


	// defer db.Close()

	error = db.Ping()

	if error != nil {
		f.Println("error verifying connection with db.Ping") 
		panic(error.Error())
	} else {
		f.Println("db alive")
	}

	return db
}


func GetTodos(db *sql.DB) *sql.Rows {
	rows, err := db.Query("SELECT * FROM students")
	if err != nil {
		f.Println("Error executing query:", err)
		// return
	}
	defer rows.Close()

	// Iterate over the rows and retrieve data
	for rows.Next() {
		var column1 int
		var column2 string
		var column3 string

		// Scan the values from the current row into variables
		err := rows.Scan(&column1, &column2, &column3)
		if err != nil {
			f.Println("Error retrieving data:", err)
			// return
		}
		// Process the retrieved data
		f.Println("Column1:", column1, "Column2:", column2, "Column3:", column3)
	}
	if err := rows.Err(); err != nil {
		f.Println("Error retrieving data:", err)
		// return
	}

	return rows
}


// insert, err := db.Query("INSERT INTO `testdb`.`students` (`id`, `firstname`, `lastname`) VALUES ('2', 'Ben', 'Ford');")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer insert.Close()
	// f.Println("Successful connection to database")

	
