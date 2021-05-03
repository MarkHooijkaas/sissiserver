# SissiServer

Simple SSI Server: A very simple webserver supporting SSI (server side includes)

## Getting Started

- Copy the sissiserver.exe to a directory
- Make a subdirectory called site with html files and other files
- Start the sissiserver executable
- Visit http:///localhost:8090 in your browser

You can run the executable with the following two optional command lines options
- `--port=8090` determine the port to listen to (default is 8090)
- `--dir=html` the directory that contains static html files (defuult is html)
this could be as follows:
```
sissiserver --port=81810 --dir=/var/www/html
```

## Purpose
This is a very simple and lightweight webserver.
It consists of just one executable file, nothing else.
It is meant for simple development and testing use, not for production use.

The intended scenario is to have a static website which uses server side includes.
These include-files can be blocks of text that an user can edit.
He can then preview his website on his browser using sissiserver.

## Caution
SissiServer is *NOT* intended for production use.
It has no security features and is not designed for heavy load.
For production purposes you can use a webserver such as NGINX which supports the same SSI syntax.

## Limitations
The only ssi directive that is supported is the "include virtual" syntax, as follows:
```html
<!--#include virtual="some-file.html" -->
```
See the test directory for an example.

## Building
The source code consist of one source file sissi.go
