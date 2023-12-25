// Package hazzy provides functionality to compute a hash of a file or byte slice
// with the format: (compression ratio).(hash of 100KB chunks).(hash of 1KB chunks).
package hazzy

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

// characters for hashing
const characters = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// hashChunk computes a hash for a byte slice.
// It calculates two indices based on the sum of byte values
// and uses these indices to pick two characters from a predefined string.
func hashChunk(chunk []byte) string {
	var sum uint16
	for _, b := range chunk {
		sum += uint16(b)
	}

	// Calculate two indices from the sum
	index1 := sum % 62 // Remainder when divided by 62
	index2 := (sum / 62) % 62

	return string(characters[index1]) + string(characters[index2])
}

// compressAndHash computes the compression ratio and hashes the data.
func compressAndHash(reader io.Reader) (hash100KB, hash1KB string, ratio int, err error) {
	var originalSize, compressedSize int64
	buf := &bytes.Buffer{}

	// Specify the compression level so that the hash is consistent
	gzipWriter, err := gzip.NewWriterLevel(buf, gzip.BestCompression)
	if err != nil {
		return "", "", 0, err
	}

	processChunk := func(chunk []byte, isFirstChunk bool) error {
		// Hash the first 100KB or full chunk
		if isFirstChunk {
			hash100KB += hashChunk(chunk)
		}

		// Hash in 1KB segments
		for i := 0; i < len(chunk); i += 1024 {
			end := i + 1024
			if end > len(chunk) {
				end = len(chunk)
			}
			hash1KB += hashChunk(chunk[i:end])
		}
		return nil
	}

	isFirstChunk := true
	for {
		chunk := make([]byte, 100*1024) // 100KB chunks
		bytesRead, err := reader.Read(chunk)
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", "", 0, err
		}

		originalSize += int64(bytesRead)
		if _, err := gzipWriter.Write(chunk[:bytesRead]); err != nil {
			return "", "", 0, err
		}

		if err := processChunk(chunk[:bytesRead], isFirstChunk); err != nil {
			return "", "", 0, err
		}
		isFirstChunk = false
	}

	// Close the gzip writer and calculate compression
	if err := gzipWriter.Close(); err != nil {
		return "", "", 0, err
	}
	compressedSize = int64(buf.Len())
	ratio = int((1 - float64(compressedSize)/float64(originalSize)) * 1000)
	return
}

// GenerateHashFromFile generates a hash from a file path.
func GenerateHashFromFile(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash100KB, hash1KB, ratio, err := compressAndHash(file)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d.%s.%s", ratio, hash100KB, hash1KB), nil
}

// GenerateHashFromBytes generates a hash from a byte slice.
func GenerateHashFromBytes(data []byte) (string, error) {
	reader := bytes.NewReader(data)
	hash100KB, hash1KB, ratio, err := compressAndHash(reader)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d.%s.%s", ratio, hash100KB, hash1KB), nil
}
