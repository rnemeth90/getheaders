package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/rnemeth90/getheaders"
	"github.com/spf13/pflag"
)

var (
	url string
)

type config struct {
	u string
}

func init() {
	pflag.StringVar(&url, "url", "", "The URL to query [required]")
}

func main() {
	pflag.Parse()
	c := config{
		u: url,
	}

	if err := run(c); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(c config) error {
	result, err := getheaders.DoRequest(http.MethodGet, c.u)
	if err != nil {
		return err
	}

	jsonData, err := json.Marshal(result)
	if err != nil {
		return err
	}

	fmt.Println(string(jsonData))
	return nil
}
