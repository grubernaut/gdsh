gDSH
====

Dancer's Shell - In Go

*Disclaimer:*
Obviously, dsh, originally written by Junichi Uekawa <dancer@debian.org> in C is light-years faster and more feature-complete than gDSH. This was/is a purely academic venture into both reading C code, and writing something somewhat useful in Golang. You can download the source for dsh [here](https://www.netfort.gr.jp/~dancer/software/downloads/#dsh)


## What is DSH?

I love DSH. There's always a use case for creating a program that can easily be replaced by a few lines of bash. However, I often use DSH daily in operations work, as it just makes life easier. Hopefully, even if you don't end up using gDSH, you'll discover DSH and use it. It's honestly an amazing piece of software.

From DSH's manpage:

> dsh executes a command remotely on several different machines at the same time. A utility to effectively do: 

```bash
for a in $(seq 1 10); do rsh $a command; done
```

gDSH aims to be as close to feature complete to DSH as possible.

## Usage

```
gdsh [options] "commandline"
```

## Currently Not Working (Not Implemented Yet) - TODO

Here's a list of things that aren't implemented yet in gDSH, that are implemented in DSH. Feel free to submit a PR, or they may be tackled as $free_time allows. 

* Adding in a single machinename via `-m` from the command line
* Adding in all machines found in `/usr/local/etc/gdsh/machines.list` to the list
* Adding in a group of machines found in `/usr/local/etc/gdsh/group/<groupname>` to the list
* Opening a terminal on each machine, and replicating input across each. (Honestly I have no idea how this even works in DSH)
* NumTopology is currently locked to 1. DSH allows spawning multiple DSH processes on remote hosts.
* Setting a forklimit. This could be useful if you're using gDSH to create an obnoxious number of remote connections.
* `--wait-shell` and `--concurrent-shell`, `-w` and `-c` are currently meaningless. gDSH is currently concurrent by default. Sorry
* TESTS, TESTS, TESTS - There aren't any. I should write some.

## Installation

gDSH looks amazing! How do I use it? 

1. Install Go and configure $GOPATH
2. `go get -u github.com/grubernaut/gdsh`

## Compilation

I should write a Makefile, but currently things are fairly small and simple.
Dependencies are vendored via `govendor`, and there aren't many of them. 

```
go build      # build the binary
go install    # build and install the binary
```

## Go API

I've aimed to abstract as much of the gDSH runtime as possible into the `dsh` package. If you wish to write a CLI tool that implements the gDSH API, I've attempted to make this a possibility. Reading through `run.go` should give enough insight into how to do this. Hopefully...

