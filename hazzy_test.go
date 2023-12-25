package hazzy

import (
	"io"
	"os"
	"strings"
	"testing"
)

// TestGenerateHashFromFile tests the hash generation from a file.
func TestGenerateHashFromFile(t *testing.T) {
	filePath := "./testData/big.png"

	got, err := GenerateHashFromFile(filePath)
	if err != nil {
		t.Errorf("GenerateHashFromFile() error = %v", err)
		return
	}

	// print hash for manual inspection, will printed to stdout
	t.Logf("hash: %v", got)

	parts := strings.Split(got, ".")
	if len(parts) != 3 {
		t.Errorf("GenerateHashFromFile() got = %v, want three parts separated by '.'", got)
	}
}

func TestGenerateHashFromFile2(t *testing.T) {
	filePath := "./testData/text.txt"

	got, err := GenerateHashFromFile(filePath)
	if err != nil {
		t.Errorf("GenerateHashFromFile() error = %v", err)
		return
	}

	// print hash for manual inspection, will printed to stdout
	t.Logf("hash: %v", got)

	parts := strings.Split(got, ".")
	if len(parts) != 3 {
		t.Errorf("GenerateHashFromFile() got = %v, want three parts separated by '.'", got)
	}
}

// TestGenerateHashFromBytes tests the hash generation from a byte slice.
func TestGenerateHashFromBytes(t *testing.T) {

	//read test bytes into a byte slice from a file
	filePath := "./testData/big.png"
	file, err := os.Open(filePath)
	if err != nil {
		t.Errorf("GenerateHashFromBytes() error = %v", err)
		return
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		t.Errorf("GenerateHashFromBytes() error = %v", err)
		return
	}

	got, err := GenerateHashFromBytes(data)
	if err != nil {
		t.Errorf("GenerateHashFromBytes() error = %v", err)
		return
	}

	parts := strings.Split(got, ".")
	if len(parts) != 3 {
		t.Errorf("GenerateHashFromBytes() got = %v, want three parts separated by '.'", got)
	}
}
