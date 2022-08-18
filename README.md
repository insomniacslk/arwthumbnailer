# arwthumbnailer

Thumbnailer for RAW files. Specifically for Sony ARW, but should work for other
formats that embed thumbnails or preview images in the RAW.

To use as a thumbnailer in GNOME's Nautilus (file viewer), copy
`arw.thumbnailer` under `/usr/share/thumbnailers`, and restart Nautilus with
`nautilus -q`. At the next Nautilus run, the thumbnailer will be used.

## Troubleshooting

If no thumbnails appear, try removing the thumbnail cache under
`~/.cache/thumbnails` and retry.
