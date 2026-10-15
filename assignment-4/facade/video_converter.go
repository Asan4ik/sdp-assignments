package facade

import (
	"fmt"
	"strings"

	"videoconverter/subsystem"
)

/* VideoConverter is the facade. The client only gives a file name and a format,
   and the order of the subsystem calls is handled here */
type VideoConverter struct {
	factory *subsystem.CodecFactory
	reader  *subsystem.BitrateReader
	mixer   *subsystem.AudioMixer
}

// NewVideoConverter creates a converter with all subsystem parts ready
func NewVideoConverter() *VideoConverter {
	return &VideoConverter{
		factory: &subsystem.CodecFactory{},
		reader:  &subsystem.BitrateReader{},
		mixer:   &subsystem.AudioMixer{},
	}
}

// Convert converts the file to the given format and returns the name of the new file
func (v *VideoConverter) Convert(filename, format string) (string, error) {
	format = strings.ToLower(format)

	// check the target format first so we fail before doing any work
	destinationCodec, err := v.factory.ForFormat(format)
	if err != nil {
		return "", fmt.Errorf("convert failed: %w", err)
	}

	file, err := subsystem.NewVideoFile(filename)
	if err != nil {
		return "", fmt.Errorf("convert failed: %w", err)
	}

	sourceCodec, err := v.factory.Extract(file)
	if err != nil {
		return "", fmt.Errorf("convert failed: %w", err)
	}

	if sourceCodec.Type() == destinationCodec.Type() {
		return "", fmt.Errorf("convert failed: %s is already %s", filename, format)
	}

	buffer := v.reader.Read(file, sourceCodec)
	result, err := v.reader.Convert(buffer, destinationCodec)
	if err != nil {
		return "", fmt.Errorf("convert failed: %w", err)
	}
	result = v.mixer.Fix(result)

	newName := withExtension(filename, format)
	fmt.Printf("VideoConverter: saved %s (%d bytes)\n", newName, len(result))
	return newName, nil
}

// withExtension swaps the extension of filename for the given one
func withExtension(filename, format string) string {
	dot := strings.LastIndex(filename, ".")
	return filename[:dot+1] + format
}
