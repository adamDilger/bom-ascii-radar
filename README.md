# bom-ascii-radar

This is a ascii-based displayer for the Bureau of Meteorology's radar images.
It's primary use is for a widget in a tmux driven dashboard.

![demo screenshot](/docs/bom-ascii-radar-demo.gif)

## Dependencies

This project uses the following libraries:

- [imagemagick](https://imagemagick.org/index.php)
  - to merge the radar images with the background
- [https://github.com/TheZoraiz/ascii-image-converter](ascii-image-converter)
  - to convert the merged image in to ascii

## Options

- `-productCode` The product code to fetch images for (IDR763)
- `-cacheDir` The directory to store cached images (/tmp/radar)
- `-timezone` The timezone to use for the radar image timestamps
  (Australia/Hobart)
- `-backgroundColor` The background color to use for the radar image" (#808080)
- `-debug` Output debugs to stderr (e.g. `go run . -debug 2> /tmp/debug.log`)
