.PHONY: dist test clean

TARGETS=\
	dist/pubsub \

SRCS_OTHER = \
	$(wildcard */*.go) \
	$(wildcard *.go)

all: $(TARGETS)
	@echo "$@ done."

clean:
	/bin/rm -f $(TARGETS)
	@echo "$@ done."

dist/pubsub: cmd/pubsub/main.go $(SRCS_OTHER)
	if [ ! -d dist ];then mkdir dist; fi
	go build -o $@ -ldflags "-X main.version=`git show -s --format=%H`" $<
	@echo "$@ done."
