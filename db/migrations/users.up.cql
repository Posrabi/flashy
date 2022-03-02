CREATE KEYSPACE IF NOT EXISTS users
  WITH replication = {'class': 'NetworkTopologyStategy', 'replication_factor': 3}
  AND durable_writes = true;;

CREATE TABLE IF NOT EXISTS users.info (
  user_id uuid,
  user_name text,
  name text,
  email text,
  phone_number text,
  hash_password text,
  auth_token text,
  PRIMARY KEY (user_id)
);

CREATE INDEX user_name ON users.info (user_name);

CREATE INDEX auth_token ON users.info ((user_id), auth_token);
