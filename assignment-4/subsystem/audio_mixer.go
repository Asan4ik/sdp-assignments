package subsystem

import "fmt"

// AudioMixer works with the audio of a video
type AudioMixer struct{}

// Fix syncs the audio with the video after conversion
func (a *AudioMixer) Fix(data []byte) []byte {
	fmt.Println("AudioMixer is fixing audio")
	return append(data, []byte("+audio")...)
}

// Extract only returns the audio part, no video
func (a *AudioMixer) Extract(data []byte) []byte {
	fmt.Println("AudioMixer is extracting audio track")
	return append(data, []byte("+audio-only")...)
}