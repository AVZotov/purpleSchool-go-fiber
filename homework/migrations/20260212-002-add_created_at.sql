ALTER TABLE users
    ADD COLUMN created_at TIMESTAMP DEFAULT now();
ALTER TABLE users
    ADD COLUMN updated_at TIMESTAMP DEFAULT now();

CREATE OR REPLACE FUNCTION update_updated_at_column()
    RETURNS TRIGGER AS
$$
BEGIN
    new.updated_at = now();
    RETURN new;
END;
$$ language plpgsql;

CREATE TRIGGER update_users_updated_at
    BEFORE UPDATE
    ON users
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();