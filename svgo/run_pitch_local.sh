#!/bin/sh -e
#
# Run our pitch on a local python http server
#
cd pitch/PITCHME
python3 -m http.server
