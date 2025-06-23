// Package swaggerui contains static SwaggerUI files filesystem.
// To fill this package with a files run `make swagger-ui`
package swaggerui

import "embed"

//go:embed *.css *.html *.js *.map *.png
var FS embed.FS
