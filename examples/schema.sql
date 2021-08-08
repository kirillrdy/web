CREATE TABLE movies (
  id serial,
  title text,
  year int,
  created_at timestamp DEFAULT now()
);
