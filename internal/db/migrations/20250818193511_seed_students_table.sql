-- +goose Up
-- +goose StatementBegin

INSERT INTO students (name, email, phone, address, createdAt)
VALUES ('Emily Chen', 'emily.chen@example.com', '123-456-7890', '123 Main St, Anytown, USA 12345', '2022-01-01 12:00:00');
INSERT INTO students (name, email, phone, address, createdAt)
VALUES    ('Liam Brown', 'liam.brown@example.com', '987-654-3210', '456 Elm St, Othertown, USA 67890', '2022-01-05 14:00:00');
INSERT INTO students (name, email, phone, address, createdAt)
VALUES    ('Ava Patel', 'ava.patel@example.com', '555-123-4567', '789 Oak St, Smalltown, USA 34567', '2022-01-10 10:00:00');
INSERT INTO students (name, email, phone, address, createdAt)
VALUES    ('Noah Lee', 'noah.lee@example.com', '901-234-5678', '321 Pine St, Bigcity, USA 90123', '2022-01-15 16:00:00');
INSERT INTO students (name, email, phone, address, createdAt)
VALUES    ('Sophia Kim', 'sophia.kim@example.com', '111-222-3333', '901 Maple St, Suburbiatown, USA 45678', '2022-01-20 12:00:00');
INSERT INTO students (name, email, phone, address, createdAt)
VALUES    ('Jackson Hall', 'jackson.hall@example.com', '444-555-6666', '234 Cedar St, Universitytown, USA 23456', '2022-01-25 14:00:00');
INSERT INTO students (name, email, phone, address, createdAt)
VALUES    ('Mia Davis', 'mia.davis@example.com', '777-888-9999', '567 Spruce St, Downtown, USA 12345', '2022-02-01 10:00:00');
INSERT INTO students (name, email, phone, address, createdAt)
VALUES    ('Alexander Martin', 'alexander.martin@example.com', '333-444-5555', '890 Walnut St, Uptown, USA 67890', '2022-02-05 16:00:00');
INSERT INTO students (name, email, phone, address, createdAt)
VALUES    ('Isabella Garcia', 'isabella.garcia@example.com', '666-777-8888', '345 Cherry St, Oldtown, USA 34567', '2022-02-10 12:00:00');
INSERT INTO students (name, email, phone, address, createdAt)
VALUES    ('Ethan Harris', 'ethan.harris@example.com', '999-000-1111', '678 Orange St, Newtown, USA 90123', '2022-02-15 14:00:00');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM students;
-- +goose StatementEnd
