CREATE USER docker;
-- CREATE DATABASE  docker;
-- GRANT ALL PRIVILEGES ON DATABASE docker to docker;
CREATE DATABASE tasksellerx;
GRANT ALL PRIVILEGES ON DATABASE tasksellerx to docker;

CREATE TABLE IF NOT EXISTS Users (
    id serial NOT NULL PRIMARY KEY,
    username varchar NOT NULL UNIQUE,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS Chat (
    id serial NOT NULL PRIMARY KEY,
    chat_name varchar NOT NULL,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS User_to_chat (
    user_id int REFERENCES Users (id),
    chat_id int REFERENCES Chat (id)
);

CREATE TABLE IF NOT EXISTS Messages (
    id serial NOT NULL PRIMARY KEY,
    chat_id int REFERENCES Chat (id),
    author_id int REFERENCES Users (id),
    message_text text NOT NULL,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP
);