\c yaus yaus;

CREATE TABLE IF NOT EXISTS url_maps (
  id SERIAL PRIMARY KEY,
  short_url VARCHAR(7) NOT NULL,
  long_url VARCHAR(255) NOT NULL
)
