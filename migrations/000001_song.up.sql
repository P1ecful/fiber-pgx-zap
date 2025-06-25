create table Song(
                     id int not null primary key,
                     album_id int not null,
                     author_id int not null,
                     title text not null,
                     release_date date not null,
                     song_text text,
                     song_link text

);