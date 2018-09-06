#!/bin/sh -e
#
# Run our pitch on a local python http server
#

#PITCH_ZIP_NAME='PITCHME.zip'
#
#cd pitch
#wget -O "${PITCH_ZIP_NAME}" https://gitpitch.com/pitchme/offline/github/stevetarver/rft-presentations/master/simple/PITCHME.zip?p=info-displays
#unzip "${PITCH_ZIP_NAME}"

cd pitch/PITCHME
python3 -m http.server
