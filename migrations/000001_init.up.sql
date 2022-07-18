CREATE TABLE users
(
    ID            SERIAL       NOT NULL UNIQUE,
    user_name     VARCHAR(255) NOT NULL,
    phone         INT          NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL
);

CREATE TABLE cafes
(
    ID        SERIAL       NOT NULL UNIQUE,
    cafe_name VARCHAR(255) NOT NULL,
    avg_time  INT          NOT NULL,
    avg_check INT          NOT NULL
);

CREATE TABLE tables
(
    ID      SERIAL                                      NOT NULL UNIQUE,
    places  INT                                         NOT NULL,
    cafe_id INT REFERENCES cafes (ID) ON DELETE CASCADE NOT NULL

);

CREATE TABLE bookings
(
    ID         SERIAL                                       NOT NULL UNIQUE,
    start_time TIMESTAMP                                    NOT NULL,
    end_time   TIMESTAMP                                    NOT NULL,
    table_id   INT REFERENCES tables (ID) ON DELETE CASCADE NOT NULL,
    user_id    INT REFERENCES users (ID) ON DELETE CASCADE  NOT NULL
);

