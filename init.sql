CREATE TABLE short_links (
    id TEXT PRIMARY KEY,
    long_url TEXT NOT NULL,
    expiry TIMESTAMP,
    usage_count INTEGER DEFAULT 0
);