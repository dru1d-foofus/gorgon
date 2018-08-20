# GOrgon
A port of jmk's Medusa written in GO

## !!DISCLAIMER!!

This project literally started during DEFCON 26 as a way to force myself into learning GO and how to be a better programmer in general. I am not an accomplished developer so if there are better, cleaner ways of writing these things. Let me know. I appreciate all forms of harsh and constructive criticism.

## About

GOrgon is a modular and cross-platform bruteforce tool written in GO. The multi-platform support of GO will allow this tool to be compiled for Windows, Mac, and Linux; this will also hopefully cut down on the many dependency issues that plague Medusa in its current form.

## Installation

```
go get github.com/dru1d-foofus/gorgon
cd $GOPATH/src/github.com/dru1d-foofus/gorgon/cmd/gorgon
go build
```

## Basics
```
./gorgon -h
```
Everything else should be self-explanatory.

![start menu](https://i.imgur.com/a7JYzYx.png)

## To-Do
- [x] Finish SSH bruteforce module; this mainly the plaintext argument
- [ ] Add output option
- [ ] Add parallelism/concurrency
- [ ] SMB module with SMBv(1/2/3) support and pth
- [ ] RDP module with pth support
- [ ] Replicating all existing and relevant modules that currently exist with Medusa
