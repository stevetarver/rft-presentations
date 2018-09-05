---?image=svgo/pitch/images/go-pencil/gopherswim.jpg&size=contain&color=white&position=right

@title[Intro]

<br/>

**_Baby gophers:_**
<br/>
**_care and feeding_**

+++

Follow along with this pitch at: 

> https://gitpitch.com/stevetarver/rft-presentations/master?p=svgo

---

### What are SVGs

> Scalable Vector Graphics (SVG) is an XML-based vector image format for two-dimensional graphics with support for interactivity and animation. The SVG specification is an open standard developed by the World Wide Web Consortium (W3C) since 1999. SVG images and their behaviors are defined in XML text files.

Why use them

Scaled Vector Graphics
Scale to any size and look great

Svg tools

---


@snap[north-east]
<img style="border:0px; box-shadow: 0px 0px 0px rgba(0, 0, 0, .0);" height="60px" src="svgo/pitch/images/pale-violet-square.png">
@snapend

SVGs are just XML

```xml
<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg" version="1.1" viewBox="61 61 142 142" width="142pt" height="142pt"
     xmlns:dc="http://purl.org/dc/elements/1.1/">
    <metadata>Produced by OmniGraffle 6.6.2
        <dc:date>2018-09-05 19:20:43 +0000</dc:date>
    </metadata>
    <defs/>
    <g stroke="none" stroke-opacity="1" stroke-dasharray="none" fill="none" fill-opacity="1">
        <title>Canvas 1</title>
        <g>
            <title>Layer 1</title>
            <rect x="72" y="72" width="120" height="120" fill="#9662d0"/>
            <rect x="72" y="72" width="120" height="120" stroke="black" stroke-linecap="round" stroke-linejoin="round"
                  stroke-width="1"/>
        </g>
    </g>
</svg>
```

---

### Why Go?

- **Minutes to learn**, sophisticated enough to conquer the world
- **Business**: small, fast, deploy anywhere, rapid product development
- **Developer joy**: easy, fun, make stuff quickly
- **Jobs**? If you can spell "Go", you can get a job on the front range

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

```
brew install httpie 
```

---

### Reverse engineering an SVG

Mock up on favorite tool, convert to code

Identify live data sections

---

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
