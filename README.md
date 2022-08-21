# arwthumbnailer

Thumbnailer for RAW files. Specifically for Sony ARW, but should work for other
formats that embed thumbnails or preview images in the RAW.

To use as a thumbnailer in GNOME's Nautilus (file viewer), copy
`arw.thumbnailer` under `/usr/share/thumbnailers`, and restart Nautilus with
`nautilus -q`. At the next Nautilus run, the thumbnailer will be used.

## Troubleshooting

If no thumbnails appear, try removing the thumbnail cache under
`~/.cache/thumbnails` and retry.


## Alternatives

* shell script: instead of this program you can use a shell script as
  thumbnailer:

```
#!/bin/bash
set -exu

input=$1
output=$2
size=$3

# Alternatives:
# -PreviewImage (large image)
# -ThumbnailImage (small image)
# -JpgFromRaw (large image)
exiftool -b -ThumbnailImage "${input}" > "${output}.tmp"
convert -resize "${size}" "${output}.tmp" "${output}"
rm -f "${output}.tmp"
```

* add more mime-types to gdk-pixbuf-thumbnainer: add the following mime-types to
  `/usr/share/thumbnailers/gdk-pixbuf-thumbnailer.thumbnailer`:

```
image/x-mef;image/x-sony-srf;image/x-pef;image/x-srf;image/x-3fr;image/x-cr2;image/x-arw;image/x-fuji-raf;image/x-x3f;image/x-eip;image/x-dng;image/x-crw;image/x-pxn;image/x-dcr;image/x-samsung-srw;image/x-adobe-dng;image/x-orf;image/x-drf;image/x-cap;image/x-sony-arw;image/x-iiq;image/x-raf;image/x-rw2;image/x-dcs;image/x-sigma-x3f;image/x-kdc;image/x-pentax-pef;image/x-dcraw;image/x-r3d;image/x-panasonic-raw;image/x-rwl;image/x-nrw;image/x-canon-cr2;image/x-bay;image/x-mrw;image/x-canon-crw;image/x-olympus-orf;image/x-k25;image/x-rwz;image/x-erf;image/x-raw;image/x-mos;image/x-nikon-nef;image/x-nef;image/x-minolta-mrw;image/x-ptx;image/x-sr2;image/x-fff;image/x-sony-sr2;image/x-panasonic-raw2
```
