package main

type Row struct {
	ID       string
	Username string
	Email    string
}

type Table struct {
	Rows []Row
}

func NewTable() *Table {
	return &Table{
		Rows: make([]Row, 0),
	}
}
