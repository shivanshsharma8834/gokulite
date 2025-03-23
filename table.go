package main

import (
	"bytes"
	"encoding/binary"
	"os"
)

type Row struct {
	ID       string
	Username string
	Email    string
}

func (r Row) Serialize() ([]byte, error) {

	buf := new(bytes.Buffer)

	IDBytes := []byte(r.ID)
	if err := binary.Write(buf, binary.LittleEndian, uint16(len(IDBytes))); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.LittleEndian, IDBytes); err != nil {
		return nil, err
	}

	UserNameBytes := []byte(r.Username)
	if err := binary.Write(buf, binary.LittleEndian, uint16(len(UserNameBytes))); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.LittleEndian, UserNameBytes); err != nil {
		return nil, err
	}

	EmailBytes := []byte(r.Email)
	if err := binary.Write(buf, binary.LittleEndian, uint16(len(EmailBytes))); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.LittleEndian, EmailBytes); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type Table struct {
	Rows []Row
}

func (t *Table) SerializeAllRows() []byte {
	buf := new(bytes.Buffer)

	for _, row := range t.Rows {
		rowBytes, _ := row.Serialize()
		buf.Write(rowBytes)
	}

	return buf.Bytes()
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
