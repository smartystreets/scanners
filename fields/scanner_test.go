package fields

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/smarty/assertions"
	"github.com/smarty/assertions/should"
)

func TestFieldsScanner(t *testing.T) {
	assert := assertions.New(t)

	reader := new(bytes.Buffer)
	for x := 0; x < 100; x++ {
		fmt.Fprintf(reader, "%d %d %d\n", x, x, x)
	}

	scanner := NewScanner(reader)
	x := 0
	for ; scanner.Scan(); x++ {
		X := fmt.Sprint(x)
		assert.So(scanner.Fields(), should.Resemble, []string{X, X, X})
	}

	assert.So(x, should.Equal, 100)
	assert.So(scanner.Err(), should.BeNil)
	assert.So(scanner.Scan(), should.BeFalse)
	assert.So(scanner.Text(), should.BeEmpty)
	assert.So(scanner.Fields(), should.BeEmpty)
}
