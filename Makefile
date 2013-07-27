.PHONY: all clean

all: dicorn

dicorn: main.go
	go build -o $@

clean:
	-rm -rf dicorn
