package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
)

type Row struct {
	ID       string
	Username string
	Email    string
}

func (r Row) Serialize() *bytes.Buffer {

	var buff bytes.Buffer
	rowByte := []byte(fmt.Sprintf("%s||%s||%s", r.ID, r.Username, r.Email))
	binary.Write(&buff, binary.LittleEndian, rowByte)
	return &buff
}

type Table struct {
	Rows []Row
}

func NewTable() *Table {
	return &Table{
		Rows: make([]Row, 0),
	}
}

func TableStart(table *Table) *Cursor {
	cursor := NewCursor(table)
	cursor.rowNum = 0

	if len(table.Rows) == 0 {
		cursor.endOfTable = true
	}

	return cursor
}

func TableEnd(table *Table) *Cursor {
	cursor := NewCursor(table)
	cursor.rowNum = len(table.Rows) - 1
	cursor.endOfTable = true

	return cursor
}

func (t *Table) SaveTableToDisk(databaseFile *os.File) {

}
