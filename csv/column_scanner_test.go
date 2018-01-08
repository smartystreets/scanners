package csv

import (
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestColumnScannerFixture(t *testing.T) {
	gunit.Run(new(ColumnScannerFixture), t)
}

type ColumnScannerFixture struct {
	*gunit.Fixture

	scanner *ColumnScanner
	err     error
	users   []User
}

func (this *ColumnScannerFixture) Setup() {
	this.scanner, this.err = NewColumnScanner(reader(csvCanon))
	this.So(this.err, should.BeNil)
	this.So(this.scanner.Header(), should.Resemble, []string{"first_name", "last_name", "username"})
}

func (this *ColumnScannerFixture) ScanAllUsers() {
	for this.scanner.Scan() {
		this.users = append(this.users, this.scanUser())
	}
}

func (this *ColumnScannerFixture) Test() {
	this.ScanAllUsers()

	this.So(this.scanner.Error(), should.BeNil)
	this.So(this.users, should.Resemble, []User{
		{FirstName: "Rob", LastName: "Pike", Username: "rob"},
		{FirstName: "Ken", LastName: "Thompson", Username: "ken"},
		{FirstName: "Robert", LastName: "Griesemer", Username: "gri"},
	})
}

func (this *ColumnScannerFixture) scanUser() User {
	return User{
		FirstName: this.scanner.Column(this.scanner.Header()[0]),
		LastName:  this.scanner.Column(this.scanner.Header()[1]),
		Username:  this.scanner.Column(this.scanner.Header()[2]),
	}
}

type User struct {
	FirstName string
	LastName  string
	Username  string
}
