package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	gcp "github.com/tckz/gcpsample"
	"github.com/tckz/gcpsample/pubsub"
)

var (
	version = "to be replaced"
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [-project-id pjid] {pull-message|send-dummy} [sub command args]\n", path.Base(os.Args[0]))

}

func main() {
	projectID := flag.String("project-id", os.Getenv("GOOGLE_CLOUD_PROJECT"), "GCP Project id")
	showVersion := flag.Bool("version", false, "Show version")
	flag.Parse()

	if *showVersion {
		fmt.Fprintf(os.Stderr, "%s\n", version)
		os.Exit(0)
	}

	if flag.NArg() == 0 {
		usage()
		os.Exit(2)
	}

	if *projectID == "" {
		fmt.Fprintf(os.Stderr, "Project ID must be specified.\n")
		os.Exit(2)

	}

	gcp.ProjectID = *projectID

	command := flag.Arg(0)
	rest := flag.Args()[1:]

	var ec int
	switch command {
	case "pull-message":
		ec = pubsub.PullMessage(rest)
	case "send-dummy":
		ec = pubsub.SendDummy(rest)
	default:
		usage()
		ec = 2
	}

	os.Exit(ec)
}
