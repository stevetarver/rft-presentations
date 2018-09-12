#!/bin/sh -e
#
# Run our pitch on a local python http server
#
# Manually fetch the offline version of the pitch using favorite theme
# extract in pitch/PITCHME
#

cd pitch/PITCHME
python3 -m http.server
