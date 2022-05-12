# go-snipeit

[![GoDoc](https://img.shields.io/static/v1?label=godoc&message=reference&color=blue)](https://pkg.go.dev/github.com/euracresearch/go-snipeit)
![tests](https://github.com/euracresearch/go-snipeit/workflows/tests/badge.svg)

go-snipeit is a Go client library for accessing the [SNIPE-IT API](https://snipe-it.readme.io/reference).

## Notice

The library is **under development** and currently supports only API endpoints and methods used internally by Eurac Research. 

If you need a full API client we suggest to download the [Swagger/OpenAPI Specification](https://snipe-it.readme.io/openapi/) and use a tool like [oapi-codegen](https://github.com/bonitoo-io/oapi-codegen) to generate the Go client code.

Supported API endpoints:
 
- Hardware
	- Return a listing of assets: https://snipe-it.readme.io/reference/hardware-list
- Locations
	- List locations: https://snipe-it.readme.io/reference/locations
	- Get location details by id: https://snipe-it.readme.io/reference/locations-1
- Categories
	- List categories: https://snipe-it.readme.io/reference/categories-1
	- Return a Category by id: https://snipe-it.readme.io/reference/categoriesid-3

During development a version was tagged as `v0.0.1` by mistake. In order to always get the latest version use the following command `go get github.com/euracresearch/go-snipeit@master` until we reach a `v1.0.0`.

Of course, contributions are very welcome!

## Example

First get the latest version by: 

```
$ go get github.com/euracresearch/go-snipeit@master
```

```go
package main

import (
	"fmt"
	"log"

	"github.com/euracresearch/go-snipeit"
)

func main() {
	// Create a new client using SnipeIT API base URL and and authentication token
	client, err := snipeit.NewClient("https://develop.snipeitapp.com/api/v1/", "my-token")
	if err != nil {
		log.Fatal(err)
	}

	opts := &snipeit.HardwareListOptions{
		Limit: 10,
	}
	// Retrieve all hardware limited to 10
	hardware, _, err := client.Hardware.List(opts)
	if err != nil {
		log.Fatal(err)
	}

	// Print all hardware AssetTags
	for _, h := range hardware {
		fmt.Println(h.AssetTag)
	}
}
```

## License

The library is distributed under the BSD-style license found in [LICENSE](./LICENSE) file.