CREATE SCHEMA IF NOT EXISTS advertising;

CREATE TABLE IF NOT EXISTS advertising.ads (
    a_id SERIAL,
    title VARCHAR(128) DEFAULT '',
    start_at VARCHAR(64) DEFAULT '',
    end_at VARCHAR(64) DEFAULT '',
    age_start NUMERIC(3) DEFAULT 0,
    age_end NUMERIC(3) DEFAULT 0,
    gender VARCHAR(16) DEFAULT '',
    country VARCHAR(256) DEFAULT '',
    platform VARCHAR(256) DEFAULT '',
    updated_at TIMESTAMP DEFAULT current_timestamp,
    CONSTRAINT wallet_a_id_pkey PRIMARY KEY (a_id)
);