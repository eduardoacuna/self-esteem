CREATE TYPE notification_rate AS ENUM ('each-day', 'each-week', 'each-two-weeks', 'each-month', 'each-two-months');
CREATE TYPE notification_weekday AS ENUM ('sunday','monday','tuesday','wednesday','thursday','friday','saturday');
CREATE TYPE estimate_outcome AS ENUM ('pending', 'positive', 'negative');

CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  email TEXT UNIQUE NOT NULL,
  notify notification_rate,
  day notification_weekday
);

CREATE TABLE tasks (
  id SERIAL,
  user_id INTEGER REFERENCES users,
  title TEXT NOT NULL,
  estimation INTEGER NOT NULL,
  done BOOLEAN NOT NULL,
  outcome estimate_outcome
);

INSERT INTO users (email, notify, day)
VALUES
('eacuna@nearsoft.com', 'each-week', 'friday'),
('mvalle@nearsoft.com', 'each-week', 'monday'),
('izepeda@nearsoft.com', 'each-two-weeks', 'monday');

INSERT INTO tasks (user_id, title, estimation, done, outcome)
VALUES
((SELECT id FROM users WHERE email='eacuna@nearsoft.com'), 'Add something', 2, true, 'negative'),
((SELECT id FROM users WHERE email='mvalle@nearsoft.com'), 'Refactor something', 4, true, 'positive'),
((SELECT id FROM users WHERE email='izepeda@nearsoft.com'), 'Fix something', 3, true, 'positive'),
((SELECT id FROM users WHERE email='mvalle@nearsoft.com'), 'Remove something', 1, false, 'pending');
