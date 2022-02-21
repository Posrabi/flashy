package api

import (
	"os"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
)

type AccessType int

const (
	ReadOnly AccessType = iota + 1
	WriteOnly
	ReadAndWrite
)

const (
	node1 string = ""
	node2 string = ""
	node3 string = ""
)

func GetAccessToDB(level AccessType) (gocqlx.Session, error) {
	cluster := gocql.NewCluster(node1, node2, node3)
	cluster.Consistency = gocql.Quorum

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
