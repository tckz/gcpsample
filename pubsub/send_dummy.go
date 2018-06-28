package pubsub

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"

	gps "cloud.google.com/go/pubsub"
	gcp "github.com/tckz/gcpsample"
)

func usageSendDummy() {
	fmt.Fprintf(os.Stderr, "Usage: send-dummy topicName\n")
}

func SendDummy(argv []string) int {

	fs := flag.NewFlagSet("send-dummy", flag.ExitOnError)
	parallel := fs.Int("parallel", 1, "Parallelism of publishing process")
	count := fs.Int("count", 100, "Number of publishing")

	if err := fs.Parse(argv); err != nil {
		return 2
	}

	if fs.NArg() == 0 {
		usageSendDummy()
		return 2
	}
	topicName := fs.Arg(0)

	ctx, _ := context.WithCancel(context.Background())

	client, err := gps.NewClient(ctx, gcp.ProjectID)
	if err != nil {
		log.Fatalf("*** Failed to pubsub.NewClient: %v", err)
	}

	ch := make(chan int, *parallel)
	wg := &sync.WaitGroup{}
	for i := 0; i < *parallel; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			topic := client.Topic(topicName)

			for n := range ch {
				sid, err := topic.Publish(ctx, &gps.Message{
					Data: []byte(strconv.Itoa(n)),
				}).Get(ctx)

				if err != nil {
					log.Fatalf("*** Failed to Publish: %v", err)
				}
				fmt.Fprintf(os.Stderr, "%s\n", sid)
			}
		}()
	}

	for i := 0; i < *count; i++ {
		ch <- i
	}
	close(ch)

	wg.Wait()

	return 0

}
