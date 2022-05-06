SELECT p.id,
    json_build_object('tags', json_agg(t))
FROM posts AS p
    INNER JOIN post_tags AS pt ON pt.post_id = p.id
    INNER JOIN tags AS t ON PT.tag_id = t.id
GROUP BY p.id;
