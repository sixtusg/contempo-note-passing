[![Go Report Card](https://goreportcard.com/badge/github.com/sixtusg/contempo-note-passing)](https://goreportcard.com/report/github.com/sixtusg/contempo-note-passing)
[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)
[![GitHub license](https://img.shields.io/github/license/sixtusg/contempo-note-passing)](https://github.com/sixtusg/contempo-note-passing/blob/main/LICENSE)

# contempo-note-passing
Simple TCP chat server+client written in Go. Note that individual usage guides are in the client and server's respective directories.

The server receives a message from the client, and then broadcasts that message to every other client.

## Usage
Run the binary provided in the release in terminal, (Windows: `contempo-note-passing.exe`, Unix: `./contempo-note-passing`) and enter server host, port and type in the provided prompts.

## Goals
Goals are subject to change.

* Add basic functionality [X]
* Make the code look good [five years]
* Add nicknames [ ]
* Add formatting [ ]
  * Colours in nicknames [ ]
  * Timestamps with messages [ ]
* Encrypted messages; clients encrypt + decrypt messages with a key [ ]
* Signed messages [ ]
* Rooms; server handles "mini-servers" separately [ ]
* Customisable config files to automatically adjust nicknames, formatting and connections
