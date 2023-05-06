CREATE TABLE IF NOT EXISTS users (
    id            BIGSERIAL PRIMARY KEY NOT NULL,
    vk_id         INT NOT NULL UNIQUE,
    photo_url     TEXT NOT NULL,
    name          VARCHAR(32) NOT NULL,
    city          VARCHAR(32) NOT NULL,
    interested_in VARCHAR(10) CONSTRAINT interested_in_constraint CHECK(interested_in = 'boyfriend' OR interested_in = 'girlfriend')
);

CREATE TABLE IF NOT EXISTS sympathy (
    id             BIGSERIAL PRIMARY KEY NOT NULL,
    first_user_id  BIGINT REFERENCES users (id) NOT NULL,
    second_user_id BIGINT REFERENCES users (id) NOT NULL,
    reciprocity    BOOLEAN NOT NULL
);
