package api

import (
	"os"
	"time"

	"github.com/gocql/gocql"
)

type AccessType int
type Keyspace string
type DBType string

const (
	ReadOnly AccessType = iota + 1
	WriteOnly
	ReadAndWrite
)

const (
	ProdDB DBType = "prod"
	DevDB  DBType = "dev"
)

const (
	UsersSpace Keyspace = "users"
)

// TODO: dynamically choose address

const (
	node1      string = "localhost:9042"
	node2      string = "localhost:9043"
	node3      string = "localhost:9044" // hosts for local running a local scylla cluster
	retries           = 5
	maxRetries        = 10 * time.Second
)

func SetupDB(level AccessType, dbType DBType) (*gocql.Session, error) {
	cluster := createClusterConfig(dbType)
	cluster.Keyspace = string(UsersSpace)
	cluster.Consistency = gocql.Quorum
	cluster.SerialConsistency = gocql.Serial
	cluster.ProtoVersion = 4
	cluster.PoolConfig.HostSelectionPolicy = gocql.TokenAwareHostPolicy(gocql.RoundRobinHostPolicy())
	cluster.RetryPolicy = &gocql.ExponentialBackoffRetryPolicy{
		Min:        time.Second,
		Max:        maxRetries,
		NumRetries: retries,
	}

	return setClusterCreds(cluster, level).CreateSession()
}

func createClusterConfig(dbType DBType) *gocql.ClusterConfig {
	switch dbType {
	case ProdDB:
		host := os.Getenv("DB_ENDPOINT")
		return gocql.NewCluster(host)
	case DevDB:
		return gocql.NewCluster(node1, node2, node3)
	default:
		return nil
	}
}

func setClusterCreds(cluster *gocql.ClusterConfig, level AccessType) *gocql.ClusterConfig {
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
	return cluster
}
