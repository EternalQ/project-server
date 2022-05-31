-- create album
CREATE OR REPLACE FUNCTION create_album(
        _name text,
        _created_at timestamp,
        _user_id BIGINT
    ) RETURNS TABLE (
        id bigint,
        name text,
        created_at timestamp,
        photos_url text
    ) AS $$
DECLARE rid BIGINT;

BEGIN
INSERT INTO album (name, created_at, user_id)
VALUES ($1, $2, $3)
RETURNING album.id INTO rid;

RETURN query
SELECT a.id AS id,
    a.name AS name,
    a.created_at AS created_at,
    array_to_string(array_agg(ap.photo_url), ',') AS photos_url
FROM album AS a
    LEFT JOIN album_photos AS ap ON ap.album_id = a.id
WHERE a.id = rid
GROUP BY a.id;

END $$ LANGUAGE plpgsql;

-- delete by id
CREATE OR REPLACE FUNCTION delete_album(_id bigint) RETURNS TABLE (
        id bigint,
        name text,
        created_at timestamp,
        photos_url text
    ) AS $$
DECLARE d_uid BIGINT;

BEGIN
DELETE FROM album
WHERE album.id = $1
RETURNING album.user_id INTO d_uid;

RETURN query
SELECT a.id AS id,
    a.name AS name,
    a.created_at AS created_at,
    array_to_string(array_agg(ap.photo_url), ',') AS photos_url
FROM album AS a
    LEFT JOIN album_photos AS ap ON ap.album_id = a.id
WHERE a.user_id = d_uid
GROUP BY a.id;

END $$ LANGUAGE plpgsql;

-- find album by id
CREATE OR REPLACE FUNCTION find_album(_id BIGINT) RETURNS TABLE (
        id bigint,
        name text,
        created_at timestamp,
        photos_url text
    ) AS $$
SELECT a.id AS id,
    a.name AS name,
    a.created_at AS created_at,
    array_to_string(array_agg(ap.photo_url), ',') AS photos_url
FROM album AS a
    LEFT JOIN album_photos AS ap ON ap.album_id = a.id
WHERE a.id = $1
GROUP BY a.id;

$$ language SQL;

-- find albums by user_id
CREATE OR REPLACE FUNCTION user_albums(_user_id BIGINT) RETURNS TABLE (
        id bigint,
        name text,
        created_at timestamp,
        photos_url text
    ) AS $$
SELECT a.id AS id,
    a.name AS name,
    a.created_at AS created_at,
    array_to_string(array_agg(ap.photo_url), ',') AS photos_url
FROM album AS a
    INNER JOIN users AS u ON u.id = a.user_id
    LEFT JOIN album_photos AS ap ON ap.album_id = a.id
WHERE u.id = $1
GROUP BY a.id;

$$ language SQL;

-- add photo in album
CREATE OR REPLACE FUNCTION add_photo(_album_id bigint, _photo_url text) RETURNS TABLE (
        id bigint,
        name text,
        created_at timestamp,
        photos_url text
    ) AS $$
INSERT INTO album_photos
VALUES($2, $1);

DELETE FROM album_photos t1 USING album_photos t2
WHERE t1.ctid < t2.ctid
    AND t1.album_id = t2.album_id
    AND t1.photo_url = t2.photo_url;

SELECT a.id AS id,
    a.name AS name,
    a.created_at AS created_at,
    array_to_string(array_agg(ap.photo_url), ',') AS photos_url
FROM album AS a
    LEFT JOIN album_photos AS ap ON ap.album_id = a.id
WHERE a.id = $1
GROUP BY a.id;

$$ language SQL;

-- remove photo from album
CREATE OR REPLACE FUNCTION remove_photo(_album_id bigint, _photo_url text) RETURNS TABLE (
        id bigint,
        name text,
        created_at timestamp,
        photos_url text
    ) AS $$
DELETE FROM album_photos AS ap
WHERE ap.photo_url = $2
    AND ap.album_id = $1;

SELECT a.id AS id,
    a.name AS name,
    a.created_at AS created_at,
    array_to_string(array_agg(ap.photo_url), ',') AS photos_url
FROM album AS a
    LEFT JOIN album_photos AS ap ON ap.album_id = a.id
WHERE a.id = $1
GROUP BY a.id;

$$ language SQL;