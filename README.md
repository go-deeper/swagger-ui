# swagger-ui

This module provides a wrapper of [SwaggerUI](https://github.com/swagger-api/swagger-ui) using the Go [embed](https://pkg.go.dev/embed) [FS](https://pkg.go.dev/io/fs) type.
The module allows you to add `swagger-ui` for an existing `swagger.json` in your project.

## Install

```shell
go get github.com/go-deeper/swagger-ui
```

## Usage

```go
package main

import (
	"fmt"
	"log"
	"net/http"

	swaggerui "github.com/go-deeper/swagger-ui/pkg"
)

func main() {
	// The module requires an existing `/swagger.json` to work.
	http.HandleFunc("/swagger.json", func(rw http.ResponseWriter, _ *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		_, _ = fmt.Fprint(rw, `{"openapi": "3.1.0", "info": {"title": "OpenAPI", "version": "1.0.0"}, "paths": {}}`)
	})

	// Optionally, we will add a redirect for paths that do not have a trailing `/`.
	http.Handle("GET /docs", http.RedirectHandler("/docs/", http.StatusMovedPermanently))

	// Let's connect the module.
	http.Handle("GET /docs/", http.StripPrefix("/docs/", http.FileServerFS(swaggerui.FS)))

	log.Fatal(http.ListenAndServe(":8000", nil))
}
```

```shell
open http://localhost:8000/docs/
```
