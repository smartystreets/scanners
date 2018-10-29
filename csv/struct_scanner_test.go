package csv

import (
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestStructScannerFixture(t *testing.T) {
	gunit.Run(new(StructScannerFixture), t)
}

type StructScannerFixture struct {
	*gunit.Fixture
	scanner *StructScanner
	err     error
	users   []TaggedUser
}

func (this *StructScannerFixture) Setup() {
	this.scanner, this.err = NewStructScanner(reader(csvCanon))
	this.So(this.err, should.BeNil)
}

func (this *StructScannerFixture) ScanAll() {
	for this.scanner.Scan() {
		var user TaggedUser
		this.scanner.Populate(&user)
		this.users = append(this.users, user)
	}
}

func (this *StructScannerFixture) Test() {
	this.ScanAll()

	this.So(this.scanner.Error(), should.BeNil)
	this.So(this.users, should.Resemble, []TaggedUser{
		{FirstName: "Rob", LastName: "Pike", Username: "rob"},
		{FirstName: "Ken", LastName: "Thompson", Username: "ken"},
		{FirstName: "Robert", LastName: "Griesemer", Username: "gri"},
	})
}

type TaggedUser struct {
	FirstName string `csv:"first_name"`
	LastName  string `csv:"last_name"`
	Username  string `csv:"username"`
}

func (this *StructScannerFixture) TestCannotReadHeader() {
	scanner, err := NewStructScanner(new(ErrorReader))
	this.So(scanner, should.BeNil)
	this.So(err, should.NotBeNil)
}

func (this *StructScannerFixture) TestScanIntoLessCompatibleType() {
	this.scanner.Scan()

	var nonPointer User
	this.So(this.scanner.Populate(nonPointer), should.NotBeNil)

	var nilPointer *User
	this.So(this.scanner.Populate(nilPointer), should.NotBeNil)
}

