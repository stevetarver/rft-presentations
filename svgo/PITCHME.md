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

**_Why Go?_**

- **Minutes to learn**, sophisticated enough to conquer the world
- **Business**: small, fast, deploy anywhere, rapid product development
- **Developer joy**: easy, fun, make stuff quickly
- **Jobs**? If you can spell "Go", you can get a job on the front range

@snap[south-east]
@fa[github fa-3x]
@snapend

@snap[south]
![golang](svgo/pitch/images/Go-Logo/PNG/Go-Logo_Aqua.png)
@snapend

@snap[south-west]
@fa[bitbucket fa-3x]
@snapend

---

# Go install & setup

1. Go to the install page: [https://golang.org/doc/install](https://golang.org/doc/install)
1. Click the "Download Go" button.
1. Run the installer.

```
ᐅ go version
go version go1.11 darwin/amd64
```

### Troubleshooting

* Did you restart your shell so paths changes can take effect?
* Did the installer setup the paths properly?

---

# Go project initialization

```
ᐅ go mod init github.com/stevetarver/rft-presentations/svgo
go: creating new go.mod: module github.com/stevetarver/rft-presentations/svgo
ᐅ cat go.mod
module github.com/stevetarver/rft-presentations/svgo
```

---

# Hello world!

TODO: fetch this from the repo

Create file `hello.go`:

```go
package main

func main() {
	println("Hello baby gophers!")
}
```

Run it

```go
ᐅ go run hello.go
Hello baby gophers!
```

This is used for testing changes frequently - and maybe all that you need.


## basic commands

go install hello.go

---

## svgo overview

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
