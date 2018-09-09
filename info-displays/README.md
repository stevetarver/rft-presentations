This project supports the PITCHME.md at repo root. You can see the pitch at [https://gitpitch.com/stevetarver/rft-presentations/master?p=info-displays](https://gitpitch.com/stevetarver/rft-presentations/master?p=info-displays)

This is an incremental journey into Go using AJ Starks `svgo` package to create SVGs. The two larger examples have numbered directories that show evolution of each from scratch to 'finished' command line program.

Project directories:

* `hello_go` - a hello world in Go
* `hello_svg` - a hello world Go program to create a simple SVG
* `weather` - a desktop wallpaper showing the standard, time, weather and news headlines
* `onthesnow` - a desktop wallpaper showing area ski resort web cams - pick the one with the best snow.

_Note:_ this is a decidedly Mac implementation. You can comment out the `screen_resolution` function and use in `run.sh` and have full functionality (not tested). Windows? ... Happy to accept PRs. ;)

## Project setup

1. Go 1.11 installed (we use Go modules)
    * After cloning, cd to `info_displays` and `go get` the dependencies.
1. Gapplin installed (Mac App Store), or decent SVG viewer - any browser will work.
1. ImageMagick installed to convert SVG to PNG for desktop wallpapers
    ```
    brew install imagemagick --with-librsvg
    ```

## Create the wallpapers

Both `weather` and `onthesnow` have shell scripts to generate the incremental phases; the numbered directories. Provide the numbered directory as the single command line parameter.

```
./run.sh 12
```

In `run.sh`, there are provisions for tight development loops and generating wallpaper. Comment out the applicable sections.
