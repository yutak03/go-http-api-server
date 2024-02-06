create table IF NOT EXISTS article (
    article_id INTEGER UNSIGNED auto_increment PRIMARY KEY,
    title VARCHAR(100) not null,
    contents text not null,
    username VARCHAR(100) not null,
    nice INTEGER not null,
    created_at datetime
);

create table IF NOT EXISTS comment (
    comment_id INTEGER UNSIGNED auto_increment PRIMARY KEY,
    article_id integer unsigned not null,
    message TEXT not NULL,
    created_at datetime,
    Foreign Key (article_id) REFERENCES article(article_id)
);