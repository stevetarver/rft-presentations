#!/bin/sh -e
#
# Run hello.go
#

go run hello_svg.go > hello_svg.svg

# Convert to png for easy display
# If you don't have ImageMagick:
# Note: librsvg improves conversion
#  brew install imagemagick --with-librsvg
#
# Reduce blurriness: -density 1200 (set this high)???
# Preserve transparency: -background none
# To keep image ratio, you also could specify only one dimension like -resize 200 for width or -resize x200 for height.
convert -density 1200 -background none hello_svg.svg hello_svg.png


convert -background none hello_svg.svg hello_svg.png
