ALTER KEYSPACE system_auth WITH replication = { 'class' : 'SimpleStrategy', 'replication_factor': 3}
  AND durable_writes = true;
