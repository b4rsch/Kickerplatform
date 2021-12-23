GRANT ALL PRIVILEGES on DATABASE kickerplatformdb TO admin;
CREATE TABLE users (
                       id serial PRIMARY KEY,
                       username VARCHAR(255) UNIQUE NOT NULL,
                       locationId int references location(id) NOT NULL
);

CREATE TABLE location(
                         id serial PRIMARY KEY,
                         name VARCHAR(255) UNIQUE NOT NULL
);
DROP TABLE matches CASCADE;

CREATE TABLE matches(
                        id serial PRIMARY KEY,
                        date DATE,
                        locationId int references location(id) NOT NULL,
                        best_of int NOT NULL,
                        team_1_player_1 VARCHAR(255) NOT NULL,
                        team_1_player_2 VARCHAR(255) NOT NULL,
                        team_2_player_1 VARCHAR(255) NOT NULL,
                        team_2_player_2 VARCHAR(255) NOT NULL
);

CREATE TABLE game(
                     id serial PRIMARY KEY,
                     match_id int references matches(id) NOT NULL,
                     points_team_1 int NOT NULL,
                     points_team_2 int NOT NULL,
                     team_1_attacker int references users(id) NOT NULL,
                     team_1_defender int references users(id) NOT NULL,
                     team_2_attacker int references users(id) NOT NULL,
                     team_2_defender int references users(id) NOT NULL
);

INSERT INTO location(name) VALUES ('Berlin'), ('München'), ('Würzburg');
