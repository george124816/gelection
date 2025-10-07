CREATE TABLE IF NOT EXISTS votes(
	id UUID PRIMARY KEY,
	inserted_at TIMESTAMP,
	election_id BIGSERIAL,
	candidate_id BIGSERIAL
);
