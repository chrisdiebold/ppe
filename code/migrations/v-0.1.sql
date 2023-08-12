-- SQLite

DROP TABLE todos;

CREATE TABLE IF NOT EXISTS todos(
    id INTEGER PRIMARY KEY,
    name text NOT NULL,
    description string NOT NULL, 
    completed boolean NOT NULL,
    createdOn Date NOT NULL,
    completedOn Date NOT NULL
);
insert into todos (name, description, completed, createdOn, completedOn)
    VALUES ('First todo', 'Do your homework!', false, '11-17-2000', '11-17-2000');

select * from todos;