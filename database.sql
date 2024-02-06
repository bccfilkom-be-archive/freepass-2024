CREATE DATABASE backend_bcc;

CREATE TABLE user (
    id INT PRIMARY KEY AUTO_INCREMENT,
    nim BIGINT,
    name TEXT,
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    major TEXT,
    faculty TEXT,
    status TEXT NOT NULL,
    description TEXT
);

CREATE TABLE post (
    id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE
);

CREATE TABLE comment (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    post_id INT NOT NULL,
    content TEXT NOT NULL,
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE,
    FOREIGN KEY (post_id) REFERENCES post(id) ON DELETE CASCADE
);
