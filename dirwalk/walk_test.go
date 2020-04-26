package dirwalk

import "testing"

func TestWalk(t *testing.T) {
	files, err := Walk(".", nil)
	if err != nil {
		t.Fatal(err)
	}

	for _, f := range files {
		t.Log(f.Name, f.Path, f.FullPath())
	}
}
