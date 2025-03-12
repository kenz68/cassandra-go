package cassandra

import "github.com/gocql/gocql"
import "github.com/sirupsen/logrus"

func ConnectDatabase(url string, keyspace string) *gocql.Session {
	cluster := gocql.NewCluster(url)
	cluster.Keyspace = keyspace

	session, err := cluster.CreateSession()

	if err != nil {
		logrus.Fatal(err)
	}

	return session
}
