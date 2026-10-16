package facade

import "testing"

func TestConvertMp4ToOgg(t *testing.T) {
	c := NewVideoConverter()
	got, err := c.Convert("clip.mp4", "ogg")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != "clip.ogg" {
		t.Errorf("expected clip.ogg, got %s", got)
	}
}

func TestConvertErrors(t *testing.T) {
	c := NewVideoConverter()
	cases := []struct{ file, format string }{
		{"notes.txt", "ogg"},
		{"noextension", "ogg"},
		{"clip.mp4", "avi"},
		{"clip.mp4", "mp4"},
	}
	for _, tc := range cases {
		if _, err := c.Convert(tc.file, tc.format); err == nil {
			t.Errorf("expected error for %s -> %s", tc.file, tc.format)
		}
	}
}

func TestExtractAudio(t *testing.T) {
	e := NewAudioExtractor()
	got, err := e.ExtractAudio("clip.ogg")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != "clip.mp3" {
		t.Errorf("expected clip.mp3, got %s", got)
	}
}
