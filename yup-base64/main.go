package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"

	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/base64"
)

const (
	flagDecode        = "decode"
	flagIgnoreGarbage = "ignore-garbage"
	flagWrap          = "wrap"
)

func main() {
	app := &cli.App{
		Name:  "base64",
		Usage: "base64 encode/decode data and print to standard output",
		UsageText: `base64 [OPTIONS] [FILE]

   Base64 encode or decode FILE, or standard input, to standard output.
   With no FILE, or when FILE is -, read standard input.`,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    flagDecode,
				Aliases: []string{"d"},
				Usage:   "decode data",
			},
			&cli.BoolFlag{
				Name:    flagIgnoreGarbage,
				Aliases: []string{"i"},
				Usage:   "when decoding, ignore non-alphabet characters",
			},
			&cli.IntFlag{
				Name:    flagWrap,
				Aliases: []string{"w"},
				Usage:   "wrap encoded lines after COLS characters (default 76, 0 to disable)",
				Value:   76,
			},
		},
		Action: action,
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "base64: %v\n", err)
		os.Exit(1)
	}
}

func action(c *cli.Context) error {
	var params []any

	// Add file arguments (or none for stdin)
	for i := 0; i < c.NArg(); i++ {
		params = append(params, yup.File(c.Args().Get(i)))
	}

	// Add flags based on CLI options
	if c.Bool(flagDecode) {
		params = append(params, Decode)
	}
	if c.Bool(flagIgnoreGarbage) {
		params = append(params, IgnoreGarbage)
	}
	if c.IsSet(flagWrap) {
		wrapWidth := c.Int(flagWrap)
		if wrapWidth > 0 {
			params = append(params, Wrap, WrapWidth(wrapWidth))
		}
	}

	// Create and execute the base64 command
	cmd := Base64(params...)
	return yup.Run(cmd)
}
