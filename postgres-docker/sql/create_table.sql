-- Creation of album table
CREATE TABLE IF NOT EXISTS album (
  id INT NOT NULL,
  title varchar(250),
  artist_id INT NOT NULL,
  price NUMERIC(10),
  PRIMARY KEY (id)
);
