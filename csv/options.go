package csv

// Option is a func type received by NewScanner.
// Each one allows configuration of the scanner and/or its internal *csv.Reader.
type Option func(*Scanner)

// Options (a singleton instance) provides access to built-in functional options.
var Options options

type options struct{}

// ContinueOnError controls scanner behavior in error scenarios.
// If true is passed, continue scanning until io.EOF is reached.
// If false is passed (default), any error encountered during scanning
// will result in the next call to Scan returning false and
// the Scanner may be considered dead. See Scanner.Error() for the exact error
// (before the next call to Scanner.Scan()).
// See https://golang.org/pkg/encoding/csv/#pkg-variables
// and https://golang.org/pkg/encoding/csv/#ParseError
// for more information regarding possible error values.
func (_ options) ContinueOnError(continue_ bool) Option {
	return func(s *Scanner) {
		s.continueOnError = continue_
	}
}
func (_ options) Comma(comma rune) Option {
	return func(s *Scanner) {
		s.reader.Comma = comma
	}
}
func (_ options) Comment(comment rune) Option {
	return func(s *Scanner) {
		s.reader.Comment = comment
	}
}
func (_ options) FieldsPerRecord(fields int) Option {
	return func(s *Scanner) {
		s.reader.FieldsPerRecord = fields
	}
}
func (_ options) LazyQuotes(lazy bool) Option {
	return func(s *Scanner) {
		s.reader.LazyQuotes = lazy
	}
}
func (_ options) ReuseRecord(reuse bool) Option {
	return func(s *Scanner) {
		s.reader.ReuseRecord = reuse
	}
}
func (_ options) TrimLeadingSpace(trim bool) Option {
	return func(s *Scanner) {
		s.reader.TrimLeadingSpace = trim
	}
}
func (o options) SkipHeaderRecord() Option {
	return func(s *Scanner) {
		s.Scan()
	}
}
func (_ options) SkipRecords(count int) Option {
	return func(s *Scanner) {
		for x := 0; x < count; x++ {
			s.Scan()
		}
	}
}
