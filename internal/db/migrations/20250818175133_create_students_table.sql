-- +goose Up
-- +goose StatementBegin
CREATE TABLE students(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL ,
    phone VARCHAR(20) NOT NULL,
    address TEXT,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE students;
-- +goose StatementEnd
