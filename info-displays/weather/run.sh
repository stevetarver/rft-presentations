#!/bin/sh -e
#
# Build one of the numbered projects
#
# NOTES:
#
# CAVEATS:
#

# Get the current screen resolution
#  'system_profiler SPDisplaysDataType' returns a line like:
#     Resolution: 2560 x 1440 (QHD/WQHD - Wide Quad High Definition)
#  that we can parse for width / height and provide to our wallpaper
screen_resolution() {
    TEXT=$(system_profiler SPDisplaysDataType | grep Resolution | awk '{$1=$1};1')
    WIDTH=$(echo ${TEXT} | cut -f2 -d' ')
    HEIGHT=$(echo ${TEXT} | cut -f4 -d' ')
    printf "%s ${WIDTH} %s ${HEIGHT}" '-w' '-h'
}

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

    # If in development, use this
#    go run wallpaper.go > wallpaper.svg
#    open wallpaper.svg

    # Otherwise, uncomment these:
    # Run the go file to produce the wallpaper
    go run wallpaper.go $(screen_resolution) > wallpaper.svg

    # Convert the SVG to PNG - macs can't show svg as wallpaper
    # The 'msvg:' prefix downloads and includes the photos
    convert msvg:wallpaper.svg wallpaper.png

    # Set the wallpaper as desktop background
    FILE_PATH="$(pwd)/wallpaper.png"
    osascript -e 'tell application "Finder" to set desktop picture to POSIX file "'${FILE_PATH}'"'
)
