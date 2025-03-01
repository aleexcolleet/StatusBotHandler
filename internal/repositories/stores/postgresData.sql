CREATE TABLE urls (
    url TEXT PRIMARY KEY
);

CREATE TABLE url_status (
    url TEXT PRIMARY KEY,
    status BOOLEAN,
    comment TEXT,
    status_code INT
);
