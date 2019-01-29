package csv

import (
	"encoding/csv"
	"fmt"
	"io"
)

// Writer wraps a csv.Writer.
type Writer struct{ *csv.Writer }

// NewWriter accepts a target io.Writer and an optional comma rune
// and builds a Writer with an internal csv.Writer.
func NewWriter(w io.Writer, comma ...rune) *Writer {
	writer := csv.NewWriter(w)
	if len(comma) > 0 {
		writer.Comma = comma[0]
	}
	return &Writer{Writer: writer}
}

// WriteFields accepts zero or more interface{} values and converts
// them to strings using fmt.Sprint and writes them as a single record
// to the underlying csv.Writer. Make sure you are comfortable with
// whatever the default format is for each field value you provide.
func (this *Writer) WriteFields(fields ...interface{}) error {
	record := make([]string, len(fields))
	for f, field := range fields {
		record[f] = fmt.Sprint(field)
	}
	return this.Write(record)
}

// WriteFormattedFields accepts a format string for 0 or more fields which
// will be passed to fmt.Sprintf before being written as a single record
// to the underlying csv.Writer.
func (this *Writer) WriteFormattedFields(format string, fields ...interface{}) error {
	record := make([]string, len(fields))
	for f, field := range fields {
		record[f] = fmt.Sprintf(format, field)
	}
	return this.Write(record)
}

// WriteStringers accepts zero or more fmt.Stinger values and converts
// them to strings by calling their String() method and writes them as
// a single record to the underlying csv.Writer.
func (this *Writer) WriteStringers(fields ...fmt.Stringer) error {
	record := make([]string, len(fields))
	for f, field := range fields {
		record[f] = field.String()
	}
	return this.Write(record)
}

// WriteStrings accepts zero or more string values and writes them as
// a single record to the underlying csv.Writer.
// IMHO, it's how csv.Writer.Write should have been defined.
func (this *Writer) WriteStrings(fields ...string) error {
	return this.Writer.Write(fields)
}

// WriteStream accepts a chan []string and ranges over it, passing
// each []string as a record to the underlying csv.Writer. Like
// it's counterpart (csv.Writer.WriteAll) it calls Flush() if
// all records are written without error. It is assumed that
// the channel is or will be closed by the caller or a separate
// goroutine, otherwise this call will block indefinitely.
func (this *Writer) WriteStream(records chan []string) error {
	for record := range records {
		err := this.Write(record)
		if err != nil {
			return err
		}
	}
	this.Flush()
	return this.Error()
}
