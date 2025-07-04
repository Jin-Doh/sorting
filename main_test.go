package main

import (
	"os"
	"testing"
)

func TestSortingIntegration(t *testing.T) {
	tempFile, err := os.CreateTemp("", "testpkglist-*.txt")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	input := "zeta\n\nalpha\nbeta\nalpha\n  \ngamma\n"
	if _, err := tempFile.WriteString(input); err != nil {
		t.Fatalf("failed to write to temp file: %v", err)
	}
	if err := tempFile.Close(); err != nil {
		t.Fatalf("failed to close temp file: %v", err)
	}

	os.Args = []string{"cmd", "--file", tempFile.Name()}
	main()

	outputBytes, err := os.ReadFile(tempFile.Name())
	if err != nil {
		t.Fatalf("failed to read output file: %v", err)
	}
	output := string(outputBytes)
	expected := "alpha\nbeta\ngamma\nzeta\n"
	if output != expected {
		t.Errorf("unexpected output.\nGot:\n%s\nWant:\n%s", output, expected)
	}
}
