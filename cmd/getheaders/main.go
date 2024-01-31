package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/rnemeth90/getheaders"
	"github.com/spf13/pflag"
	"gopkg.in/yaml.v2"
)

var (
	url     string
	help    bool
	jsonOut bool
	yamlOut bool
)

type config struct {
	url string
}

func init() {
	pflag.StringVarP(&url, "url", "u", "", "The URL to query [required]")
	pflag.BoolVarP(&help, "help", "h", false, "Help")
	pflag.BoolVarP(&jsonOut, "json", "j", false, "Output in JSON format")
	pflag.BoolVarP(&yamlOut, "yaml", "y", false, "Output in YAML format")
	pflag.Usage = usage
	pflag.ErrHelp = nil
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	fmt.Fprintln(os.Stderr, "This program shows HTTP headers from a given URL")
	fmt.Fprintln(os.Stderr, "Flags:")
	pflag.PrintDefaults()
	fmt.Fprintln(os.Stderr, "\nExamples:")
	fmt.Fprintf(os.Stderr, "  %s --url https://rnemeth90.github.io --yaml\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "  %s --url https://rnemeth90.github.io --json\n", os.Args[0])
}

func main() {
	pflag.Parse()

	nonPflagArgs := pflag.Args()

	if url == "" && len(nonPflagArgs) == 0 {
		usage()
		os.Exit(0)
	}

	if len(nonPflagArgs) == 1 {
		url = nonPflagArgs[0]
	}

	c := config{
		url: url,
	}

	if err := run(c); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(c config) error {
	result, err := getheaders.DoRequest(http.MethodGet, c.url)
	if err != nil {
		return err
	}

	switch {
	case yamlOut:
		yamlData, err := yaml.Marshal(result)
		if err != nil {
			return err
		}
		fmt.Println(string(yamlData))
	case jsonOut:
		fallthrough
	default:
		jsonData, err := json.Marshal(result)
		if err != nil {
			return err
		}
		fmt.Println(string(jsonData))
	}

	return nil
}
