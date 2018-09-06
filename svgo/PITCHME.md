---?image=svgo/pitch/images/go-pencil/gopherswim.jpg&size=contain&color=white&position=right

@title[Intro]

<br/>

**_Automating Digital Asset_**
<br/>
**_ Workflows... with Go_**

+++

### ***What we'll do***

* Introduce Go
* Code SVGs
* Think through opportunities for automating your workflow


+++

Follow along with this pitch at: 

> https://gitpitch.com/stevetarver/rft-presentations/master?p=svgo

---

# One hash #
## Two hashses ##
### Three hashses ###
#### Four hashses ####
##### Five hashses #####
###### Six hashes ######

+++

Plain text _italics_ **BOLD** **_BOLD italics_**

@ul

- bullet 1
- bullet 2

@ulend

@ol

1. number 1
1. number 2

@olend

> quote

---


### ***What are SVGs***

> Scalable Vector Graphics (SVG) is an XML-based vector image format for two-dimensional graphics with support for interactivity and animation.

+++

### ***Why use them?***

* Display perfectly at any scale
* All elements scale together
* Export to `png` at target resolutions for devices that can't render `svg`

+++

@snap[north-east]
<img style="border:0px; box-shadow: 0px 0px 0px rgba(0, 0, 0, .0);" height="60px" src="svgo/pitch/images/pale-violet-square.png">
@snapend

#### SVGs are just XML

```xml
<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" 
    "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg" version="1.1" 
    viewBox="61 61 142 142" width="142pt" height="142pt"
     xmlns:dc="http://purl.org/dc/elements/1.1/">
    <metadata>
        Produced by OmniGraffle 6.6.2
        <dc:date>2018-09-05 19:20:43 +0000</dc:date>
    </metadata>
    <defs/>
    <g stroke="none" stroke-opacity="1" stroke-dasharray="none" 
        fill="none" fill-opacity="1">
        <title>Canvas 1</title>
        <g>
            <title>Layer 1</title>
            <rect x="72" y="72" width="120" height="120" fill="#9662d0"/>
            <rect x="72" y="72" width="120" height="120" stroke="black" 
                stroke-linecap="round" stroke-linejoin="round" stroke-width="1"/>
        </g>
    </g>
</svg>
```

---

### ***When _can_ you use them?***

Classically:

@ul

- Web page icons: [Font Awesome](https://fontawesome.com/?utm_source=v4_homepage&utm_medium=display&utm_campaign=fa5_released&utm_content=banner)
- Custon icon sets: [CTL Monocle](https://www.ctl.io/developers/blog/post/monocle-pattern-library)

SVGs are great for simple drawings, but you can do more...

+++

You can create charts:

<img style="border:0px; box-shadow: 0px 0px 0px rgba(0, 0, 0, .0);" height="500px" src="svgo/pitch/images/starks_examples/stocks.png">

+++

Imagine camera snapshots from the four mountains:

<img style="border:0px; box-shadow: 0px 0px 0px rgba(0, 0, 0, .0);" height="450px" src="svgo/pitch/images/starks_examples/jtree_thumbs.png">

+++

And then putting that on a large format TV:

<img style="border:0px; box-shadow: 0px 0px 0px rgba(0, 0, 0, .0);" height="450px" src="svgo/pitch/images/starks_examples/twh-20180417.png">

---

Svg tools

Strategy: generate the svg and then pngs at a variety of resolutions

Design Asset Workflow

* Design
* Draw
* Export
* Publish
* Done

---

How does this change if you want to show live data?

* Design
* Fetch data
* Draw
* Export
* Publish
* Fetch data... at some period and republish

If done manually, time sink
If you automate, effortless - increase your bill rate

---


Many automation tools, we're going to look at a new one, because it's different, and really easy

### Why Go?

@ul

- **Minutes to learn**, sophisticated enough to conquer the world
- **Business**: small, fast, deploy anywhere, rapid product development
- **Developer joy**: easy, fun, make stuff quickly
- **Tooling exists**: svgo is just a simple way to create these xml lines


Recording your drawing in code lets you version control it, revert to any previous stage, have many in-progress pieces.

@ulend

@snap[south-west]
@fa[git fa-2x]
@snapend

@snap[south]
<img style="border:0px; box-shadow: 0px 0px 0px rgba(0, 0, 0, .0);"  height="50" src="svgo/pitch/images/Go-Logo/PNG/Go-Logo_Aqua_sm.png">
@snapend

@snap[south-east]
@fa[github fa-2x]
@snapend

---

### Install Go

1. Go to the install page: [https://golang.org/doc/install](https://golang.org/doc/install)
1. Click the "Download Go" button.
1. Run the installer.
1. Check the version on the command line.

```
ᐅ go version
go version go1.11 darwin/amd64
```

### Troubleshooting

* Restart the shell to pickup paths changes.
* Check paths setup properly.

---

### Create a new project

```
ᐅ go mod init github.com/stevetarver/rft-presentations/svgo
go: creating new go.mod: module github.com/stevetarver/rft-presentations/svgo

ᐅ cat go.mod
module github.com/stevetarver/rft-presentations/svgo
```

---

### Create your first app

Create file `hello.go`:

```go
package main

func main() {
    println("Hello baby gophers!")
}
```

Run it:

```go
ᐅ go run hello.go
Hello baby gophers!
```

---

### svgo overview

What are SVGs
Why are they cool
Where can they be used
Show examples

What does an SVG file look like?

---

## Hello SVG

```go
 ᐅ go get github.com/ajstarks/svgo
go: finding github.com/ajstarks/svgo latest
go: downloading github.com/ajstarks/svgo v0.0.0-20180830174826-7338bd80e790
ᐅ cat go.mod
module github.com/stevetarver/rft-presentations/svgo

require github.com/ajstarks/svgo v0.0.0-20180830174826-7338bd80e790 // indirect
```

show code, run, view

---

Project ideas

Install https://github.com/ajstarks/svgo

* GitHub activity
* Custom screen display

install svgo: `go get github.com/ajstarks/svgo`


---

### Using live data

How to fetch

curl, wget, httpie

Focus on live data - this is where it shines

```
brew install httpie 
```

---

### Reverse engineering an SVG

Mock up on favorite tool, convert to code

Identify live data sections

---

### Use on website

embed

simple webpage that references external url

____


### Install as mac wallpaper

qlmanage -t -s 1000 -o . picture.svg 
It will produce picture.svg.png that is 1000 pixels wide.
May only produce squares.

rsvg-convert
brew install librsvg and is used like this:

rsvg-convert -h 32 icon.svg > icon-32.png


ImageMagick is an extremely versatile command-line image editor, which would probably rival Photoshop if it had, you know, a GUI. But who needs those anyways. :P

Something like the following would convert a .svg to .png, after installation:

$ convert picture.svg picture.png
The original .svg isn't deleted.

osascript -e 'tell application "Finder" to set desktop picture to POSIX file "/Users/starver/code/rft/rft-presentations/svgo/pitch/images/pale-violet-square.png"'

could do the same on iphone


___

### Install on appletv

Strategies:

Many large TVs have a store mode for just this type of thing - my LG for example. This retail mode may jack with home user mode.
You can set up a local media server and many tvs can connect to it (webOS)
Use an AppleTV
* Use screen saver: Set to a photo or album and use as a screen saver
* Use Photos: Create a photos album holding a slide show and view that on the TV.
* Use Home Sharing: 
* Get an app to show content from dropbox, google drive, etc.

Create an app that loads slides from a remote hosted site.


___

# References

## AJ Starks

* [svgo Github page: https://github.com/ajstarks/svgo](https://github.com/ajstarks/svgo)
* [svgo info displays: https://github.com/ajstarks/go-info-displays](https://github.com/ajstarks/go-info-displays)
* [Deck Github page: https://github.com/ajstarks/deck](https://github.com/ajstarks/deck)

Tons of examples in go-info-displays


---

# TODO

---?image=pitch/images/go-pencil/gopherhat.jpg
---?image=pitch/images/go-pencil/gopherhelmet.jpg
---?image=pitch/images/go-pencil/gophermega.jpg
---?image=pitch/images/go-pencil/gopherrunning.jpg
---?image=pitch/images/go-pencil/gopherswim.jpg
---?image=pitch/images/go-pencil/gopherswrench.jpg
