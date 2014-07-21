Usage
=====

```go
package main

import (
	`fmt`
	`github.com/Hellyna/go-whois`
	`os`
)

func main() {
	res, err := whois.Whois(os.Args[1])

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(res)
}
```

<!---
vim:ts=4 sw=4 noet:
-->
