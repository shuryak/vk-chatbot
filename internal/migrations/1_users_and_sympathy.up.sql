CREATE TABLE IF NOT EXISTS users (
    vk_id         BIGINT PRIMARY KEY NOT NULL UNIQUE,
    photo_url     TEXT,
    name          VARCHAR(32),
    city          VARCHAR(32),
    interested_in VARCHAR(10) CONSTRAINT interested_in_constraint CHECK(interested_in = 'boys' OR interested_in = 'girls')
);

CREATE TABLE IF NOT EXISTS sympathy (
    id             BIGSERIAL PRIMARY KEY NOT NULL,
    first_user_vk_id  BIGINT REFERENCES users (vk_id) NOT NULL,
    second_user_vk_id BIGINT REFERENCES users (vk_id) NOT NULL,
    reciprocity    BOOLEAN NOT NULL
);
