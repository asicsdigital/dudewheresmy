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

This tool is built with [spf13/cobra](https://github.com/spf13/cobra).  To add another subcommand:

1. `go get github.com/spf13/cobra/cobra`
2. Ensure that `$GOPATH/bin` is in your PATH ([direnv](https://direnv.net/) will do this for you).
3. `cobra add <COMMAND_NAME>`
4. `$EDITOR cmd/<COMMAND_NAME>.go`
