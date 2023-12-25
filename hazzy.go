// Package hazzy provides functionality to compute a hash of a file or byte slice
// with the format: (compression ratio).(hash of 100KB chunks).(hash of 1KB chunks).
package hazzy

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"strconv"
)

// hashChunk computes a simple hash for a byte slice.
func hashChunk(chunk []byte) string {
	var sum uint16 = 0
	for _, b := range chunk {
		sum += uint16(b)
	}

	characters := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	index1 := sum % 62
	index2 := (sum / 62) % 62

	return string(characters[index1]) + string(characters[index2])
}

// compressAndHash computes the compression ratio and hashes the data from an io.Reader.
func compressAndHash(reader io.Reader) (string, string, int, error) {
	var originalSize, compressedSize int64
	hash100KB, hash1KB := "", ""

	buf := &bytes.Buffer{}
	gzipWriter := gzip.NewWriter(buf)
	chunk := make([]byte, 100*1024) // 100KB chunks
	isFirstChunk := true

	for {
		bytesRead, err := reader.Read(chunk)
		if err != nil {
			if err == io.EOF {
				// Check if this is the only chunk and it's smaller than 100KB
				if isFirstChunk && originalSize < 100*1024 {
					hash100KB = hashChunk(chunk[:bytesRead]) // Hash the entire content
				}
				break
			}
			return "", "", 0, err
		}
		originalSize += int64(bytesRead)

		// Specify the compression level here
		gzipWriter, err := gzip.NewWriterLevel(buf, gzip.BestSpeed)
		if err != nil {
			return "", "", 0, err
		}

		if _, err := gzipWriter.Write(chunk[:bytesRead]); err != nil {
			return "", "", 0, err
		}

		// For the first chunk, or a full 100KB chunk, calculate the hash
		if isFirstChunk || bytesRead == 100*1024 {
			hash100KB += hashChunk(chunk[:bytesRead])
			isFirstChunk = false
		}

		for i := 0; i < bytesRead; i += 1024 {
			end := i + 1024
			if end > bytesRead {
				end = bytesRead
			}
			hash1KB += hashChunk(chunk[i:end])
		}
	}

	if err := gzipWriter.Close(); err != nil {
		return "", "", 0, err
	}
	compressedSize = int64(buf.Len())

	// Calculate the compression ratio as an integer between 0 and 1000
	fmt.Println("originalSize", originalSize)
	fmt.Println("compressedSize", compressedSize)
	compressionReductionPercentage := int((1 - float64(compressedSize)/float64(originalSize)) * 1000)
	return hash100KB, hash1KB, compressionReductionPercentage, nil
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

	return strconv.Itoa(ratio) + "." + hash100KB + "." + hash1KB, nil
}

// GenerateHashFromBytes generates a hash from a byte slice.
func GenerateHashFromBytes(data []byte) (string, error) {
	reader := bytes.NewReader(data)
	hash100KB, hash1KB, ratio, err := compressAndHash(reader)
	if err != nil {
		return "", err
	}

	return strconv.Itoa(ratio) + "." + hash100KB + "." + hash1KB, nil
}
