DROP INDEX IF EXISTS users.user_name;
DROP INDEX IF EXISTS users.auth_token;
DROP INDEX IF EXISTS users.facebook_access_token;
DROP TABLE IF EXISTS users.info;
DROP KEYSPACE IF EXISTS users;
