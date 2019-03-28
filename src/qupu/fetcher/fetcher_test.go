package fetcher

import (
	"fmt"
	"testing"
)

func TestFetch(t *testing.T) {
	contents, err := Fetch("http://www.qupu123.com")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", contents)
}
