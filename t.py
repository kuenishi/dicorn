def memcached_test():
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

def cassandra_test():
    import pycassa
    from pycassa.pool import ConnectionPool
    from pycassa.columnfamily import ColumnFamily

    pool = ConnectionPool('Keyspace1', ['localhost:9979'])
    col_fam = ColumnFamily(pool, 'ColumnFamily1')
    r = col_fam.insert('row_key', {'col_name': 'col_val'})
    print(r)
    r = col_fam.get('row_key', {'col_name':'col_val'})
    print(r)

def cassandra_manage_test():
    import pycassa
    from pycassa.system_manager import SystemManager
    sys = SystemManager('localhost:9979')
    # Create a SimpleStrategy keyspace
    sys.create_keyspace('Keyspace1', pycassa.system_manager.SIMPLE_STRATEGY, {'replication_factor': '1'})
    # Create a NetworkTopologyStrategy keyspace
    #sys.create_keyspace('NTS_KS', pycassa.system_manager.NETWORK_TOPOLOGY_STRATEGY, {'DC1': '2', 'DC2': '1'})
    r = sys.create_column_family('Keyspace1', 'ColumnFamily1')
    print(r)
    sys.close()

cassandra_manage_test()
cassandra_test()
