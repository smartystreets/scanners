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

// NewColumnScannerWithoutHeader initializes a scanner using the provided reader and options
// and uses the provided header to build a column index such that columns in each record can
// be fetched using the Column* methods. It requires that all records in the data have the
// same column count as the provided header.
func NewColumnScannerWithoutHeader(reader io.Reader, header []string, options ...Option) *ColumnScanner {
	return &ColumnScanner{
		Scanner:      NewScanner(reader, append(options, FieldsPerRecord(len(header)))...),
		headerRecord: header,
		columnIndex:  deriveColumnIndices(header),
	}
}

// NewColumnScanner initializes a scanner using the provided reader and options which reads and
// analyzes the first record (as a header) such that specific columns can be fetched using the
// Column* methods. It returns an error if the header record cannot be read. It assumes that
// all records in the data have the same column count as the first/header record.
func NewColumnScanner(reader io.Reader, options ...Option) (*ColumnScanner, error) {
	inner := NewScanner(reader, append(options, FieldsPerRecord(0))...)
	if !inner.Scan() {
		return nil, inner.Error()
	}
	header := inner.Record()
	return &ColumnScanner{
		Scanner:      inner,
		headerRecord: header,
		columnIndex:  deriveColumnIndices(header),
	}, nil
}

func deriveColumnIndices(header []string) map[string]int {
	index := make(map[string]int)
	for i, value := range header {
		index[value] = i
	}
	return index
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
