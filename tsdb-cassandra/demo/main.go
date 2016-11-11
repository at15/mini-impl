package main

// https://academy.datastax.com/resources/getting-started-apache-cassandra-and-go

import (
    "log"
    "github.com/gocql/gocql"
)

func main() {
    cluster := gocql.NewCluster("127.0.0.1")
    cluster.Keyspace = "demo"
    session, err := cluster.CreateSession()
    defer session.Close()
    if err != nil {
        log.Fatal(err)
    }
    log.Print("connected!")
}