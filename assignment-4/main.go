package main

import (
	"fmt"
	"log"

	"videoconverter/facade"
)

// main is the client that only uses the facade package
func main() {
	converter := facade.NewVideoConverter()

	fmt.Println("--- convert mp4 to ogg ---")
	result, err := converter.Convert("holiday.mp4", "ogg")
	if err != nil {
		log.Fatalf("Error: %s\n", err)
	}
	fmt.Println("Result:", result)

	fmt.Println("\n--- extract audio ---")
	extractor := facade.NewAudioExtractor()
	audio, err := extractor.ExtractAudio("holiday.mp4")
	if err != nil {
		log.Fatalf("Error: %s\n", err)
	}
	fmt.Println("Result:", audio)
	
	fmt.Println("\n--- bad input ---")
	_, err = converter.Convert("notes.txt", "ogg")
	fmt.Println("Error: ", err)
	
	_, err = converter.Convert("holiday.mp4", "avi")
	fmt.Println("Error:", err)
}
