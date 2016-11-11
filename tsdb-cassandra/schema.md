# Schema

Schema from existing tsdb and blog posts

## KairosDB

KairosDB

NOTE: the table created using thrift not cql, so it will have `column1` as column name when query

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

## Spotify Heroic

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
