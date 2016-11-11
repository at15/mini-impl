# Time series database based on Cassandra

## Usage

- `docker run --name tsdb-cassandra -p 9042:9042 -d cassandra:3.9`
    - `docker stop tsdb-cassandra`
    - `docker start tsdb-cassandra`
- `docker exec -it tsdb-cassandra bash`


- `docker run --name tsdb-scylla -p 9042:9042 -d scylladb/scylla:1.4.1`

## Key points

> Typically, a cluster has one keyspace per application composed of many different tables.
- [Cassandra support user defined aggregation, min, max, group by etc](http://docs.datastax.com/en/cql/3.3/cql/cql_using/useCreateUDA.html)

## TODO

- [x] setup cassandra using docker
- [ ] query cassandra
- [ ] create time series tables
- [ ] query https://github.com/scylladb/scylla  (maybe hard to setup?)

## NOTE

- need to use `-p` to map port to host, `expose` is not enough
- `failed to connect to 127.0.0.1:9042 due to error: Keyspace 'demo' does not exist`
> Before you execute the program, Launch `cqlsh` and execute:
create keyspace demo with replication = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };
(this is not include in the datastax documentation ... )

## Ref

- https://academy.datastax.com/resources/getting-started-apache-cassandra-and-go
- http://www.scylladb.com/2016/07/26/kairosdb/

- http://www.datastax.com/dev/blog/advanced-time-series-with-cassandra
> either the column values contain row keys pointing to a separate column family which contains the actual data for events,
or the complete set of data for each event is stored in the timeline itself
- https://www.tutorialspoint.com/cassandra/cassandra_data_model.htm

All the primary key partition stuff

- http://stackoverflow.com/questions/24949676/difference-between-partition-key-composite-key-and-clustering-key-in-cassandra

Wide row

- http://www.datastax.com/dev/blog/does-cql-support-dynamic-columns-wide-rows

Time sereis data modeling with apache cassandra

- https://academy.datastax.com/resources/getting-started-time-series-data-modeling

- http://dtrapezoid.com/time-series-data-modeling-for-medical-devices.html
