-- create comment
DROP FUNCTION IF EXISTS create_COMMENT;

CREATE FUNCTION create_comment(
    "comm" text,
    created BIGINT,
    postid BIGINT,
    userid bigint
) RETURNS TABLE (
    id bigint,
    "comment" text,
    created_at timestamp,
    user_email text
) AS $$
DECLARE ret_id bigint;

BEGIN
INSERT INTO comments("comment", created_at, post_id, user_id)
VALUES($1, to_timestamp($2), $3, $4)
RETURNING comments.id INTO ret_id;

RETURN query
SELECT c.id AS id,
    c.comment AS "comment",
    c.created_at AS created_at,
    u.email AS user_email
FROM comments AS c
    INNER JOIN users AS u ON c.user_id = u.id
    INNER JOIN posts AS p ON c.post_id = p.id
WHERE c.id = ret_id;

END $$ LANGUAGE plpgsql;

-- find by post id
CREATE OR REPLACE FUNCTION find_comments(_post_id BIGINT) RETURNS TABLE (
        id bigint,
        "comment" text,
        created_at timestamp,
        user_email text
    ) AS $$ BEGIN RETURN query
SELECT c.id AS id,
    c.comment AS "comment",
    c.created_at AS created_at,
    u.email AS user_email
FROM comments AS c
    INNER JOIN users AS u ON c.user_id = u.id
    INNER JOIN posts AS p ON c.post_id = p.id
WHERE p.id = $1;

END $$ LANGUAGE plpgsql;

SELECT find_comment(9);

-- Test
SELECT *
FROM comments;

SELECT create_comment('hehehe', 1652175676, 7, 2);