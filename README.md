# M3U Generator

A command-line tool written in Go that automatically generates M3U playlists for music directories. The tool generates separate M3U files for each first-level subdirectory, making it easy to manage music collections organized in folders.

## Features

- Generates M3U playlists for first-level subdirectories only
- Supports multiple audio formats (mp3, wav, flac, m4a, ogg, aac)
- Ignores nested subdirectories deeper than one level
- Fast and efficient processing
- Simple command-line interface

## Installation

```bash
go install github.com/linnv/m3ugenerator@latest
```

Or clone and build from source:

```bash
git clone https://github.com/linnv/m3ugenerator.git
cd m3ugenerator
go build
```

## Usage

```bash
./m3ugenerator -dir <path-to-music-directory>
```

### Example

Given the following directory structure:

```
music/
├── music.m3u
├── subdir1
│   ├── song1.mp3
│   └── song2.wav
├── subdir2
│   ├── nested_subdir
│   │   └── song4.m4a
│   └── song3.flac
└── subdir3
    ├── song5.ogg
    └── song6.aac
```

Running the command:

```bash
./m3ugenerator -dir ./music
```

Will generate:

```
music/
├── music.m3u
├── subdir1
│   ├── song1.mp3
│   ├── song2.wav
│   └── subdir1.m3u
├── subdir2
│   ├── nested_subdir
│   │   └── song4.m4a
│   ├── song3.flac
│   └── subdir2.m3u
└── subdir3
    ├── song5.ogg
    ├── song6.aac
    └── subdir3.m3u
```

Note: Files in directories deeper than one level (e.g., `nested_subdir`) will not have M3U files generated.

## Supported Audio Formats

- MP3 (.mp3)
- WAV (.wav)
- FLAC (.flac)
- M4A (.m4a)
- OGG (.ogg)
- AAC (.aac)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

[MIT License](LICENSE)
