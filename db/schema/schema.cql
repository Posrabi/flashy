CREATE KEYSPACE IF NOT EXISTS users
  WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 3}
  AND durable_writes = true;;

CREATE TABLE IF NOT EXISTS users.info (
  user_id uuid,
  user_name text,
  name text,
  email text,
  phone_number text,
  hash_password text,
  auth_token text,
  PRIMARY KEY ((user_id), user_name, name)
) WITH CLUSTERING ORDER BY (name DESC);
