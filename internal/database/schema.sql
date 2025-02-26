CREATE TABLE stocks (
	ticker TEXT PRIMARY KEY,
	name TEXT NOT NULL,
	market_cap NUMERIC NOT NULL,
	sector TEXT NOT NULL,
	last_updated TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE sector(
	name TEXT PRIMARY KEY,
	sector_market_cap NUMERIC NOT NULL
);
