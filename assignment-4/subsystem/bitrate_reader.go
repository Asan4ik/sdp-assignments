package subsystem

import "fmt"

// BitrateReader reads video data and re-encodes it
type BitrateReader struct{}

// Read reads the raw data of the file using the source codec
func (b *BitrateReader) Read(file *VideoFile, source Codec) []byte {
	fmt.Printf("BitrateReader is reading %s using %s codec\n", file.Name, source.Type())
	return []byte("raw:" + file.Name)
}

// Convert re-encodes the data for the destination codec
func (b *BitrateReader) Convert(buffer []byte, destination Codec) ([]byte, error) {
	if len(buffer) == 0 {
		return nil, fmt.Errorf("buffer is empty, nothing to convert")
	}
	fmt.Printf("BitrateReader is converting buffer to %s\n", destination.Type())
	return append(buffer, []byte("->"+destination.Type())...), nil
}