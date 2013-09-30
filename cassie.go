package dicorn

import (
	"fmt"
	"github.com/samuel/go-thrift/thrift"
	"net"
	"net/rpc"
	"os"

	"dicorn/cassandra"
	"log"
	"runtime"
)

type cassie struct { //implementation in memory??
	BackendType     string
	BackendHosts    string
	KeySpaces       []*cassandra.KsDef
	ColumnDefs      map[string]map[string]*cassandra.ColumnDef
	Columns         map[string](map[string]*cassandra.Column)
	CurrentKeySpace string
}

func NewCassie(backend, backend_hosts string) *cassie {
	m := &cassie{
		BackendType:  backend,
		BackendHosts: backend_hosts,
	}
	return m
}

func (c *cassie) getKeySpace() (*cassandra.KsDef, error) {
	for _, ks := range c.KeySpaces {
		if ks.Name == c.CurrentKeySpace {
			log.Printf("keyspace %s found: %v.", c.CurrentKeySpace, ks)
			return ks, nil
		}
	}
	log.Printf("keyspace %s not found.", c.CurrentKeySpace)
	return nil, &cassandra.NotFoundException{}
}

func (c *cassie) Add(
	Key []byte,
	ColumnParent *cassandra.ColumnParent,
	Column *cassandra.CounterColumn,
	ConsistencyLevel cassandra.ConsistencyLevel) error {
	return nil
}

func (c *cassie) AtomicBatchMutate(MutationMap map[string]map[string][]*cassandra.Mutation, ConsistencyLevel cassandra.ConsistencyLevel) error {
	_, _, l, _ := runtime.Caller(0)
	log.Printf("line %d\n", l)
	return nil
}
func (c *cassie) BatchMutate(MutationMap map[string]map[string][]*cassandra.Mutation,
	ConsistencyLevel cassandra.ConsistencyLevel) error {
	log.Printf("BatchMutate: %v", MutationMap)
	for rowKey, rowMutation := range MutationMap {
		for columnFamily, mutations := range rowMutation {
			for _, mutation := range mutations {
				log.Printf("%s>%s> %v", rowKey, columnFamily, mutation)
			}
		}
	}
	return nil
}
func (c *cassie) Cas(Key []byte, ColumnFamily string, Expected []*cassandra.Column, Updates []*cassandra.Column, SerialConsistencyLevel cassandra.ConsistencyLevel, CommitConsistencyLevel cassandra.ConsistencyLevel) (*cassandra.CASResult, error) {
	_, _, l, _ := runtime.Caller(0)
	log.Printf("line %d\n", l)
	return nil, nil
}
func (c *cassie) DescribeClusterName() (string, error) {
	_, _, l, _ := runtime.Caller(0)
	log.Printf("line %d\n", l)
	return "", nil
}
func (c *cassie) DescribeKeyspace(Keyspace string) (*cassandra.KsDef, error) {
	pc, _, l, _ := runtime.Caller(0)
	log.Printf("line %d at %s\n", l, runtime.FuncForPC(pc).Name())
	ks, _ := c.getKeySpace()
	log.Printf("%v", ks.CfDefs[0].ColumnMetadata)
	return c.getKeySpace()
}
func (c *cassie) DescribeKeyspaces() ([]*cassandra.KsDef, error) {
	_, _, l, _ := runtime.Caller(0)
	log.Printf("line %d\n", l)
	return nil, nil
}
func (c *cassie) DescribePartitioner() (string, error) {
	_, _, l, _ := runtime.Caller(0)
	log.Printf("line %d\n", l)
	return "", nil
}
func (c *cassie) DescribeRing(Keyspace string) ([]*cassandra.TokenRange, error) {
	_, _, l, _ := runtime.Caller(0)
	log.Printf("line %d\n", l)
	return nil, nil
}
func (c *cassie) DescribeSchemaVersions() (map[string][]string, error) {
	_, _, l, _ := runtime.Caller(0)
	log.Printf("line %d\n", l)
	return nil, nil
}
func (c *cassie) DescribeSnitch() (string, error) {
	_, _, l, _ := runtime.Caller(0)
	log.Printf("line %d\n", l)
	return "", nil
}
func (c *cassie) DescribeSplits(CfName string, StartToken string, EndToken string, KeysPerSplit int32) ([]string, error) {
	_, _, l, _ := runtime.Caller(0)
	log.Printf("line %d\n", l)
	return nil, nil
}
func (c *cassie) DescribeSplitsEx(CfName string, StartToken string, EndToken string, KeysPerSplit int32) ([]*cassandra.CfSplit, error) {
	_, _, l, _ := runtime.Caller(0)
	log.Printf("line %d\n", l)
	return nil, nil
}
func (c *cassie) DescribeTokenMap() (map[string]string, error) {
	_, _, l, _ := runtime.Caller(0)
	log.Printf("line %d\n", l)
	return nil, nil
}
func (c *cassie) DescribeVersion() (string, error) {
	_, _, l, _ := runtime.Caller(0)
	log.Printf("line %d\n", l)
	return "", nil
}
func (c *cassie) ExecuteCql3Query(Query []byte, Compression cassandra.Compression, Consistency cassandra.ConsistencyLevel) (*cassandra.CqlResult, error) {
	_, _, l, _ := runtime.Caller(0)
	log.Printf("line %d\n", l)
	return nil, nil
}
func (c *cassie) ExecuteCqlQuery(Query []byte, Compression cassandra.Compression) (*cassandra.CqlResult, error) {
	_, _, l, _ := runtime.Caller(0)
	log.Printf("line %d\n", l)
	return nil, nil
}
func (c *cassie) ExecutePreparedCql3Query(ItemId int32, Values [][]byte, Consistency cassandra.ConsistencyLevel) (*cassandra.CqlResult, error) {
	_, _, l, _ := runtime.Caller(0)
	log.Printf("line %d\n", l)
	return nil, nil
}
func (c *cassie) ExecutePreparedCqlQuery(ItemId int32, Values [][]byte) (*cassandra.CqlResult, error) {
	_, _, l, _ := runtime.Caller(0)
	log.Printf("line %d\n", l)
	return nil, nil
}
func (c *cassie) Get(Key []byte, ColumnPath *cassandra.ColumnPath, ConsistencyLevel cassandra.ConsistencyLevel) (*cassandra.ColumnOrSuperColumn, error) {
	_, _, l, _ := runtime.Caller(0)
	log.Printf("line %d\n", l)
	return nil, nil
}
func (c *cassie) GetCount(Key []byte, ColumnParent *cassandra.ColumnParent, Predicate *cassandra.SlicePredicate, ConsistencyLevel cassandra.ConsistencyLevel) (int32, error) {
	_, _, l, _ := runtime.Caller(0)
	log.Printf("line %d\n", l)
	return -1, nil
}
func (c *cassie) GetIndexedSlices(ColumnParent *cassandra.ColumnParent, IndexClause *cassandra.IndexClause, ColumnPredicate *cassandra.SlicePredicate, ConsistencyLevel cassandra.ConsistencyLevel) ([]*cassandra.KeySlice, error) {
	_, _, l, _ := runtime.Caller(0)
	log.Printf("line %d\n", l)
	return nil, nil
}
func (c *cassie) GetPagedSlice(ColumnFamily string, Range *cassandra.KeyRange, StartColumn []byte, ConsistencyLevel cassandra.ConsistencyLevel) ([]*cassandra.KeySlice, error) {
	_, _, l, _ := runtime.Caller(0)
	log.Printf("line %d\n", l)
	return nil, nil
}
func (c *cassie) GetRangeSlices(ColumnParent *cassandra.ColumnParent, Predicate *cassandra.SlicePredicate, Range *cassandra.KeyRange, ConsistencyLevel cassandra.ConsistencyLevel) ([]*cassandra.KeySlice, error) {
	_, _, l, _ := runtime.Caller(0)
	log.Printf("line %d\n", l)
	return nil, nil
}
func (c *cassie) GetSlice(Key []byte, ColumnParent *cassandra.ColumnParent, Predicate *cassandra.SlicePredicate, ConsistencyLevel cassandra.ConsistencyLevel) ([]*cassandra.ColumnOrSuperColumn, error) {
	_, _, l, _ := runtime.Caller(0)
	log.Printf("line %d\n", l)
	return nil, nil
}
func (c *cassie) Insert(Key []byte, ColumnParent *cassandra.ColumnParent, Column *cassandra.Column, ConsistencyLevel cassandra.ConsistencyLevel) error {
	_, _, l, _ := runtime.Caller(0)
	log.Printf("line %d\n", l)
	return nil
}
func (c *cassie) Login(AuthRequest *cassandra.AuthenticationRequest) error {
	_, _, l, _ := runtime.Caller(0)
	log.Printf("line %d\n", l)
	return nil
}
func (c *cassie) MultigetCount(Keys [][]byte, ColumnParent *cassandra.ColumnParent, Predicate *cassandra.SlicePredicate, ConsistencyLevel cassandra.ConsistencyLevel) (map[string]int32, error) {
	_, _, l, _ := runtime.Caller(0)
	log.Printf("line %d\n", l)
	return nil, nil
}
func (c *cassie) MultigetSlice(Keys [][]byte, ColumnParent *cassandra.ColumnParent, Predicate *cassandra.SlicePredicate, ConsistencyLevel cassandra.ConsistencyLevel) (map[string][]*cassandra.ColumnOrSuperColumn, error) {
	_, _, l, _ := runtime.Caller(0)
	log.Printf("line %d\n", l)
	return nil, nil
}
func (c *cassie) PrepareCql3Query(Query []byte, Compression cassandra.Compression) (*cassandra.CqlPreparedResult, error) {
	_, _, l, _ := runtime.Caller(0)
	log.Printf("line %d\n", l)
	return nil, nil
}
func (c *cassie) PrepareCqlQuery(Query []byte, Compression cassandra.Compression) (*cassandra.CqlPreparedResult, error) {
	_, _, l, _ := runtime.Caller(0)
	log.Printf("line %d\n", l)
	return nil, nil
}
func (c *cassie) Remove(Key []byte, ColumnPath *cassandra.ColumnPath, Timestamp int64, ConsistencyLevel cassandra.ConsistencyLevel) error {
	_, _, l, _ := runtime.Caller(0)
	log.Printf("line %d\n", l)
	return nil
}
func (c *cassie) RemoveCounter(Key []byte, Path *cassandra.ColumnPath, ConsistencyLevel cassandra.ConsistencyLevel) error {
	_, _, l, _ := runtime.Caller(0)
	log.Printf("line %d\n", l)
	return nil
}
func (c *cassie) SetCqlVersion(Version string) error {
	_, _, l, _ := runtime.Caller(0)
	log.Printf("line %d\n", l)
	return nil
}
func (c *cassie) SetKeyspace(Keyspace string) error {
	// TODO draw metadata from database if ever exists...
	c.CurrentKeySpace = Keyspace
	log.Printf("CurrentKeySpace set to %s", Keyspace)
	return nil
}
func (c *cassie) SystemAddColumnFamily(CfDef *cassandra.CfDef) (string, error) {
	ks, err := c.getKeySpace()
	if err == nil {
		ks.CfDefs = append(ks.CfDefs, CfDef)
		log.Printf("SystemAddColumnFamily: ColumnFamily added: %v", c)
		//log.Printf("%v", ks.CfDefs)
		return CfDef.Name, nil
	}
	log.Printf("SystemAddColumnFamily: KeySpace %s not found.", c.CurrentKeySpace)
	return CfDef.Name, &cassandra.NotFoundException{}
}
func (c *cassie) SystemAddKeyspace(KsDef *cassandra.KsDef) (string, error) {
	//_, _, l, _ := runtime.Caller(0)
	//log.Printf("line %d, %v\n", l, KsDef)
	c.KeySpaces = append(c.KeySpaces, KsDef)
	name := KsDef.Name
	return name, nil
}
func (c *cassie) SystemDropColumnFamily(ColumnFamily string) (string, error) {
	_, _, l, _ := runtime.Caller(0)
	log.Printf("line %d\n", l)
	return "", nil
}
func (c *cassie) SystemDropKeyspace(Keyspace string) (string, error) {
	_, _, l, _ := runtime.Caller(0)
	log.Printf("line %d\n", l)
	return "", nil
}
func (c *cassie) SystemUpdateColumnFamily(CfDef *cassandra.CfDef) (string, error) {
	_, _, l, _ := runtime.Caller(0)
	log.Printf("line %d\n", l)
	return "", nil
}
func (c *cassie) SystemUpdateKeyspace(KsDef *cassandra.KsDef) (string, error) {
	_, _, l, _ := runtime.Caller(0)
	log.Printf("line %d\n", l)
	return "", nil
}
func (c *cassie) TraceNextQuery() ([]byte, error) {
	_, _, l, _ := runtime.Caller(0)
	log.Printf("line %d\n", l)
	return nil, nil
}
func (c *cassie) Truncate(Cfname string) error {
	_, _, l, _ := runtime.Caller(0)
	log.Printf("line %d\n", l)
	return nil
}

func RunCass(addr, backend_type, backend_hosts string) {
	cassieService := NewCassie(backend_type, backend_hosts)
	rpc.RegisterName("Thrift", &cassandra.CassandraServer{cassieService})

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		println("error")
		os.Exit(-1)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("ERROR: %+v\n", err)
			continue
		}
		fmt.Printf("New connection %+v\n", conn)
		go rpc.ServeCodec(thrift.NewServerCodec(thrift.NewFramedReadWriteCloser(conn, 0), thrift.NewBinaryProtocol(true, false)))
	}
}
