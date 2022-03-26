CREATE TABLE IF NOT EXISTS users.phrases (
  user_id uuid,
  phrase_time timestamp,
  word text,
  sentence text,
  PRIMARY KEY ((user_id), cur_time) 
) WITH CLUSTERING ORDER BY (cur_time DESC);
