package parser

import (
	"fmt"
	"testing"
)

func TestParseList(t *testing.T) {
	content := ParseList("https://m.tititoy2688.com//query/book?id=1405")
	fmt.Printf("%s\n", content)
}
