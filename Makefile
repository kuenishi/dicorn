.PHONY: all clean

all: dicorn

dicorn: cmd/main.go *.go
	go build -o $@ $<

clean:
	-rm -rf dicorn
