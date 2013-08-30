import memcache

c = memcache.Client(['localhost:9979'])
print("set>", c.set('key', 'value'))
print("get>", c.get('key'))
print("del>", c.delete('key'))
print("get>", c.get('key'))

c = memcache.Client(['localhost:9979'])
print("set>", c.set('i', 0))
print("inc>", c.incr('i', 1))
