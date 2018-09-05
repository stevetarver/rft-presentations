# Intro

TODO: not working - see this for more images

---?https://golang.org/doc/gopher/pencil/gopherswim.jpg
---?image=assets/img/gp-logo-tagline.svg

---?image=pitch/images/gopherhat.jpg
---?image=pitch/images/gopherhelmet.jpg
---?image=pitch/images/gophermega.jpg
---?image=pitch/images/gopherrunning.jpg
---?image=pitch/images/gopherswim.jpg
---?image=pitch/images/gopherswrench.jpg

inline image example

![SAMBA Deployment](https://onetapbeyond.github.io/resource/img/samba/new-samba-deploy.jpg)

top slide

+++

Follow along with [this pitch](https://gitpitch.com/stevetarver/rft-presentations/master?p=svgo).


---

# Why go

TODO: markdown fragments not working

@ul

- Plain text list item
- Rich **markdown** list *item*
- Link [within](https://gitpitch.com) list item

@ulend

- Like the board game, Go takes minutes to learn, but you can keep improving skills for a long time.
- Resume builder
- Small and fast for business use
- Fun - immediate feedback

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
