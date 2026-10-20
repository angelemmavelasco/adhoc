-- name: GetIdea :one
select * from ideas
where id = ? limit 1;

-- name: ListIdeas :many
select * from ideas
order by title;

-- name: CreateIdea :one
insert into ideas (
    title, content
) values (
    ?, ?
)