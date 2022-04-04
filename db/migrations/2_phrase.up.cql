CREATE TABLE IF NOT EXISTS users.phrases (
  user_id uuid,
  phrase_time timestamp,
  word text,
  sentence text,
  correct boolean,
  PRIMARY KEY ((user_id), phrase_time) 
) WITH CLUSTERING ORDER BY (phrase_time DESC);
