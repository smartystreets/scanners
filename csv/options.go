package csv

// internal is used as an unexported placeholder type in the option func signature
// thereby preventing users of this package from defining additional options. Hah!
type internal struct{}

// option is a func type received by NewScanner.
// Each one allows configuration of the scanner and/or its internal *csv.Reader.
type option func(*Scanner, ...internal)

// Comma sets *csv.Reader.Comma. https://golang.org/pkg/encoding/csv/#Reader
func Comma(comma rune) option {
	return func(scanner *Scanner, _ ...internal) { scanner.reader.Comma = comma }
}

// Comment sets *csv.Reader.Comment. https://golang.org/pkg/encoding/csv/#Reader
func Comment(comment rune) option {
	return func(scanner *Scanner, _ ...internal) { scanner.reader.Comment = comment }
}

// FieldsPerRecord sets *csv.Reader.FieldsPerRecord: https://golang.org/pkg/encoding/csv/#Reader
func FieldsPerRecord(fields int) option {
	return func(scanner *Scanner, _ ...internal) { scanner.reader.FieldsPerRecord = fields }
}

// LazyQuotes sets *csv.Reader.LazyQuotes: https://golang.org/pkg/encoding/csv/#Reader
func LazyQuotes(scanner *Scanner, _ ...internal) {
	scanner.reader.LazyQuotes = true
}

// ReuseRecord sets *csv.Reader.ReuseRecord: https://golang.org/pkg/encoding/csv/#Reader
func ReuseRecord(scanner *Scanner, _ ...internal) {
	scanner.reader.ReuseRecord = true
}

// TrimLeadingSpace sets *csv.Reader.ReuseRecord: https://golang.org/pkg/encoding/csv/#Reader
func TrimLeadingSpace(scanner *Scanner, _ ...internal) {
	scanner.reader.TrimLeadingSpace = true
}

// ContinueOnError allows the scanner to continue past errors. See Scanner.Error() for
// the exact error (before the next call to Scanner.Scan())
func ContinueOnError(scanner *Scanner, _ ...internal) {
	scanner.continueOnError = true
}
