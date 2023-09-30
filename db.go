package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type DBConnection struct {
	Success bool   `json:"success"`
	Err     string `json:"err"`
}

type ColumnType struct {
	DBType string `json:"type"`
}

type TableData struct {
	Columns     []string     `json:"columns"`
	ColumnTypes []ColumnType `json:"columnTypes"`
	Data        [][]string   `json:"data"`
}

func ShowTables(db *sql.DB) ([]string, error) {
	rows, err := db.Query("SHOW TABLES")
	if err != nil {
		fmt.Println(err.Error())
	}

	cols, err := rows.Columns()
	if err != nil {
		fmt.Println(err.Error())
	}

	rawResult := make([][]byte, len(cols))
	var results []string

	dest := make([]interface{}, len(cols))
	for i := range rawResult {
		dest[i] = &rawResult[i]
	}

	for rows.Next() {
		err = rows.Scan(dest...)
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}

		for _, raw := range rawResult {
			if raw == nil {
				results = append(results, "\\N")
			} else {
				results = append(results, string(raw))
			}
		}
	}
	return results, nil
}

func SelectAll(tableName string, db *sql.DB) (TableData, error) {
	var tableData TableData
	rows, err := db.Query(fmt.Sprintf("SELECT * FROM %s LIMIT 100", tableName))
	if err != nil {
		fmt.Println(err.Error())
		return tableData, err
	}
	cols, err := rows.Columns()
	if err != nil {
		fmt.Println(err.Error())
		return tableData, err
	}

	columnTypes, _ := rows.ColumnTypes()
	tableData.Columns = cols
	for _, c := range columnTypes {
		tableData.ColumnTypes = append(tableData.ColumnTypes, ColumnType{DBType: (*c).DatabaseTypeName()})
	}

	rawResult := make([][]byte, len(cols))
	result := make([]string, len(cols))
	var results [][]string

	dest := make([]interface{}, len(cols))
	for i := range rawResult {
		dest[i] = &rawResult[i]
	}

	for rows.Next() {
		err = rows.Scan(dest...)
		if err != nil {
			fmt.Println("Failed to scan row", err)
			return tableData, err
		}

		for i, raw := range rawResult {
			if raw == nil {
				result[i] = ""
			} else {
				result[i] = string(raw)
			}
		}
		results = append(results, result)
	}
	tableData.Data = results
	return tableData, nil
}
