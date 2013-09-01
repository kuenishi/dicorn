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

## Riak client

Currently using [riakpbc](https://github.com/mrb/riakpbc). Another option is [goriakpbc](https://github.com/tpjg/goriakpbc) .

## bench

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


## version

-200 million

## author

2013, @kuenishi

## license

Apache 2.0
