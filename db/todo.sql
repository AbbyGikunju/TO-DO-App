DROP TABLE IF EXISTS users;

CREATE TABLE users (
    id         INT AUTO_INCREMENT NOT NULL,
    username   VARCHAR(50) NOT NULL,
    email      VARCHAR(100) NOT NULL,
    password   VARCHAR(255) NOT NULL,
    PRIMARY KEY(`id`)
);

DROP TABLE IF EXISTS todos;
 
CREATE TABLE todos(
    id           INT AUTO_INCREMENT PRIMARY KEY,
    user_id      INT NOT NULL,
    title        VARCHAR(50) NOT NULL,
    priority     INT NOT NULL,
    date_created DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP, 
    deadline     DATETIME,
    completed    BOOLEAN NOT NULL DEFAULT FALSE,
    description  TEXT,
    UNIQUE(user_id, title),
    FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
);