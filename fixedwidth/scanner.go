package fixedwidth

import (
	"bufio"
	"io"
)

type Substring func(line string) (field string)

func Field(index, width int) Substring {
	return func(line string) string {
		return line[index : index+width]
	}
}

type Scanner struct {
	*bufio.Scanner
}

func NewScanner(reader io.Reader) *Scanner {
	return &Scanner{Scanner: bufio.NewScanner(reader)}
}

func (this *Scanner) Field(field Substring) string {
	return field(this.Text())
}

func (this *Scanner) Fields(fields ...Substring) (values []string) {
	for _, field := range fields {
		values = append(values, this.Field(field))
	}
	return values
}
