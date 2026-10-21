-- name: ReadIdea :one
select * from ideas
where id = ? limit 1;

-- name: ReadIdeas :many
select * from ideas
order by title;

-- name: NewIdea :one
insert into ideas (
    title, content
) values (
    ?, ?
)
returning *;



-- name: ReadKeyword :one
select * from keywords
where keyword = ? collate nocase
limit 1;

-- name: ReadKeywords :many
select * from keywords
order by keyword;

-- name: NewKeyword :one
insert into keywords (
                     keyword
) values (
          ?
         )
returning *;