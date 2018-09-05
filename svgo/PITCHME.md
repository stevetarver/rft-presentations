---?image=svgo/pitch/images/gopherswim.jpg&size=contain&color=white&position=right

@title[Intro]

<br/>

**_Baby gophers:_**
<br/>
**_care and feeding_**

+++

Follow along with this pitch at: 

> https://gitpitch.com/stevetarver/rft-presentations/master?p=svgo

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

---

## Project 1

Install https://github.com/ajstarks/svgo

* GitHub activity
* Custom screen display

---

# References

## AJ Starks

* [svgo Github page: https://github.com/ajstarks/svgo](https://github.com/ajstarks/svgo)
* [Deck Github page: https://github.com/ajstarks/deck](https://github.com/ajstarks/deck)


---

# TODO

---?image=pitch/images/gopherhat.jpg
---?image=pitch/images/gopherhelmet.jpg
---?image=pitch/images/gophermega.jpg
---?image=pitch/images/gopherrunning.jpg
---?image=pitch/images/gopherswim.jpg
---?image=pitch/images/gopherswrench.jpg
