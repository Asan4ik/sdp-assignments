# Video Converter Facade (Go)

Assignment 4 for Software Design Patterns. It shows the Facade pattern using a video converter

The converter has a few small classes working together (codec factory, bitrate reader, audio mixer)
The client does not want to deal with them, so the facade gives it one simple method:

```go
converter := facade.NewVideoConverter()
newFile, err := converter.Convert("holiday.mp4", "ogg")
```

## Structure

```
.
├── main.go                    client, only imports the facade package
├── facade/
│   ├── video_converter.go     main facade: Convert(filename, format)
│   ├── audio_extractor.go     second facade: ExtractAudio(filename)
│   └── facade_test.go
├── subsystem/
│   ├── video_file.go
│   ├── codec.go               Codec interface, MPEG4Codec, OggCodec, CodecFactory
│   ├── bitrate_reader.go
│   └── audio_mixer.go
└── docs/
    ├── class-diagram.puml
    └── class-diagram.png
```

The subsystem package does not import the facade package, so it can still be used directly

## Run

```
go run .
go test ./...
```

Needs Go 1.27 or newer