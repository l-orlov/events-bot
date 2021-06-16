CREATE OR REPLACE FUNCTION trigger_set_timestamp()
    RETURNS TRIGGER AS
$$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TABLE IF NOT EXISTS userprofile
(
    userprofile_id        BIGSERIAL PRIMARY KEY,
    location              VARCHAR(255) NOT NULL,
    tags                  TEXT[]       NOT NULL DEFAULT ARRAY []::TEXT[],
    created_at            TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at            TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    is_needed_free_events BOOLEAN      NOT NULL DEFAULT FALSE
);
CREATE TRIGGER set_timestamp
    BEFORE UPDATE
    ON userprofile
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();
