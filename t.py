import memcache

c = memcache.Client(['localhost:9979'])
print(c.set('key', 'value'))
c = memcache.Client(['localhost:9979'])
print(c.get('key'))
