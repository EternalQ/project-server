-- user creating returns his id
CREATE OR REPLACE FUNCTION create_user(
        email text,
        encrypted_password text,
        created_at timestamp
    ) RETURNS SETOF users AS $$
DECLARE rid bigint;

BEGIN
INSERT INTO users (email, encrypted_password, created_at)
VALUES ($1, $2, $3)
RETURNING users.id INTO rid;

RETURN query
SELECT *
FROM users
WHERE users.id = rid;

END $$ LANGUAGE plpgsql;

select * from users
SELECT create_user('test5', 'test', 1652114047);

-- get all users
DROP FUNCTION IF EXISTS all_users();

CREATE FUNCTION all_users() RETURNS SETOF users AS $$
SELECT *
FROM users;

$$ LANGUAGE SQL;

-- finds user by id
DROP FUNCTION IF EXISTS find_user(bigint);

CREATE FUNCTION find_user(id bigint) RETURNS SETOF users AS $$
SELECT *
FROM users AS u
WHERE u.id = $1;

$$ LANGUAGE SQL;

-- finds user by email
DROP FUNCTION IF EXISTS find_user(text);

CREATE FUNCTION find_user(email text) RETURNS SETOF users AS $$
SELECT *
FROM users AS u
WHERE u.email = $1;

$$ LANGUAGE SQL;

-- Tests
SELECT *
FROM users AS u
WHERE u.id = 2;

SELECT *
FROM users;

SELECT create_user('test3', 'test', 1652114047);

SELECT all_users();

-- all FUNCTIONs
SELECT routine_name
FROM information_schema.routines
WHERE routine_type = 'FUNCTION'
    AND routine_schema = 'public';