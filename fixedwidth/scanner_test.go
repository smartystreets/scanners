package fixedwidth

import (
	"bytes"
	"strings"
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestScannerFixture(t *testing.T) {
	gunit.Run(new(ScannerFixture), t)
}

type ScannerFixture struct {
	*gunit.Fixture
}

func (this *ScannerFixture) TestScanner() {
	scanner := NewScanner(bytes.NewBufferString(strings.Repeat("122333444455555\n", 10)))
	records := 0
	for ; scanner.Scan(); records++ {
		this.So(scanner.Field(Field(0, 1)), should.Equal, "1")
		this.So(scanner.Field(Field(1, 2)), should.Equal, "22")
		this.So(scanner.Field(Field(3, 3)), should.Equal, "333")
		this.So(scanner.Field(Field(6, 4)), should.Equal, "4444")
		this.So(scanner.Field(Field(10, 5)), should.Equal, "55555")
	}
	this.So(records, should.Equal, 10)
	this.So(scanner.Err(), should.BeNil)
}
