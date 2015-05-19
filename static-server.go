package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/jessevdk/go-flags"
)

type Options struct {
	Host string `short:"p" long:"port" description:"port" default:""`
	Port int    `short:"h" long:"host" description:"host" default:"8080"`
	Path string `long:"path" description:"path" default:"."`
}

var opts Options

var parser = flags.NewParser(&opts, flags.Default)

func init() {
	if _, err := parser.Parse(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	spew.Dump(opts)
	log.Fatal(
		http.ListenAndServe(
			fmt.Sprintf("%s:%d", opts.Host, opts.Port),
			http.FileServer(http.Dir(opts.Path)),
		),
	)
}
