package main

import "fmt"

func executeStatement(statement *Statement, table *Table) error {
	switch statement.Type {
	case "insert":
		table.Rows = append(table.Rows, Row{
			ID:       statement.Data[0],
			Username: statement.Data[1],
			Email:    statement.Data[2],
		})
	case "select":
		if len(table.Rows) == 0 {
			return fmt.Errorf("no rows found")
			break
		}
		for index, row := range table.Rows {
			fmt.Printf("[%d] ID: %s, Username: %s, Email: %s\n",
				index+1,
				row.ID,
				row.Username,
				row.Email,
			)
		}
	}

	return nil
}
