package csv

import (
	"encoding/csv"
	"io"
	"strings"
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestScanAllFixture(t *testing.T) {
	gunit.Run(new(ScanAllFixture), t)
}

type ScanAllFixture struct {
	*gunit.Fixture
}

func (this *ScanAllFixture) scanAll(inputs []string, options ...Option) (scanned []Record) {
	scanner := NewScanner(reader(inputs), options...)
	line := 1
	for ; scanner.Scan(); line++ {
		scanned = append(scanned, Record{
			line:   line,
			record: scanner.Record(),
			err:    scanner.Error(),
		})
	}
	if err := scanner.Error(); err != nil {
		scanned = append(scanned, Record{
			line: line,
			err:  err,
		})
	}
	return scanned
}

func (this *ScanAllFixture) TestCanonical() {
	scanned := this.scanAll(csvCanon, Comma(','), FieldsPerRecord(3))
	this.So(scanned, should.Resemble, expectedScannedOutput)
}

func (this *ScanAllFixture) TestCanonicalWithOptions() {
	scanned := this.scanAll(csvCanonRequiringConfigOptions, Comma(';'), Comment('#'))
	this.So(scanned, should.Resemble, expectedScannedOutput)
}

func (this *ScanAllFixture) TestOptions() {
	scanner := NewScanner(nil, ReuseRecord(true), TrimLeadingSpace(true), LazyQuotes(true))
	this.So(scanner.reader.ReuseRecord, should.BeTrue)
	this.So(scanner.reader.LazyQuotes, should.BeTrue)
	this.So(scanner.reader.TrimLeadingSpace, should.BeTrue)
}

func (this *ScanAllFixture) TestInconsistentFieldCounts_ContinueOnError() {
	scanned := this.scanAll(csvCanonInconsistentFieldCounts, ContinueOnError(true))
	this.So(scanned, should.Resemble, []Record{
		{line: 1, record: []string{"1", "2", "3"}, err: nil},
		{line: 2, record: []string{"1", "2", "3", "4"}, err: &csv.ParseError{StartLine: 2, Line: 2, Column: 0, Err: csv.ErrFieldCount}},
		{line: 3, record: []string{"1", "2", "3"}, err: nil},
	})
}

func (this *ScanAllFixture) TestInconsistentFieldCounts_HaltOnError() {
	scanned := this.scanAll(csvCanonInconsistentFieldCounts)
	this.So(scanned, should.Resemble, []Record{
		{line: 1, record: []string{"1", "2", "3"}, err: nil},
		{line: 2, record: nil, err: &csv.ParseError{StartLine: 2, Line: 2, Column: 0, Err: csv.ErrFieldCount}},
	})
}

func (this *ScanAllFixture) TestCallsToScanAfterEOFReturnFalse() {
	scanner := NewScanner(strings.NewReader("1,2,3"), Comma(','))

	this.So(scanner.Scan(), should.BeTrue)
	this.So(scanner.Record(), should.Resemble, []string{"1", "2", "3"})
	this.So(scanner.Error(), should.BeNil)

	for x := 0; x < 100; x++ {
		this.So(scanner.Scan(), should.BeFalse)
		this.So(scanner.Record(), should.BeNil)
		this.So(scanner.Error(), should.BeNil)
	}
}

func reader(lines []string) io.Reader {
	return strings.NewReader(strings.Join(lines, "\n"))
}

var ( // https://golang.org/pkg/encoding/csv/#example_Reader
	csvCanon = []string{
		"first_name,last_name,username",
		`"Rob","Pike",rob`,
		`Ken,Thompson,ken`,
		`"Robert","Griesemer","gri"`,
	}
	csvCanonRequiringConfigOptions = []string{
		`first_name;last_name;username`,
		`"Rob";"Pike";rob`,
		`# lines beginning with a # character are ignored`,
		`Ken;Thompson;ken`,
		`"Robert";"Griesemer";"gri"`,
	}
	csvCanonInconsistentFieldCounts = []string{
		`1,2,3`,
		`1,2,3,4`,
		`1,2,3`,
	}
	expectedScannedOutput = []Record{
		{1, []string{"first_name", "last_name", "username"}, nil},
		{2, []string{"Rob", "Pike", "rob"}, nil},
		{3, []string{"Ken", "Thompson", "ken"}, nil},
		{4, []string{"Robert", "Griesemer", "gri"}, nil},
	}
)

type Record struct {
	line   int
	record []string
	err    error
}
