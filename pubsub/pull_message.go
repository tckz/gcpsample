package pubsub

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	gps "cloud.google.com/go/pubsub"
	gcp "github.com/tckz/gcpsample"
)

func usagePullMessage() {
	fmt.Fprintf(os.Stderr, "Usage: pull-message subscripitonName\n")
}

func PullMessage(argv []string) int {

	fs := flag.NewFlagSet("pull-message", flag.ExitOnError)
	if err := fs.Parse(argv); err != nil {
		return 2
	}

	if fs.NArg() == 0 {
		usagePullMessage()
		return 2
	}

	subName := fs.Arg(0)
	ctx, _ := context.WithCancel(context.Background())

	client, err := gps.NewClient(ctx, gcp.ProjectID)
	if err != nil {
		log.Fatalf("*** Failed to pubsub.NewClient: %v", err)
	}

	subs := client.Subscription(subName)
	cctx, _ := context.WithCancel(ctx)
	for {
		err := subs.Receive(cctx, func(ctx context.Context, m *gps.Message) {
			m.Ack()

			fmt.Printf("%s\n", string(m.Data))
		})
		if err != nil {
			log.Fatalf("*** Failed to Receive(): %v", err)
		}
	}

	return 0
}
