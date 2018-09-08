#!/bin/sh -e
#
# Build one of the numbered projects
#
# NOTES:
#
# CAVEATS:
#
if [ "$(uname -s)" = "Darwin" ]; then
    # If called through a symlink, this will point to the symlink
    THIS_SCRIPT_DIR="$( cd "$( dirname "${0}" )" && pwd )"
else
    # All linux should use this...
    THIS_SCRIPT_DIR=$(dirname $(readlink -f "${0}"))
fi
(
    # Ensure we are in the target directory for relative paths
    cd ${THIS_SCRIPT_DIR}/${1}

    # Run the go file to produce the wallpaper
    go run wallpaper.go > wallpaper.svg

    # Convert the SVG to PNG - macs can't show svg as wallpaper
    # The 'msg:' prefix downloads and includes the photos
#    convert msvg:wallpaper.svg wallpaper.png

    # Set the wallpaper as desktop background
#    FILE_PATH="$(pwd)/wallpaper.png"
#    osascript -e 'tell application "Finder" to set desktop picture to POSIX file "'${FILE_PATH}'"'
)
