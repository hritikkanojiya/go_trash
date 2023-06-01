package main

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	username := "gm_dev"
	password := "GM@1302$#"
	host := "3.110.4.157"
	port := "5432"
	database := "gm_dev"

	escapedPassword := url.QueryEscape(password)

	connectionURL := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", username, escapedPassword, host, port, database)

	pool, err := pgxpool.Connect(context.Background(), connectionURL)
	if err != nil {
		panic(err.Error())
	}
	defer pool.Close()

	allUsersResult, err := fetchRecords(pool, "get_all_users")
	if err != nil {
		panic(err.Error())
	}

	for _, row := range allUsersResult {
		fmt.Println("Row:", row)
	}

}

func fetchRecords(pool *pgxpool.Pool, procName string, params ...interface{}) ([]map[string]interface{}, error) {
	stmt := fmt.Sprintf("SELECT * FROM %s(", procName)
	var placeholders []string
	for i := 1; i <= len(params); i++ {
		placeholders = append(placeholders, fmt.Sprintf("$%d", i))
	}

	stmt += strings.Join(placeholders, ",") + ")"

	rows, err := pool.Query(context.Background(), stmt, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columnDescriptions := rows.FieldDescriptions()
	colCount := len(columnDescriptions)

	values := make([]interface{}, colCount)
	for i := range values {
		values[i] = new(interface{})
	}

	var result []map[string]interface{}

	for rows.Next() {
		err := rows.Scan(values...)
		if err != nil {
			return nil, err
		}

		rowData := make(map[string]interface{})
		for i, colDesc := range columnDescriptions {
			colName := string(colDesc.Name)
			colValue := *values[i].(*interface{})
			rowData[colName] = colValue
		}

		result = append(result, rowData)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func createRecord(pool *pgxpool.Pool, procName string, values ...interface{}) (map[string]interface{}, error) {
	stmt := fmt.Sprintf("SELECT * FROM %s($1, $2)", procName)

	rows, err := pool.Query(context.Background(), stmt, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columnDescriptions := rows.FieldDescriptions()
	colCount := len(columnDescriptions)

	columns := make([]string, colCount)
	valuePtrs := make([]interface{}, colCount)

	for i, desc := range columnDescriptions {
		columns[i] = string(desc.Name)
		valuePtrs[i] = new(interface{})
	}

	result := make(map[string]interface{})

	if rows.Next() {
		err := rows.Scan(valuePtrs...)
		if err != nil {
			return nil, err
		}

		for i, column := range columns {
			result[column] = *valuePtrs[i].(*interface{})
		}
	} else {
		return nil, fmt.Errorf("no result found")
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return result, nil
}
