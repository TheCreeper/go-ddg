package ddg

import (
	"fmt"
	"testing"
)

func TestQuery(t *testing.T) {

	c := NewClient()

	q, err := c.Query("DuckDuckGo")
	if err != nil {

		t.Error(err)
	}

	fmt.Println(q.Abstract)
}

func TestFeelingLucky(t *testing.T) {

	c := NewClient()

	_, text, err := c.FeelingLucky("Ducks")
	if err != nil {

		t.Error(err)
	}

	fmt.Println(text)
}
