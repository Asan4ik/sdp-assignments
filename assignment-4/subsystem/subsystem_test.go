package subsystem

import "testing"

// the subsystem should work without the facade
func TestUseSubsystemDirectly(t *testing.T) {
	file, err := NewVideoFile("clip.ogg")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	factory := &CodecFactory{}
	source, err := factory.Extract(file)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	target, err := factory.ForFormat("mp4")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	reader := &BitrateReader{}
	if _, err := reader.Convert(reader.Read(file, source), target); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestConvertEmptyBuffer(t *testing.T) {
	reader := &BitrateReader{}
	if _, err := reader.Convert(nil, &OggCodec{}); err == nil {
		t.Error("expected error for empty buffer")
	}
}

func TestNewVideoFileErrors(t *testing.T) {
	for _, name := range []string{"noextension", "notes.txt"} {
		if _, err := NewVideoFile(name); err == nil {
			t.Errorf("expected error for %q", name)
		}
	}
}
