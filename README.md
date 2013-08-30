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

Currently [goriakpbc](https://github.com/tpjg/goriakpbc) and [riakpbc](https://github.com/mrb/riakpbc) looks fine. Which to choose.

## version

-200 million

## author

2013, @kuenishi

## license

Apache 2.0
