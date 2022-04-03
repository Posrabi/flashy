ALTER TABLE users.info ADD facebook_access_token text ;

CREATE INDEX facebook_access_token on users.info((user_id), facebook_access_token);
