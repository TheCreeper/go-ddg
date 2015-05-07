go-ddg
=====================

[![go-ddg](https://godoc.org/github.com/TheCreeper/go-ddg?status.png)](http://godoc.org/github.com/TheCreeper/go-ddg)

Package ddg provides an implementation of the [DuckDuckGo API](https://duckduckgo.com/api).

## Example

```
package main

import (
	"log"
	"fmt"

	"github.com/TheCreeper/go-ddg"
)

func main() {

	c := ddg.NewClient()

	_, text, err := c.FeelingLucky("Ducks")
	if err != nil {

		log.Fatal(err)
	}

	fmt.Println(text)
}
```