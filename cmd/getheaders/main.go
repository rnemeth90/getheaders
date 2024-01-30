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
	jsonOut bool
	yamlOut bool
)

type config struct {
	url string
}

func init() {
	pflag.StringVarP(&url, "url", "u", "", "The URL to query [required]")
	pflag.BoolVarP(&jsonOut, "json", "j", false, "Output in JSON format")
	pflag.BoolVarP(&yamlOut, "yaml", "y", false, "Output in YAML format")
}

// usage displays the usage information for the program
func usage() {
	fmt.Println("Usage of getheaders:")
	fmt.Println("  -u, --url string")
	fmt.Println("        The URL to query [required]")
	fmt.Println("  -j, --json")
	fmt.Println("        Output in JSON format")
	fmt.Println("  -y, --yaml")
	fmt.Println("        Output in YAML format")
}

func main() {
	pflag.Parse()

	nonPflagArgs := pflag.Args()

	if url == "" && len(nonPflagArgs) == 0 {
		usage()
		os.Exit(1)
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
