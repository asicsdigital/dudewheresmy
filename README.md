# Dude, Where's My

`dudewheresmy` is a portable tool for obtaining environment- or context-specific information at runtime.

## Installation

You need [Go](http://golang.org) to build this tool.

```sh
$ go get github.com/asicsdigital/dudewheresmy
```

You should now have an executable at `$GOPATH/bin/dudewheresmy`.

## Usage

```sh
$ dudewheresmy
NAME:
   dudewheresmy - find things you're looking for

USAGE:
   dudewheresmy [global options] command [command options] [arguments...]

VERSION:
   v0.1.0

COMMANDS:
     hostip, i, ip, host  find IP address of process host
     help, h              Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

### `hostip`

From userspace on a POSIX system, `hostip` returns localhost:

```sh
$ dudewheresmy hostip
127.0.0.1
```

From inside a Docker container, `hostip` returns a lookup of `host.docker.internal.`:

```sh
$ docker run --rm -it asicsdigital/dudewheresmy:latest hostip
192.168.65.2
```

From an [EC2](https://aws.amazon.com/ec2/) instance, `hostip` returns a lookup of `local-ipv4` from the metadata service:

```sh
$ dudewheresmy hostip
10.1.20.222
```

## Contributing

This tool is built with [urfave/cli](https://github.com/urfave/cli).  To add another subcommand:

1. `mkdir -p commands/<YOUR SUBCOMMAND>`
2. `$EDITOR commands/<YOUR SUBCOMMAND>/<YOUR SUBCOMMAND>.go`
3. Declare a new package namespace for your subcommand.
4. Define a `Command()` function that returns a `cli.Command` struct
5. Import `github.com/asicsdigital/dudewheresmy/commands/<YOUR SUBCOMMAND>` from main.go
