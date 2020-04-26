package corm

import "github.com/gocql/gocql"

type DB struct {
	Value   interface{}
	Error   error
	session *gocql.Session
}

// Experimental
// use PlugIn(*gocql.Session) alternatively
func Open(hosts []string, keyspace string, user string, password string, port int) (*DB, error) {
	db := DB{}
	var err error

	// connect to the cluster
	cluster := gocql.NewCluster(hosts...)
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: user,
		Password: password,
	}
	cluster.Port = port
	cluster.Keyspace = keyspace
	cluster.Consistency = gocql.Quorum // TODO: should be passed in

	if db.session, err = cluster.CreateSession(); err != nil {
		return &DB{}, err
	}

	return &db, nil
}

// Plug in an existing GoCQL connection
func PlugIn(session *gocql.Session) *DB {
	db := DB{
		session: session,
	}

	return &db
}

func (d *DB) GetSession() *gocql.Session {
	return d.session
}
