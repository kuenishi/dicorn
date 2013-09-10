# dicorn

`dicorn` is a-sort-of translator which creates abstract interface on every database


```
memcache interface -----+                +-- leveldb backend
                        |                |
raw http interface -----+   +--------+   +-- Riak backend
                        +---+ dicorn +---+
cassandra interface ----+   +--------+   +-- spam backend
                        |                |
hbase interface --------+                +-- memory backend

```

Dunno all of this get success.

- [memcache interface](https://github.com/memcached/memcached/blob/master/doc/protocol.txt)

## quickstart

[Install Go](http://golang.org/doc/install) . Clone this repo and build.

```sh
$ git clone git://github.com/kuenishi/dicorn
$ cd dicorn
$ go get
$ make
```

## Riak client

Currently using [riakpbc](https://github.com/mrb/riakpbc). Another option is [goriakpbc](https://github.com/tpjg/goriakpbc) .

## bench

Micro benchmark on my MacBook Air

memcached

```
$ ./memcached
...
$ ./clients/memslap  --servers=localhost:11211 --non-blocking --concurrency=10
	Threads connecting to servers 10
	Took 1.987 seconds to load data
```

dicorn

```
$ ./dicorn -backend=mem -cmd=memcache -listen=localhost:11211
listening on localhost:11211
backend: mem
...
$ ./clients/memslap  --servers=localhost:11211 --non-blocking --concurrency=10
	Threads connecting to servers 10
	Took 2.705 seconds to load data
```

dicorn backed by Riak (single and local node)
```
$ ./dicorn -backend=riak -cmd=memcache -listen=localhost:11211    [master ~/src/dicorn]
2013/09/01 23:13:43 server.go:20: listening on localhost:11211
2013/09/01 23:13:43 server.go:21: backend: riak
2013/09/01 23:13:43 riak_backend.go:31: riak is ok: &{[localhost:8087] 0xf84005c680 <nil> false 1000}
...
$ ./clients/memslap  --servers=localhost:11211 --non-blocking --concurrency=10
	Threads connecting to servers 10
	Took 88.487 seconds to load data
```

Yay! So slow.

## version

-200 million

## author

2013, @kuenishi

## license

Apache 2.0
