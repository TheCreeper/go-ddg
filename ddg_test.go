package ddg

import "testing"

func TestQuery(t *testing.T) {

	c := NewClient()
	q, err := c.Query("DuckDuckGo")
	if err != nil {

		t.Fatal(err)
	}

	t.Log(q.Abstract)
}

func TestFeelingLucky(t *testing.T) {

	c := NewClient()
	_, text, err := c.FeelingLucky("Ducks")
	if err != nil {

		t.Fatal(err)
	}

	t.Log(text)
}
