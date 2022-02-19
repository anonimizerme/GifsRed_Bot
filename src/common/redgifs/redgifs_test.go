package redgifs

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	list, err := Search("movie")

	if err != nil {
		t.Errorf("Error %v", err)
	}

	fmt.Println(list)
}
