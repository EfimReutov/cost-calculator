CREATE TABLE IF NOT EXISTS categories
(
    id       SMALLSERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);

INSERT INTO categories (name)
VALUES ('food');
INSERT INTO categories (name)
VALUES ('car');
INSERT INTO categories (name)
VALUES ('communal');
INSERT INTO categories (name)
VALUES ('clothes');
INSERT INTO categories (name)
VALUES ('other');


CREATE TABLE IF NOT EXISTS spends
(
    id          BIGSERIAL PRIMARY KEY,
    category_id SMALLINT  NOT NULL REFERENCES categories (id),
    amount      DECIMAL   NOT NULL,
    description TEXT,
    date        TIMESTAMP NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS sources
(
    id   SMALLSERIAL PRIMARY KEY,
    name TEXT      NOT NULL
);

INSERT INTO sources (name)
VALUES ('Victoria');
INSERT INTO sources (name)
VALUES ('Efim');

CREATE TABLE IF NOT EXISTS incoming
(
    id        BIGSERIAL NOT NULL,
    source_id SMALLINT  NOT NULL REFERENCES sources (id),
    amount    DECIMAL   NOT NULL,
    date TIMESTAMP NOT NULL DEFAULT (now())

);