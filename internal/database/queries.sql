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