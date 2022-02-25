package api

import (
	"os"
	"time"

	"github.com/gocql/gocql"
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
	node1      string = "10.101.109.2"
	node2      string = "10.98.165.168"
	node3      string = "10.99.215.125"
	retries           = 5
	maxRetries        = 10 * time.Second
)

func GetAccessToDB(level AccessType, space Keyspace) (*gocql.Session, error) {
	cluster := gocql.NewCluster(node1, node2, node3)
	cluster.Keyspace = string(space)
	cluster.Consistency = gocql.Quorum
	cluster.SerialConsistency = gocql.Serial
	cluster.ProtoVersion = 4
	cluster.PoolConfig.HostSelectionPolicy = gocql.TokenAwareHostPolicy(gocql.RoundRobinHostPolicy())
	cluster.RetryPolicy = &gocql.ExponentialBackoffRetryPolicy{
		Min:        time.Second,
		Max:        maxRetries,
		NumRetries: retries,
	}

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
	return cluster.CreateSession()
}
