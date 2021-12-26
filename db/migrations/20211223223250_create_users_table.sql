-- migrate:up
CREATE TABLE users (
    id SERIAL NOT NULL PRIMARY KEY,
    spotify_username VARCHAR(255) UNIQUE,
    lastfm_username VARCHAR(255),
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT (NOW() AT TIME ZONE 'utc'),
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT (NOW() AT TIME ZONE 'utc')
);

-- migrate:down
DROP TABLE users;
