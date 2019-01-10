package csv

// Option is a func type received by NewScanner.
// Each one allows configuration of the scanner and/or its internal *csv.Reader.
type Option func(*Scanner)

// ContinueOnError controls scanner behavior in error scenarios.
// If true is passed, continue scanning until io.EOF is reached.
// If false is passed (default), any error encountered during scanning
// will result in the next call to Scan returning false and
// the Scanner may be considered dead. See Scanner.Error() for the exact error
// (before the next call to Scanner.Scan()).
// See https://golang.org/pkg/encoding/csv/#pkg-variables
// and https://golang.org/pkg/encoding/csv/#ParseError
// for more information regarding possible error values.
func ContinueOnError(continue_ bool) Option { return func(s *Scanner) { s.continueOnError = continue_ } }
func Comma(comma rune) Option               { return func(s *Scanner) { s.reader.Comma = comma } }
func Comment(comment rune) Option           { return func(s *Scanner) { s.reader.Comment = comment } }
func FieldsPerRecord(fields int) Option     { return func(s *Scanner) { s.reader.FieldsPerRecord = fields } }
func LazyQuotes(lazy bool) Option           { return func(s *Scanner) { s.reader.LazyQuotes = lazy } }
func ReuseRecord(reuse bool) Option         { return func(s *Scanner) { s.reader.ReuseRecord = reuse } }
func TrimLeadingSpace(trim bool) Option     { return func(s *Scanner) { s.reader.TrimLeadingSpace = trim } }
func SkipHeaderRecord() Option              { return SkipRecords(1) }
func SkipRecords(count int) Option {
	return func(s *Scanner) {
		for x := 0; x < count; x++ {
			s.Scan()
		}
	}
}
