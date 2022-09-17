package main

import (
	"fmt"
	"log"

	"github.com/zan8in/goflags"
)

type options struct {
	silent bool
	inputs goflags.StringSlice
	config string
	values goflags.RuntimeMap
}

func main() {
	opt := &options{}

	flagSet := goflags.NewFlagSet()
	flagSet.SetDescription("Test program to demonstrate goflags options")

	flagSet.BoolVar(&opt.silent, "silent", true, "show silent output")
	flagSet.StringSliceVarP(&opt.inputs, "inputs", "i", nil, "list of inputs (file,comma-separated)", goflags.FileCommaSeparatedStringSliceOptions)

	// Group example
	flagSet.CreateGroup("config", "Configuration",
		flagSet.StringVar(&opt.config, "config", "", "file to read config from"),
		flagSet.RuntimeMapVar(&opt.values, "values", nil, "key-value runtime values"),
	)
	if err := flagSet.Parse(); err != nil {
		log.Fatalf("Could not parse flags: %s\n", err)
	}
	if opt.config != "" {
		if err := flagSet.MergeConfigFile(opt.config); err != nil {
			log.Fatalf("Could not merge config file: %s\n", err)
		}
	}
	fmt.Printf("silent: %v inputs: %v config: %v values: %v\n", opt.silent, opt.inputs, opt.config, opt.values)

}
