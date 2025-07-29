package swaggerui_test

import (
	"testing"

	swaggerui "github.com/go-deeper/swagger-ui/pkg"
)

func TestFS(t *testing.T) {
	entries, err := swaggerui.FS.ReadDir(".")
	if err != nil {
		t.Fatalf("failed to read directory: %v", err)
	}

	wantNames := [...]string{
		"favicon-16x16.png", "favicon-32x32.png",
		"index.css", "index.html",
		"oauth2-redirect.html", "swagger-initializer.js",
		"swagger-ui-bundle.js", "swagger-ui-bundle.js.map",
		"swagger-ui-es-bundle-core.js", "swagger-ui-es-bundle-core.js.map",
		"swagger-ui-es-bundle.js", "swagger-ui-es-bundle.js.map",
		"swagger-ui-standalone-preset.js", "swagger-ui-standalone-preset.js.map",
		"swagger-ui.css", "swagger-ui.css.map",
		"swagger-ui.js", "swagger-ui.js.map",
	}

	if len(entries) != len(wantNames) {
		t.Fatalf("want %d entries, got %d", len(wantNames), len(entries))
	}

	for i, wantName := range wantNames {
		if gotName := entries[i].Name(); wantName != gotName {
			t.Fatalf("want entry name %q, got %q", wantName, gotName)
		}
	}
}
