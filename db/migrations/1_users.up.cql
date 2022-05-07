CREATE KEYSPACE IF NOT EXISTS users
  WITH replication = {'class': 'NetworkTopologyStrategy', 'replication_factor': 3}
  AND durable_writes = true;

CREATE TABLE IF NOT EXISTS users.info (
  user_id text,
  user_name text,
  name text,
  email text,
  hash_password text,
  facebook_access_token text,
  auth_token text,
  PRIMARY KEY ((user_id))
);

CREATE INDEX user_name ON users.info (user_name);

CREATE INDEX auth_token ON users.info ((user_id), auth_token);

CREATE INDEX facebook_access_token on users.info((user_id), facebook_access_token);
