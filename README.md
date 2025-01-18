
## m3u generator
 
### a tool to generate m3u files for music only for the first level sub dirs e.g. 

```
❯ tree music/
music/
├── music.m3u
├── subdir1
│   ├── song1.mp3
│   └── song2.wav
├── subdir2
│   ├── nested_subdir
│   │   └── song4.m4a
│   └── song3.flac
└── subdir3
    ├── song5.ogg
    └── song6.aac

5 directories, 7 files

elapsed 0.006s

jialinwu in go/src/m3ugenerator via 🐹 v1.23.2 via 🐍 v3.13.1
❯ ./m3ugenerator -dir ./music
Generating M3U file for music/subdir1 file count 2 m3uFilePath music/subdir1/subdir1.m3u...
M3U file generated: music/subdir1/subdir1.m3u
Generating M3U file for music/subdir2 file count 2 m3uFilePath music/subdir2/subdir2.m3u...
M3U file generated: music/subdir2/subdir2.m3u
Generating M3U file for music/subdir3 file count 2 m3uFilePath music/subdir3/subdir3.m3u...
M3U file generated: music/subdir3/subdir3.m3u
M3U files generated successfully.

elapsed 0.010s

jialinwu in go/src/m3ugenerator via 🐹 v1.23.2 via 🐍 v3.13.1
❯ tree music/
music/
├── music.m3u
├── subdir1
│   ├── song1.mp3
│   ├── song2.wav
│   └── subdir1.m3u
├── subdir2
│   ├── nested_subdir  #notice subdir deeper than 2 levels will be ignored which won't generate m3u file deeper dirs
│   │   └── song4.m4a
│   ├── song3.flac
│   └── subdir2.m3u
└── subdir3
    ├── song5.ogg
    ├── song6.aac
    └── subdir3.m3u
```
