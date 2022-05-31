-- create post
DROP FUNCTION IF EXISTS create_post;

CREATE FUNCTION create_post(
    "txt" text,
    created timestamp,
    photourl text,
    userid bigint
) RETURNS TABLE(
    id BIGINT,
    "text" text,
    created_at timestamp,
    photo_url text,
    user_email text,
    comments_count BIGINT
) AS $$
DECLARE rid bigint;

BEGIN
INSERT INTO posts(
        "text",
        created_at,
        photo_url,
        user_id
    )
VALUES ($1, $2, $3, $4)
RETURNING posts.id INTO rid;

RETURN query
SELECT p.id AS id,
    p.text AS "text",
    p.created_at AS created_at,
    p.photo_url AS photo_url,
    u.email AS user_email,
    COUNT(c.*) AS comments_count
FROM posts AS p
    INNER JOIN users AS u ON u.id = p.user_id
    LEFT JOIN comments AS c ON c.post_id = p.id
WHERE p.id = rid
GROUP BY p.id,
    u.email;

END $$ LANGUAGE plpgsql;

SELECT create_post('haha', 1652131509, 'photo', 2);

-- delete post
DROP FUNCTION IF EXISTS delete_post;

CREATE FUNCTION delete_post(_id bigint) RETURNS TABLE(
    id BIGINT,
    "text" text,
    created_at timestamp,
    photo_url text,
    user_email text,
    comments_count BIGINT
) AS $$
DECLARE d_uid bigint;

BEGIN
DELETE FROM posts
WHERE posts.id = $1
RETURNING posts.user_id INTO d_uid;

RETURN query
SELECT p.id AS id,
    p.text AS "text",
    p.created_at AS created_at,
    p.photo_url AS photo_url,
    u.email AS user_email,
    COUNT(c.*) AS comments_count
FROM posts AS p
    INNER JOIN users AS u ON u.id = p.user_id
    LEFT JOIN comments AS c ON c.post_id = p.id
WHERE p.user_id = d_uid
GROUP BY p.id,
    u.email;

END $$ LANGUAGE plpgsql;

SELECT delete_post(23);

-- find post by tag
CREATE OR REPLACE FUNCTION find_post(tag text) RETURNS TABLE(
        id BIGINT,
        "text" text,
        created_at timestamp,
        photo_url text,
        user_email text,
        comments_count int
    ) AS $$
SELECT p.id AS id,
    p.text AS "text",
    p.created_at AS created_at,
    p.photo_url AS photo_url,
    u.email AS user_email,
    COUNT(c) AS comments_count
FROM posts AS p
    INNER JOIN users AS u ON u.id = p.user_id
    LEFT JOIN comments AS c ON c.post_id = p.id
    LEFT JOIN post_tags AS pt ON pt.post_id = p.id
WHERE pt.tag = $1
GROUP BY p.id,
    u.email
ORDER BY p.created_at;

$$ LANGUAGE SQL;

-- find post by user_id
CREATE OR REPLACE FUNCTION user_posts(_user_id BIGINT) RETURNS TABLE(
        id BIGINT,
        "text" text,
        created_at timestamp,
        photo_url text,
        user_email text,
        comments_count int
    ) AS $$
SELECT p.id AS id,
    p.text AS "text",
    p.created_at AS created_at,
    p.photo_url AS photo_url,
    u.email AS user_email,
    COUNT(c) AS comments_count
FROM posts AS p
    INNER JOIN users AS u ON u.id = p.user_id
    LEFT JOIN comments AS c ON c.post_id = p.id
WHERE u.id = $1
GROUP BY p.id,
    u.email
ORDER BY p.created_at;

$$ LANGUAGE SQL;

-- add post tags with tags string separated by comma
DROP FUNCTION IF EXISTS add_post_tags;

CREATE FUNCTION add_post_tags(post_id BIGINT, tags_str text) RETURNS TABLE(tag text) AS $$
DECLARE tags text [];

BEGIN tags = string_to_array($2, ',');

INSERT INTO post_tags
VALUES($1, trim(unnest(tags)));

DELETE FROM post_tags t1 USING post_tags t2
WHERE t1.ctid > t2.ctid
    AND t1.post_id = t2.post_id
    AND t1.tag = t2.tag;

RETURN query
SELECT pt.tag AS tag
FROM post_tags AS pt
WHERE pt.post_id = $1;

END $$ LANGUAGE plpgsql;

-- get posts (page_size, page_size * (page_number - 1   ))
DROP FUNCTION IF EXISTS get_posts;

CREATE FUNCTION get_posts(size int, offs int) RETURNS TABLE(
    id BIGINT,
    "text" text,
    created_at timestamp,
    photo_url text,
    user_email text,
    comments_count int
) AS $$
SELECT p.id AS id,
    p.text AS "text",
    p.created_at AS created_at,
    p.photo_url AS photo_url,
    u.email AS user_email,
    COUNT(c) AS comments_count
FROM posts AS p
    INNER JOIN users AS u ON u.id = p.user_id
    LEFT JOIN comments AS c ON c.post_id = p.id
GROUP BY p.id,
    u.email
ORDER BY p.created_at
LIMIT $1 OFFSET $2;

$$ LANGUAGE SQL;

-- Tests
SELECT *
FROM posts;

SELECT get_posts(2, 2);

SELECT add_post_tags(9, 'tag1 , tag2,tag3');

SELECT delete_post(7);

SELECT find_post('tag2');