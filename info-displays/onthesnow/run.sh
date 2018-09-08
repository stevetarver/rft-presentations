#!/bin/sh -e
#
# Build one of the numbered projects
#
# NOTES:
#
# CAVEATS:
#
# TODO:
if [ "$(uname -s)" = "Darwin" ]; then
    # If called through a symlink, this will point to the symlink
    THIS_SCRIPT_DIR="$( cd "$( dirname "${0}" )" && pwd )"
else
    THIS_SCRIPT_DIR=$(dirname $(readlink -f "${0}"))
fi
(
    # Ensure we are in this directory for relative paths
    cd ${THIS_SCRIPT_DIR}/${1}

    go run wallpaper.go > wallpaper.svg

    #convert -background none wallpaper.svg wallpaper.png
)
