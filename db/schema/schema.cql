CREATE KEYSPACE IF NOT EXISTS users
  WITH replication = {'class': 'NetworkTopologyStategy', 'replication_factor': 3}
  AND durable_writes = true;;

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

CREATE INDEX user_name ON users.info (user_name); # global secondary index

CREATE INDEX auth_token ON users.info ((user_id), auth_token);

CREATE INDEX facebook_access_token on users.info((user_id), facebook_access_token); # local secondary index

CREATE TABLE IF NOT EXISTS users.phrases (
  user_id text,
  phrase_time timestamp,
  word text,
  sentence text,
  correct boolean,
  PRIMARY KEY ((user_id), phrase_time) 
) WITH CLUSTERING ORDER BY (phrase_time DESC);
