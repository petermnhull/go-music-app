-- migrate:up
CREATE OR REPLACE FUNCTION update_updated_at_column()   
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW() AT TIME ZONE 'utc';
    RETURN NEW;   
END;
$$ language 'plpgsql';


-- migrate:down
DROP FUNCTION IF EXISTS update_updated_at_column();
