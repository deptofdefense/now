// =================================================================
//
// Work of the U.S. Department of Defense, Defense Digital Service.
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	NowVersion = "1.0.0"
)

const (
	flagEpoch     = "epoch"
	flagFormat    = "format"
	flagPrecision = "precision"
	flagDelta     = "delta"
	flagTimeZone  = "time-zone"
	flagVersion   = "version"
)

func initFlags(flag *pflag.FlagSet) {
	flag.BoolP(flagEpoch, "e", false, "print the UNIX Epoch time, which is the duration since midnight on January 1, 1970 UTC.")
	flag.StringP(flagFormat, "f", "RFC3339Nano", "a constant or a verbose time format")
	flag.StringP(flagPrecision, "p", "s", "the precision to use for printing the UNIX Epoch time: seconds (s), milliseconds (ms), or nanoseconds (ns)")
	flag.StringP(flagDelta, "d", "0s", "the time delta from the current time in the go duration format")
	flag.StringP(flagTimeZone, "z", "", "the time zone: either UTC, Local, or name in the IANA Time Zone database (defaults to local time zone)")
	flag.BoolP(flagVersion, "v", false, "print the version")
}

func initViper(cmd *cobra.Command) (*viper.Viper, error) {
	v := viper.New()
	err := v.BindPFlags(cmd.Flags())
	if err != nil {
		return v, fmt.Errorf("error binding flag set to viper: %w", err)
	}
	v.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	v.AutomaticEnv() // set environment variables to overwrite config
	return v, nil
}

func checkConfig(v *viper.Viper) error {
	epoch := v.GetBool(flagEpoch)
	if epoch {
		precision := v.GetString(flagPrecision)
		if len(precision) == 0 {
			return fmt.Errorf("precision is missing")
		}
	} else {
		format := v.GetString(flagFormat)
		if len(format) == 0 {
			return fmt.Errorf("unix epoch or format is required")
		}
	}
	delta := v.GetString(flagDelta)
	if len(delta) > 0 {
		if _, err := time.ParseDuration(delta); err != nil {
			return fmt.Errorf("error parsing delta %q: %w", delta, err)
		}
	}
	return nil
}

func formatDate(d time.Time, format string) (int, error) {

	switch strings.ToLower(format) {
	case "ansic":
		return fmt.Fprintf(os.Stdout, "%s\n", d.Format(time.ANSIC))
	case "rfc822":
		return fmt.Fprintf(os.Stdout, "%s\n", d.Format(time.RFC822))
	case "rfc822z":
		return fmt.Fprintf(os.Stdout, "%s\n", d.Format(time.RFC822Z))
	case "rfc850":
		return fmt.Fprintf(os.Stdout, "%s\n", d.Format(time.RFC850))
	case "rfc1123":
		return fmt.Fprintf(os.Stdout, "%s\n", d.Format(time.RFC1123))
	case "rfc1123z":
		return fmt.Fprintf(os.Stdout, "%s\n", d.Format(time.RFC1123Z))
	case "rfc3339":
		return fmt.Fprintf(os.Stdout, "%s\n", d.Format(time.RFC3339))
	case "rfc3339nano":
		return fmt.Fprintf(os.Stdout, "%s\n", d.Format(time.RFC3339Nano))
	case "kitchen":
		return fmt.Fprintf(os.Stdout, "%s\n", d.Format(time.Kitchen))
	case "stamp":
		return fmt.Fprintf(os.Stdout, "%s\n", d.Format(time.Stamp))
	case "stampmilli":
		return fmt.Fprintf(os.Stdout, "%s\n", d.Format(time.StampMilli))
	case "stampmicro":
		return fmt.Fprintf(os.Stdout, "%s\n", d.Format(time.StampMicro))
	case "stampnano":
		return fmt.Fprintf(os.Stdout, "%s\n", d.Format(time.StampNano))
	}

	return fmt.Fprintf(os.Stdout, "%s\n", d.Format(format))
}

func main() {

	now := time.Now()

	rootCommand := &cobra.Command{
		Use:                   `now [flags]`,
		DisableFlagsInUseLine: true,
		Short: `Now is a simple command line utility for printing the current time in a variety of formats.  Now also supports time deltas.  Now is built in Go and uses the time package to format the current time.

The value for the format flag can be in the Go time format or one of the following constants from the Go time package: ANSIC, RFC822, RFC822Z, RFC850, RFC1123, RFC1123Z, RFC3339, RFC3339Nano, Kitchen, Stamp, StampMilli, StampMicro, and StampNano.
		`,
		SilenceErrors: true,
		SilenceUsage:  true,
		RunE: func(cmd *cobra.Command, args []string) error {

			v, err := initViper(cmd)
			if err != nil {
				return fmt.Errorf("error initializing viper: %w", err)
			}

			if len(args) > 0 {
				return cmd.Usage()
			}

			if errConfig := checkConfig(v); errConfig != nil {
				return errConfig
			}

			if v.GetBool(flagVersion) {
				fmt.Println(NowVersion)
				return nil
			}

			tz := v.GetString(flagTimeZone)

			d := now.Add(v.GetDuration(flagDelta))

			if len(tz) > 0 {
				location, err := time.LoadLocation(tz)
				if err != nil {
					return fmt.Errorf("error parsing time zone %q: %w", tz, err)
				}
				d = d.In(location)
			}

			if v.GetBool(flagEpoch) {
				precision := v.GetString(flagPrecision)
				switch precision {
				case "seconds", "second", "s":
					_, _ = fmt.Fprintf(os.Stdout, "%d\n", d.Unix())
					return nil
				case "milliseconds", "millisecond", "ms":
					_, _ = fmt.Fprintf(os.Stdout, "%d\n", d.UnixNano()/1000000)
					return nil
				case "nanoseconds", "nanosecond", "ns":
					_, _ = fmt.Fprintf(os.Stdout, "%d\n", d.UnixNano())
					return nil
				}
				return fmt.Errorf("unknown precision (%q) for unix epoch time", precision)
			}

			_, _ = formatDate(d, v.GetString(flagFormat))

			return nil
		},
	}
	initFlags(rootCommand.Flags())

	if err := rootCommand.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "now: "+err.Error())
		_, _ = fmt.Fprintln(os.Stderr, "Try now --help for more information.")
		os.Exit(1)
	}
}
