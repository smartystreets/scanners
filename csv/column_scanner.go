package csv

import (
	"fmt"
	"io"
	"log"
)

type ColumnScanner struct {
	*Scanner
	headerRecord []string
	columnIndex  map[string]int
}

func NewColumnScanner(reader io.Reader, options ...Option) (*ColumnScanner, error) {
	inner := NewScanner(reader, append(options, Options.FieldsPerRecord(0))...)
	if !inner.Scan() {
		return nil, inner.Error()
	}
	scanner := &ColumnScanner{
		Scanner:      inner,
		headerRecord: inner.Record(),
		columnIndex:  make(map[string]int),
	}
	scanner.readHeader()
	return scanner, nil
}

func (this *ColumnScanner) readHeader() {
	for i, value := range this.headerRecord {
		this.columnIndex[value] = i
	}
}

func (this *ColumnScanner) Header() []string {
	return this.headerRecord
}

func (this *ColumnScanner) ColumnErr(column string) (string, error) {
	index, ok := this.columnIndex[column]
	if !ok {
		return "", fmt.Errorf("Column [%s] not present in header record: %#v\n", column, this.headerRecord)
	}
	return this.Record()[index], nil
}

func (this *ColumnScanner) Column(column string) string {
	value, err := this.ColumnErr(column)
	if err != nil {
		log.Panic(err)
	}
	return value
}
