# swagger-ui

## Install

```shell
go get github.com/go-deeper/swagger-ui
```

## Usage

```go
package main

import (
	"log"
	"net/http"

	swaggerui "github.com/go-deeper/swagger-ui/pkg"
)

const swaggerJSON = `{"openapi": "3.1.0", "info": {"title": "OpenAPI", "version": "1.0.0"}, "paths": {}}`

func main() {
	http.HandleFunc("/swagger.json", func(rw http.ResponseWriter, _ *http.Request) {
		_, _ = rw.Write([]byte(swaggerJSON))
	})

	http.Handle("GET /docs/", http.StripPrefix("/docs/", http.FileServerFS(swaggerui.FS)))
	http.Handle("GET /docs", http.RedirectHandler("/docs/", http.StatusMovedPermanently))

	log.Fatal(http.ListenAndServe(":8000", nil))
}
```

```shell
open http://localhost:8000/docs/
```
