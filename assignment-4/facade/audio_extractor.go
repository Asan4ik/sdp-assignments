package facade

import (
	"fmt"

	"videoconverter/subsystem"
)

/* AudioExtractor is a second facade. It uses a different part of the same
   subsystem, so we keep it away from VideoConverter */
type AudioExtractor struct {
	factory *subsystem.CodecFactory
	reader *subsystem.BitrateReader
	mixer *subsystem.AudioMixer
}

// NewAudioExtractor creates an extractor with all subsystem parts ready
func NewAudioExtractor() *AudioExtractor {
	return &AudioExtractor{
		factory: &subsystem.CodecFactory{},
		reader: &subsystem.BitrateReader{},
		mixer: &subsystem.AudioMixer{},
	}
}

// ExtractAudio returns the name of the audio file made from the video
func (a *AudioExtractor) ExtractAudio(filename string) (string, error) {
	file, err := subsystem.NewVideoFile(filename)
	if err != nil {
		return "", fmt.Errorf("extract audio failed: %w", err)
	}

	codec, err := a.factory.Extract(file)
	if err != nil {
		return "", fmt.Errorf("extract audio failed: %w", err)
	}

	buffer := a.reader.Read(file, codec)
	audio := a.mixer.Extract(buffer)

	newName := withExtension(filename, "mp3")
	fmt.Printf("AudioExtractor: saved %s (%d bytes)\n", newName, len(audio))
	return newName, nil
}