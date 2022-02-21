package api

import (
	"os"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
)

type AccessType int
type Keyspace string

const (
	ReadOnly AccessType = iota + 1
	WriteOnly
	ReadAndWrite
)

const (
	UsersSpace Keyspace = "users"
)

const (
	// TODO: find a way to proxy these addresses.
	node1 string = "10.101.109.2"
	node2 string = "10.98.165.168"
	node3 string = "10.99.215.125"
)

func GetAccessToDB(level AccessType, space Keyspace) (gocqlx.Session, error) {
	cluster := gocql.NewCluster(node1, node2, node3)
	cluster.Keyspace = string(space)
	cluster.Consistency = gocql.Quorum
	cluster.SerialConsistency = gocql.Serial
	cluster.PoolConfig.HostSelectionPolicy = gocql.TokenAwareHostPolicy(gocql.RoundRobinHostPolicy())

	switch level {
	case ReadOnly:
		cluster.Authenticator = gocql.PasswordAuthenticator{
			Username: os.Getenv("READ_USER"),
			Password: os.Getenv("READ_PASS"),
		}
	case WriteOnly:
		cluster.Authenticator = gocql.PasswordAuthenticator{
			Username: os.Getenv("WRITE_USER"),
			Password: os.Getenv("WRITE_PASS"),
		}
	case ReadAndWrite:
		cluster.Authenticator = gocql.PasswordAuthenticator{
			Username: os.Getenv("READWRITE_USER"),
			Password: os.Getenv("READWRITE_PASS"),
		}
	}
	return gocqlx.WrapSession(cluster.CreateSession())
}
