# now

## Description

**now** is a simple command line utility for printing the current time in a variety of formats.  **now** also supports time deltas.

**now** is built in [Go](https://golang.org/). **now** uses the [time](https://pkg.go.dev/time) package to format the current time.

## Usage

Below is the usage for the `now` command.

```text
Now is a simple command line utility for printing the current time in a variety of formats.  Now also supports time deltas.  Now is built in Go and uses the time package to format the current time.

The value for the format flag can be in the Go time format or one of the following constants from the Go time package: ANSIC, RFC822, RFC822Z, RFC850, RFC1123, RFC1123Z, RFC3339, RFC3339Nano, Kitchen, Stamp, StampMilli, StampMicro, and StampNano.

Usage:
  now [flags]

Flags:
  -d, --delta string       the time delta from the current time in the go duration format (default "0s")
  -e, --epoch              print the UNIX Epoch time, which is the duration since midnight on January 1, 1970 UTC.
  -f, --format string      a constant or a verbose time format (default "RFC3339Nano")
  -h, --help               help for now
  -p, --precision string   the precision to use for printing the UNIX Epoch time: seconds (s), milliseconds (ms), or nanoseconds (ns) (default "s")
  -z, --time-zone string   the time zone: either UTC, Local, or name in the IANA Time Zone database (defaults to local time zone)
  -v, --version            print the version
```

Use the `--epoch` and `--precision` flags to print the UNIX Epoch time.

Use the `--format` flag to print the current time using a custom go format string.  The following constants from the Go [time](https://pkg.go.dev/time?tab=doc#pkg-constants) package can also be used: `ANSIC`, `RFC822`, `RFC822Z`, `RFC850`, `RFC1123`, `RFC1123Z`, `RFC3339`, `RFC3339Nano`, `Kitchen`, `Stamp`, `StampMilli`, `StampMicro`, and `StampNano`.

Use the `--delta` flag to print the `{current time} + {delta}`.  The value for `--delta` is parsed using the [ParseDuration](https://pkg.go.dev/time?tab=doc#ParseDuration) function and supports the following valid units: "ns", "us" (or "µs"), "ms", "s", "m", "h".

## Examples

### UNIX Epoch Time in Seconds

```shell
now -e -p s
```

### UNIX Epoch Time in Milliseconds

```shell
now -e -p ms
```

### UNIX Epoch Time in Microseconds

```shell
now -e -p us
```

### UNIX Epoch Time in Nanoseconds

```shell
now -e -p ns
```

### Custom Go Format

```shell
now -f 2006-01-02
```

### 2 Hours ago

```shell
now -f RFC3339 -d '-2h'
```

### 1 Minute from now

```shell
now -f RFC3339 -d '1m'
```

### Time Zone

```shell
now -f RFC3339 -z America/Los_Angeles
```

### UTC

```shell
now -f RFC3339 -z UTC
```

## Building

**now** is written in pure Go, so the only dependency needed to compile the server is [Go](https://golang.org/).  Go can be downloaded from <https://golang.org/dl/>.

This project uses [direnv](https://direnv.net/) to manage environment variables and automatically adding the `bin` and `scripts` folder to the path.  Install direnv and hook it into your shell.  The use of `direnv` is optional as you can always call `now` directly with `bin/now`.

If using `macOS`, follow the `macOS` instructions below.

To build a binary for your local operating system you can use `make bin/now`.  To build for a release, you can use `make build_release`.  Additionally, you can call `go build` directly to support specific use cases.

### macOS

You can install `go` on macOS using homebrew with `brew install go`.

To install `direnv` on `macOS` use `brew install direnv`.  If using bash, then add `eval \"$(direnv hook bash)\"` to the `~/.bash_profile` file .  If using zsh, then add `eval \"$(direnv hook zsh)\"` to the `~/.zshrc` file.

## Contributing

We'd love to have your contributions!  Please see [CONTRIBUTING.md](CONTRIBUTING.md) for more info.

## Security

Please see [SECURITY.md](SECURITY.md) for more info.

## License

This project constitutes a work of the United States Government and is not subject to domestic copyright protection under 17 USC § 105.  However, because the project utilizes code licensed from contributors and other third parties, it therefore is licensed under the MIT License.  See LICENSE file for more information.
