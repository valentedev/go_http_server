CREATE TABLE users (
   id SERIAL PRIMARY KEY,
   name VARCHAR(50)
);

INSERT INTO users (name)
VALUES ('Bruce Wayne');

INSERT INTO users (name)
VALUES ('Clark Kent');