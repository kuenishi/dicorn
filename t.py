import memcache

c = memcache.Client(['localhost:9979'])

for i in xrange(0,10):
    print("set>", c.set('key', 'value'))
    print("get>", c.get('key'))
    print("del>", c.delete('key'))
    print("get>", c.get('key'))

    print("set>", c.set('i', 0))
#    print("inc>", c.incr('i', 2))
#    print("inc>", c.decr('i', 1))

print("set>", c.set('booo', 'value'))
