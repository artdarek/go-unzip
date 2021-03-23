Package go-unzip
===================

Package go-unzip provides a very simple library to extract zip archive

## Installation
```shell
go get -u github.com/artdarek/go-unzip
```

## Examples

```go
package main

import (
	"fmt"

	"github.com/artdarek/go-unzip/pkg/unzip"
)

func main() {
	uz := unzip.New()

	files, err := uz.Extract("./data/file.zip", "./data/directory")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("extracted files count: %d", len(files))
	fmt.Printf("files list: %v", files)
}
```

## Contributing

Pull requests, bug fixes and issue reports are welcome.

Before proposing a change, please discuss your change by raising an issue.
