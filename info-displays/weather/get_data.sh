#!/bin/sh -e
#
# Grab live data we will use in our SVG
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
    cd ${THIS_SCRIPT_DIR}

    # Create the data dir if it doesn't exist
    ! [ -d "data" ] && mkdir data

    # Get screen caps of 4 mountains

)
