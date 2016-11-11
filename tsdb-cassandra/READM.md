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

## cql

KairosDB 

NOTE: the table created by thrift will have `column1` as column name when query 

````java
private void createSchema(int replicationFactor)
{
    List<ColumnFamilyDefinition> cfDef = new ArrayList<ColumnFamilyDefinition>();

    cfDef.add(HFactory.createColumnFamilyDefinition(
            m_keyspaceName, CF_DATA_POINTS, ComparatorType.BYTESTYPE));

    cfDef.add(HFactory.createColumnFamilyDefinition(
            m_keyspaceName, CF_ROW_KEY_INDEX, ComparatorType.BYTESTYPE));

    cfDef.add(HFactory.createColumnFamilyDefinition(
            m_keyspaceName, CF_STRING_INDEX, ComparatorType.UTF8TYPE));

    KeyspaceDefinition newKeyspace = HFactory.createKeyspaceDefinition(
            m_keyspaceName, ThriftKsDef.DEF_STRATEGY_CLASS,
            replicationFactor, cfDef);

    m_cluster.addKeyspace(newKeyspace, true);
}
````

Heroic

````
CREATE KEYSPACE IF NOT EXISTS {{keyspace}}
  WITH REPLICATION = {
    'class' : 'SimpleStrategy',
    'replication_factor' : 3
  };

CREATE TABLE IF NOT EXISTS {{keyspace}}.metrics (
  metric_key blob,
  data_timestamp_offset int,
  data_value double,
  PRIMARY KEY(metric_key, data_timestamp_offset)
) WITH COMPACT STORAGE;
````
Heroic store meta data in elastic search

````json
{
  "metadata": {
    "properties": {
      "key": {
        "index": "not_analyzed",
        "type": "string",
        "doc_values": true,
        "include_in_all": false
      },
      "tags": {
        "type": "string",
        "index": "not_analyzed",
        "doc_values": true,
        "include_in_all": false
      },
      "tag_keys": {
        "type": "string",
        "index": "not_analyzed",
        "doc_values": true,
        "include_in_all": false
      }
    }
  }
}
````

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
