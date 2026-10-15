package subsystem

import "fmt"

// Codec is a video codec, for example mp4 or ogg
type Codec interface {
	Type() string
}

type MPEG4Codec struct{}

func (c *MPEG4Codec) Type() string {
	return "mp4"
}

type OggCodec struct{}

func (c *OggCodec) Type() string {
	return "ogg"
}

// CodecFactory creates codecs so nobody else needs a switch on the type
type CodecFactory struct{}

// Extract looks at the file and returns the codec that was used to encode it
func (f *CodecFactory) Extract(file *VideoFile) (Codec, error) {
	switch file.CodecType {
	case "mpeg4":
		fmt.Println("CodecFactory is extracting mpeg4 codec")
		return &MPEG4Codec{}, nil
	case "ogg":
		fmt.Println("CodecFactory: extracting ogg codec")
		return &OggCodec{}, nil
	}
	return nil, fmt.Errorf("no codec for type %q", file.CodecType)
}

// ForFormat returns the codec we want to convert to
func (f *CodecFactory) ForFormat(format string) (Codec, error) {
	switch format {
	case "mp4":
		return &MPEG4Codec{}, nil
	case "ogg":
		return &OggCodec{}, nil
	}
	return nil, fmt.Errorf("format %q is not supported", format)
}