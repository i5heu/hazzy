! WORK IN PROGRESS !


# Hazzy

<p align="center">
  <img src=".media/logo.png"  width="33%">
</p>

## Description

Hazzy is a Go package that offers a unique approach to file hashing, particularly useful for identifying duplicates in files, either wholly or partially. It achieves this by computing hash values in chunks, enabling users to compare parts of files for potential similarities. This feature is particularly beneficial in data deduplication, storage optimization, and efficient file management.

## Features

- **Compression Ratio Calculation**: Hazzy calculates the compression ratio of a file, providing insight into the level complexity of the file's content.
- **Chunk-Based Hashing**: Files are hashed in chunks (100KB and 1KB sizes), enabling partial file comparison and detection of similarities within different files.
- **Format**: The hash output format is `(compression ratio).(hash of 100KB chunks).(hash of 1KB chunks)`.

## Advantages Over Traditional Hashing Methods

### Compared to Pure Fuzzy Hashing

- **More Informative**: Hazzy not only identifies similar files but also quantifies the level of similarity with its compression ratio and detailed chunk-based hashes.
- **Efficient Partial Matching**: Unlike traditional fuzzy hashes that provide a single hash value, Hazzy's chunk-based approach allows for more granular comparison, making it easier to locate which parts of the files are similar.

### Compared to Bloom Trees

- **Simpler Implementation**: Hazzy offers a more straightforward approach without the need for complex data structures like Bloom trees.
- **Reduced False Positives**: By providing detailed chunk-based hashes, Hazzy reduces the likelihood of false positives that can occur with Bloom tree implementations in large datasets.
- **Versatility in File Size**: Effective for both large and small files, Hazzy ensures accurate hashing and comparison regardless of file size.

## Example

This image will give this hash

<img src="./testData/smol.jpeg"  width="33%">

`4.Is.33PidPezAaudnCIjEcl0Nd5jjeOaOZL2iVhsDZMutapsNSYqe3LW2EiikBfVxYB1sWAXwmKqr0BxqAzSjbDO1Uy9krHUYjnridr5xajV72leTJLp6uFNZF1swoVkDgsiFyIZODlRgdz979lhLJU7jVmii8878wJZwgCPEBHs715C5FVlFJHjK3OHuMhbqueAPVSzBGmoUfj21T0FOb4qjfp0qn7OhDeoJ2WAM568KfhXqtZPgUvvHGHDf5n1iGz00Qvukiv8kxB9SYIFcuHILjxJ6L7SK4eNDhvjo0LR7ETyPUspX25aRiVEwhpdh2weRYj6RYkilaoFORCo4aS7QU3xLLCeP8di35LbVlxZw5HLDlKzwCGt6igYjIihoifSxwlHYOYS4Q9ujti907BCiPQlKKm8HqDFC9vqudGZMyeR14ybCcpO6c3Td6FjnTHXRILCvkCRkqsZYKeDq2mLHUUTZ2M6ORj8odKrsjSefIEjhrddnSsY7ODfgWmbc3aloYXZtnQezwDhcuEUbUezKbPYhfRglZj29MOYWriHS0Y4HnAAO9jkhrWQE9OylHf3XWuRHcjmn6Ilv605Jb1Uwer4SMWyWE9S1HD0q2qKor6HmmSywC
`


## Installation

To install Hazzy, use the following go get command:

```bash
go get -u github.com/i5heu/hazzy
```

## Usage

Here's a basic example of how to use Hazzy:

! WORK IN PROGRESS !  

```go
package main

import (
    "fmt"
    "github.com/i5heu/hazzy"
)

func main() {
    // Generate hash from a file
    hash, err := hazzy.GenerateHashFromFile("path/to/your/file")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("File Hash:", hash)

    // Generate hash from a byte slice
    data := []byte("example data")
    hash, err = hazzy.GenerateHashFromBytes(data)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Data Hash:", hash)
}
```

## Contributing

Contributions are welcome! Please feel free to submit pull requests, open issues, or suggest improvements.

## License

hazzy (c) 2023 Mia Heidenstedt and contributors

SPDX-License-Identifier: AGPL-3.0