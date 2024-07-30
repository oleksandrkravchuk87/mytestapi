package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// .\migrate --dbUser=root --dbPass=root --dbHost=localhost --dbPort=3306 --sqlFile=./migrations.sql
func main() {
	dbHost := flag.String("dbHost", "localhost", "database host")
	dbPort := flag.String("dbPort", "3306", "database port")
	dbUser := flag.String("dbUser", "", "database user")
	dbPass := flag.String("dbPass", "", "database password")
	sqlFile := flag.String("sqlFile", "", "SQL script file")
	flag.Parse()

	if *dbUser == "" || *dbPass == "" || *sqlFile == "" {
		log.Fatal("dbUser, dbPass, and sqlFile flags are required")
	}

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/", *dbUser, *dbPass, *dbHost, *dbPort))
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}

	sqlBytes, err := os.ReadFile(*sqlFile)
	if err != nil {
		log.Fatalf("Error reading SQL file: %v", err)
	}
	sqlScript := string(sqlBytes)

	sqlStatements := strings.Split(sqlScript, ";")

	for _, stmt := range sqlStatements {
		stmt = strings.TrimSpace(stmt)
		if stmt == "" {
			continue
		}
		_, err := db.Exec(stmt)
		if err != nil {
			log.Fatalf("Error executing statement (%s): %v", stmt, err)
		}
	}

	fmt.Println("SQL script executed successfully")
}
