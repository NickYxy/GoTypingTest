package ast_test

import (
	"bytes"
	"flag"
	"os"
	"path/filepath"
	"testing"

	"github.com/NickYxy/GoTypingTest/instrument_trace/instrumenter/ast"
)

var update = flag.Bool("update", false, "update .golden files")

func TestRewrite(t *testing.T) {
	tests := []struct {
		goldFile string
		srcFile  string
	}{
		{
			goldFile: "no_import.golden",
			srcFile:  "no_import.go",
		},
		{
			goldFile: "with_import_no_trace.golden",
			srcFile:  "with_import_no_trace.go",
		},
		{
			goldFile: "with_import_with_trace.golden",
			srcFile:  "with_import_with_trace.go",
		},
	}

	for _, tt := range tests {
		golden := filepath.Join("testdata", tt.goldFile)
		got, err := ast.Rewrite(filepath.Join("testdata", tt.srcFile))
		if err != nil {
			t.Fatalf("rewrite failed: %v\n", err)
		}
		if *update {
			err := os.WriteFile(golden, got, 0644)
			if err != nil {
				return
			}
		}

		want, err := os.ReadFile(golden)
		if err != nil {
			t.Fatalf("open file %s failed: %v", tt.goldFile, err)
		}

		if !bytes.Equal(got, want) {
			t.Errorf("want %s, got %s", string(want), string(got))
		}
	}
}