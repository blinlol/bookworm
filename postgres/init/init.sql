create table books (
    id uuid default gen_random_uuid() primary key,
    author varchar(64) not null,
    title varchar(64) not null,
    constraint unique unique_author_title_pair (author, title)
);

create table quotes (
    book_id uuid references books(id) not null,
    quotes text[] default array[]::text[]
);
