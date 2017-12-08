package fields

import (
	"bufio"
	"io"
	"strings"
)

type Scanner struct {
	*bufio.Scanner
}

func NewScanner(reader io.Reader) *Scanner {
	return &Scanner{Scanner: bufio.NewScanner(reader)}
}

func (this *Scanner) Fields() []string {
	return strings.Fields(this.Text())
}
