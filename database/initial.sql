/* Initial database script for the boilerplate */
CREATE DATABASE IF NOT EXISTS boilerplate;

CREATE TABLE tab_todo (
    todo_id INTEGER UNSIGNED AUTO_INCREMENT,
    title VARCHAR(50),
    PRIMARY KEY(todo_id)
);