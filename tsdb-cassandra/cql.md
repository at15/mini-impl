# CQL

common CQL commands that are used for time series in Cassandra

## Create

Keyspace

- `create keyspace demo with replication = {'class':'SimpleStrategy', 'replication_factor': 1};`

Table

- [ ] [Compact storage](https://docs.datastax.com/en/cql/3.1/cql/cql_reference/create_table_r.html#reference_ds_v3f_vfk_xj__using-compact-storage)
Both KairosDB and Heroic are using compact storage

````
CREATE TABLE users (
  user_name varchar PRIMARY KEY,
  password varchar,
  gender varchar,
  )
````

http://www.datastax.com/dev/blog/does-cql-support-dynamic-columns-wide-rows

> Thus, the way to model dynamic cells in CQL is with a compound primary key. For the gory details on things like CompositeType, see my previous post.

## Describe

- `select * from system.schema_keyspaces;` or `describe keyspaces`
- `use kairosdb`
- `describe tables` list all tables in a key space
- `describe table string_index`

## Insert

- `insert into table_name (col1, col2, col3) values (v1, v2, v3)`
- use `uuid()` to generate uuid
- use `textAsBlob` to insert blob value in cqlsh
- when insert existing uuid, don't add quote mark around it

## Key

see [model](model.md) for more

- Must have primary key, when only one column is used as primary key, then it is also the partition key
- The first column in primary key is the partition key, the reset are clustering key
- COMPOSITE/COMPOUND KEY can retrieve "wide rows"

## Ref

Key

- http://stackoverflow.com/questions/24949676/difference-between-partition-key-composite-key-and-clustering-key-in-cassandra
- http://www.datastax.com/dev/blog/cql3-for-cassandra-experts
