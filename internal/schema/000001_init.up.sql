CREATE TABLE participant (
                             id SERIAL PRIMARY KEY,
                             first_name VARCHAR(255),
                             last_name VARCHAR(255),
                             email VARCHAR(255) UNIQUE NOT NULL,
                             birth_date DATE,
                             username VARCHAR(255) UNIQUE NOT NULL,
                             password VARCHAR(255) NOT NULL,
                             role VARCHAR(50) CHECK (role IN ('user', 'organizer'))
);

CREATE TABLE event (
                       id SERIAL PRIMARY KEY,
                       organizer_id INT REFERENCES participant(id),
                       name VARCHAR(255) NOT NULL,
                       description TEXT,
                       category VARCHAR(100),
                       date_time TIMESTAMP NOT NULL,
                       address VARCHAR(255),
                       ticket_price NUMERIC(10, 2)
);

CREATE TABLE event_feedback (
                                id SERIAL PRIMARY KEY,
                                event_id INT REFERENCES event(id),
                                sender_id INT REFERENCES participant(id),
                                text TEXT
);

CREATE TABLE user_event_link (
                                 user_id INT REFERENCES participant(id),
                                 event_id INT REFERENCES event(id),
                                 link_type VARCHAR(100) CHECK (link_type IN ('like', 'going'))
                                 PRIMARY KEY (user_id, event_id)
);

CREATE TABLE ticket (
                        id SERIAL PRIMARY KEY,
                        event_id INT REFERENCES event(id),
                        owner_id INT REFERENCES participant(id),
                        purchase_date TIMESTAMP NOT NULL,
                        seat_number VARCHAR(50)
);
