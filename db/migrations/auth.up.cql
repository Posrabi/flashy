ALTER KEYSPACE system_auth WITH replication = { 'class' : 'NetworkTopologyStrategy', 'replication_factor': 3}
  AND durable_writes = true;
