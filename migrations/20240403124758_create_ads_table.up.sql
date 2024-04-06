CREATE SCHEMA IF NOT EXISTS advertising;

CREATE TABLE IF NOT EXISTS advertising.ads (
    aid SERIAL,
    title VARCHAR(128) DEFAULT '',
    start_at TIMESTAMP WITH TIME ZONE DEFAULT NULL,
    end_at TIMESTAMP WITH TIME ZONE DEFAULT NULL,
    age_start  INTEGER DEFAULT 0,
    age_end  INTEGER DEFAULT 0,
    gender VARCHAR(16)[] DEFAULT '{}',
    country VARCHAR(256)[] DEFAULT '{}',
    platform VARCHAR(256)[] DEFAULT '{}',
    updated_at TIMESTAMP DEFAULT current_timestamp,
    CONSTRAINT wallet_a_id_pkey PRIMARY KEY (aid)
);