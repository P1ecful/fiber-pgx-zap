create table song(
                     song_name text not null unique,
                     song_group text not null,
                     release_date date,
                     song_text text,
                     link text
);