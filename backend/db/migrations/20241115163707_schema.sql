-- +goose Up
-- +goose StatementBegin
CREATE TABLE participants(
                             id SERIAL NOT NULL,
                             team_id BIGINT NOT NULL,
                             name VARCHAR(255) NOT NULL,
                             role VARCHAR(255) NOT NULL,
                             birthdate TIMESTAMP(0) WITH
                                 TIME zone NOT NULL,
                             birthplace VARCHAR(255) NOT NULL
);
ALTER TABLE
    participants ADD PRIMARY KEY(id);
CREATE TABLE teams(
                      id SERIAL NOT NULL,
                      user_id BIGINT NOT NULL,
                      name VARCHAR(255) NOT NULL
);
ALTER TABLE
    teams ADD PRIMARY KEY(id);
CREATE TABLE users(
                      id SERIAL NOT NULL,
                      username VARCHAR(128) NOT NULL,
                      password VARCHAR(255) NOT NULL,
                      email VARCHAR(255) NOT NULL
);
ALTER TABLE
    users ADD PRIMARY KEY(id);
ALTER TABLE
    users ADD CONSTRAINT users_username_unique UNIQUE(username);
ALTER TABLE
    teams ADD CONSTRAINT teams_user_id_foreign FOREIGN KEY(user_id) REFERENCES users(id);
ALTER TABLE
    participants ADD CONSTRAINT participants_team_id_foreign FOREIGN KEY(team_id) REFERENCES teams(id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
