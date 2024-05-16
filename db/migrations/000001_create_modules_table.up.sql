CREATE TABLE modules (
    id INTEGER PRIMARY KEY,
    namespace character varying(256) not null,
    name character varying(256) not null,
    provider character varying(256) not null,
    version character varying(256) not null,
    owner character varying(256) not null,
    location character varying(256) not null,
    definition character varying(256) not null,
    downloads integer not null default 0,
    published_at timestamp with time zone,
    root jsonb,
    submodules jsonb
);
