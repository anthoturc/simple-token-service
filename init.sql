CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE api_tokens(
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  token_hash VARCHAR(64) UNIQUE NOT NULL,
  is_enabled BOOLEAN NOT NULL DEFAULT TRUE,
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  expires_at TIMESTAMPTZ NOT NULL
);

CREATE OR REPLACE FUNCTION set_expiry_date()
RETURNS TRIGGER AS $$
BEGIN
  NEW.expires_at := NEW.created_at + INTERVAL '2 weeks';
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER before_insert_set_expiration
BEFORE INSERT ON api_tokens
FOR EACH ROW
EXECUTE FUNCTION set_expiry_date();