package main

type Cursor struct {
	table      *Table
	rowNum     int
	endOfTable bool
}

func NewCursor(table *Table) *Cursor {

	return &Cursor{
		table:      table,
		rowNum:     0,
		endOfTable: false,
	}
}

func (c *Cursor) Advance() {
	c.rowNum += 1
	if c.rowNum >= len(c.table.Rows) {
		c.endOfTable = true
	}
}
