#!/bin/sh -e
#
# Run hello.go
#

go run hello_svg.go > hello_svg.svg

# Convert to png for easy display
# If you don't have ImageMagick:
# Note: librsvg improves conversion
#  brew install imagemagick --with-librsvg
convert -background none hello_svg.svg hello_svg.png
