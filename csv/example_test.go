package csv_test

import (
	"fmt"
	"log"
	"strings"

	"github.com/smartystreets/scanners/csv"
)

func ExampleScanner() {
	in := strings.Join([]string{
		`first_name,last_name,username`,
		`"Rob","Pike",rob`,
		`Ken,Thompson,ken`,
		`"Robert","Griesemer","gri"`,
	}, "\n")
	scanner := csv.NewScanner(strings.NewReader(in))

	for scanner.Scan() {
		fmt.Println(scanner.Record())
	}

	if err := scanner.Error(); err != nil {
		log.Panic(err)
	}

	// Output:
	// [first_name last_name username]
	// [Rob Pike rob]
	// [Ken Thompson ken]
	// [Robert Griesemer gri]
}

// This example shows how csv.Scanner can be configured to handle other
// types of CSV files.
func ExampleScanner_options() {
	in := strings.Join([]string{
		`first_name;last_name;username`,
		`"Rob";"Pike";rob`,
		`# lines beginning with a # character are ignored`,
		`Ken;Thompson;ken`,
		`"Robert";"Griesemer";"gri"`,
	}, "\n")

	scanner := csv.NewScanner(strings.NewReader(in), csv.Comma(';'), csv.Comment('#'))

	for scanner.Scan() {
		fmt.Println(scanner.Record())
	}

	if err := scanner.Error(); err != nil {
		log.Panic(err)
	}

	// Output:
	// [first_name last_name username]
	// [Rob Pike rob]
	// [Ken Thompson ken]
	// [Robert Griesemer gri]
}
