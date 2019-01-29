package content_test

import (
	"github.com/sundogrd-templates/golang-api-template/services/content"
	"testing"
)

// TestContentFindOne ...
func TestUnmarshalContentExtraJson(t *testing.T) {
	testStr := "{}"
	extra, err := content.UnmarshalContentExtraJson(testStr)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(extra)
}