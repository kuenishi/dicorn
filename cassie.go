package dicorn

import (
	"fmt"
	"net"
	"net/rpc"
	"os"
	"github.com/samuel/go-thrift/thrift"

	"dicorn/cassandra"
 )

type cassie struct { //implementation
}

func (c *cassie) Add(
	Key []byte,
	ColumnParent *cassandra.ColumnParent,
	Column *cassandra.CounterColumn,
	ConsistencyLevel cassandra.ConsistencyLevel) error {
	return nil
}
func (c *cassie)  AtomicBatchMutate(MutationMap map[string]map[string][]* cassandra.Mutation, ConsistencyLevel cassandra.ConsistencyLevel) error  {return nil}
func (c *cassie)  BatchMutate(MutationMap map[string]map[string][]*cassandra.Mutation,
	ConsistencyLevel cassandra.ConsistencyLevel) error {return nil}
func (c *cassie)  Cas(Key []byte, ColumnFamily string, Expected []*cassandra.Column, Updates []*cassandra.Column, SerialConsistencyLevel cassandra.ConsistencyLevel, CommitConsistencyLevel cassandra.ConsistencyLevel) (*cassandra.CASResult, error) {return nil,nil}
func (c *cassie)  DescribeClusterName() (string, error) {return "",nil}
func (c *cassie)  DescribeKeyspace(Keyspace string) (*cassandra.KsDef, error) {return nil,nil}
func (c *cassie)  DescribeKeyspaces() ([]*cassandra.KsDef, error) {return nil,nil}
func (c *cassie)  DescribePartitioner() (string, error) {return "",nil}
func (c *cassie)  DescribeRing(Keyspace string) ([]*cassandra.TokenRange, error) {return nil,nil}
func (c *cassie)  DescribeSchemaVersions() (map[string][]string, error) {return nil,nil}
func (c *cassie)  DescribeSnitch() (string, error) {return "",nil}
func (c *cassie)  DescribeSplits(CfName string, StartToken string, EndToken string, KeysPerSplit int32) ([]string, error) {return nil,nil}
func (c *cassie)  DescribeSplitsEx(CfName string, StartToken string, EndToken string, KeysPerSplit int32) ([]*cassandra.CfSplit, error) {return nil,nil}
func (c *cassie)  DescribeTokenMap() (map[string]string, error) {return nil,nil}
func (c *cassie)  DescribeVersion() (string, error) {return "",nil}
func (c *cassie)  ExecuteCql3Query(Query []byte, Compression cassandra.Compression, Consistency cassandra.ConsistencyLevel) (*cassandra.CqlResult, error) {return nil,nil}
func (c *cassie)  ExecuteCqlQuery(Query []byte, Compression cassandra.Compression) (*cassandra.CqlResult, error) {return nil,nil}
func (c *cassie)  ExecutePreparedCql3Query(ItemId int32, Values [][]byte, Consistency cassandra.ConsistencyLevel) (*cassandra.CqlResult, error) {return nil,nil}
func (c *cassie)  ExecutePreparedCqlQuery(ItemId int32, Values [][]byte) (*cassandra.CqlResult, error) {return nil,nil}
func (c *cassie)  Get(Key []byte, ColumnPath *cassandra.ColumnPath, ConsistencyLevel cassandra.ConsistencyLevel) (*cassandra.ColumnOrSuperColumn, error) {return nil,nil}
func (c *cassie)  GetCount(Key []byte, ColumnParent *cassandra.ColumnParent, Predicate *cassandra.SlicePredicate, ConsistencyLevel cassandra.ConsistencyLevel) (int32, error) {return -1,nil}
func (c *cassie)  GetIndexedSlices(ColumnParent *cassandra.ColumnParent, IndexClause *cassandra.IndexClause, ColumnPredicate *cassandra.SlicePredicate, ConsistencyLevel cassandra.ConsistencyLevel) ([]*cassandra.KeySlice, error) {return nil,nil}
func (c *cassie)  GetPagedSlice(ColumnFamily string, Range *cassandra.KeyRange, StartColumn []byte, ConsistencyLevel cassandra.ConsistencyLevel) ([]*cassandra.KeySlice, error) {return nil,nil}
func (c *cassie)  GetRangeSlices(ColumnParent *cassandra.ColumnParent, Predicate *cassandra.SlicePredicate, Range *cassandra.KeyRange, ConsistencyLevel cassandra.ConsistencyLevel) ([]*cassandra.KeySlice, error) {return nil,nil}
func (c *cassie)  GetSlice(Key []byte, ColumnParent *cassandra.ColumnParent, Predicate *cassandra.SlicePredicate, ConsistencyLevel cassandra.ConsistencyLevel) ([]*cassandra.ColumnOrSuperColumn, error) {return nil,nil}
func (c *cassie)  Insert(Key []byte, ColumnParent *cassandra.ColumnParent, Column *cassandra.Column, ConsistencyLevel cassandra.ConsistencyLevel) error {return nil}
func (c *cassie)  Login(AuthRequest *cassandra.AuthenticationRequest) error {return nil}
func (c *cassie)  MultigetCount(Keys [][]byte, ColumnParent *cassandra.ColumnParent, Predicate *cassandra.SlicePredicate, ConsistencyLevel cassandra.ConsistencyLevel) (map[string]int32, error) {return nil,nil}
func (c *cassie)  MultigetSlice(Keys [][]byte, ColumnParent *cassandra.ColumnParent, Predicate *cassandra.SlicePredicate, ConsistencyLevel cassandra.ConsistencyLevel) (map[string][]*cassandra.ColumnOrSuperColumn, error) {return nil,nil}
func (c *cassie)  PrepareCql3Query(Query []byte, Compression cassandra.Compression) (*cassandra.CqlPreparedResult, error) {return nil,nil}
func (c *cassie)  PrepareCqlQuery(Query []byte, Compression cassandra.Compression) (*cassandra.CqlPreparedResult, error) {return nil,nil}
func (c *cassie)  Remove(Key []byte, ColumnPath *cassandra.ColumnPath, Timestamp int64, ConsistencyLevel cassandra.ConsistencyLevel) error {return nil}
func (c *cassie)  RemoveCounter(Key []byte, Path *cassandra.ColumnPath, ConsistencyLevel cassandra.ConsistencyLevel) error {return nil}
func (c *cassie)  SetCqlVersion(Version string) error {return nil}
func (c *cassie)  SetKeyspace(Keyspace string) error {return nil}
func (c *cassie)  SystemAddColumnFamily(CfDef *cassandra.CfDef) (string, error) {return "",nil}
func (c *cassie)  SystemAddKeyspace(KsDef *cassandra.KsDef) (string, error) {return "",nil}
func (c *cassie)  SystemDropColumnFamily(ColumnFamily string) (string, error) {return "",nil}
func (c *cassie)  SystemDropKeyspace(Keyspace string) (string, error) {return "",nil}
func (c *cassie)  SystemUpdateColumnFamily(CfDef *cassandra.CfDef) (string, error) {return "",nil}
func (c *cassie)  SystemUpdateKeyspace(KsDef *cassandra.KsDef) (string, error) {return "",nil}
func (c *cassie)  TraceNextQuery() ([]byte, error) {return nil,nil}
func (c *cassie)  Truncate(Cfname string) error {return nil}


func RunCass(addr, backend_type, backend_hosts string) {
	cassieService := new(cassie)
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

