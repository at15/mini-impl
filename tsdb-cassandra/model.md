# Data model

http://www.datastax.com/dev/blog/cql3-for-cassandra-experts

denormalization & wide row

> Early observers of the NoSQL space concluded that if Cassandra does not do joins for you, then joins should be done client-side. Not so! Instead of joins, Cassandra encourages denormalization: at update time, clients should write the records to multiple columnfamilies, such that queries (known ahead of time) can be retrieved from a single row.

> To support this, Cassandra’s storage engine provides wide, sparse rows. These rows can correspond 1:1 with business objects, but more often they encode data in the cell name as well as the value — thus a “row” becomes more of an (ordered) map, than a relational row. This is the kind of model we had in mind when we improved the storage engine to allow up to two billion columns per row.

> Materialized views: Cassandra rows are resultsets; objects are each packed into one column in the cassandra row using CompositeType

think rows as partitions from relational DB view

> CQL3 makes one very important changes to how it presents Cassandra data: wide rows are “transposed” and unpacked into named columns. From a relational standpoint, you can think of storage engine rows as partitions, within which (object) rows are clustered.

CQL3 does not like Thrift
while Thrift feels like expose the storage engine directly to the users
CQL3 provide some abstraction to make it more familiar to RDBMS users


with single key as primary key, C* works like normal RDBMS

The following show wide row is achieved by using compound key

````
cqlsh:demo> create table t1 (id int, tag text, PRIMARY KEY (id));
cqlsh:demo> insert t1 (id, tag) values (1, 'a');
cqlsh:demo> insert into t1 (id, tag) values (1, 'a');
cqlsh:demo> select * from t1;

 id | tag
----+-----
  1 |   a

(1 rows)
cqlsh:demo> insert into t1 (id, tag) values (1, 'b');
cqlsh:demo> select * from t1;

 id | tag
----+-----
  1 |   b

(1 rows)
cqlsh:demo> create table t2 (id int, tag text, PRIMARY KEY (id, tag));
cqlsh:demo> insert into t2 (id, tag) values (1, 'a');
cqlsh:demo> insert into t2 (id, tag) values (1, 'b');
cqlsh:demo> select * from t2;

 id | tag
----+-----
  1 |   a
  1 |   b

(2 rows)
````
