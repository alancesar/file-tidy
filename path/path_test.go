package path

import (
	"os"
	"strings"
	"testing"
)

func TestBuilder_FromPattern(t *testing.T) {
	myStruct := struct {
		Track  int
		Title  string
		Artist string
		Album  string
		Year   int
	}{
		Track:  7,
		Title:  "Basket Case",
		Artist: "Green Day",
		Album:  "Dookie",
		Year:   1994,
	}

	path, err := NewBuilder().
		FromPattern("{{.Artist}}/[{{.Year}}] {{.Album}}/{{printf \"%02d\" .Track}} - {{.Title}}", myStruct)
	if err != nil {
		t.Error(err)
	}

	expected := strings.ReplaceAll("Green Day/[1994] Dookie/07 - Basket Case", "/", string(os.PathSeparator))
	if expected != path {
		t.Error("unexpected path")
	}
}
