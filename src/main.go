package main

import (
  "fmt"
  "database/sql"
	"github.com/tealeg/xlsx"
  "./connection"
)

const (
  host     = "localhost"
  port     = 5432
  user     = "postgres"
  password = "qburst"
  dbname   = "my_db"
)

func main() {
  db := dbConnection.ConnectDB()
  defer db.Close()
  fmt.Println("## Writing from Excel to DB ##")
  excelToDB(db)
  fmt.Println("## Writing from DB to Excel ##")
  dbToExcel(db)

}

func excelToDB(db *sql.DB) {
  sqlStatement := `
INSERT INTO users (name, age, email)
VALUES ($1, $2, $3)`

  excelFile := "./xlFiles/sourceXL.xlsx"
    xlFile, err := xlsx.OpenFile(excelFile)
    if err != nil {
        fmt.Println("Error")
    }
    for _, sheet := range xlFile.Sheets {
        for rowIndex, row := range sheet.Rows {
          if rowIndex == 0 {
            continue
          }
          fmt.Println("Inserting row\n")
          currentRow := make([]string, 3)
          for i, cell := range row.Cells {
                text := cell.String()
                currentRow[i] = text
            }
            _, err = db.Exec(sqlStatement, currentRow[0], currentRow[1], currentRow[2])
            if err != nil {
              fmt.Println("Error:", err)
            } else {
              fmt.Println("\n")
            }
        }
    }
}

func dbToExcel(db *sql.DB) {
	sqlStatement := `SELECT * FROM users`

  users, err := db.Query(sqlStatement)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	file := xlsx.NewFile()
  sheet, err := file.AddSheet("sheet1")
  if err != nil {
    fmt.Printf(err.Error())
  }

	for users.Next() {
		details := make([]string, 4)
		err := users.Scan(&details[0], &details[1], &details[2], &details[3])
		if err != nil {
			fmt.Println("Error: ", err)
		}
		row := sheet.AddRow()
		for _, value := range details {
			cell := row.AddCell()
			cell.Value = value
		}
		err = file.Save("./xlFiles/destXL.xlsx")
  	if err != nil {
    	fmt.Printf(err.Error())
  	} else {
			fmt.Println("Row added!")
		}
	}
}
