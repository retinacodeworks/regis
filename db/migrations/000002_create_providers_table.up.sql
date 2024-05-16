CREATE TABLE providers (
    id INTEGER PRIMARY KEY,
    namespace character varying(256) not null,
    type character varying(256) not null,
    version character varying(256) not null,
    protocols jsonb not null,
    platforms jsonb not null,
    gpg_public_keys jsonb not null,
    published_at timestamp with time zone
);