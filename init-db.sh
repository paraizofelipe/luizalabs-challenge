#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL

CREATE TABLE IF NOT EXISTS products (
        id uuid NOT NULL,
        title varchar(64) NOT NULL,
        image varchar(150) NOT NULL,
        price float NOT NULL DEFAULT 0,
        brand varchar(20) NOT NULL,
        review_score float NOT NULL DEFAULT 0, 
        created_at timestamptz NOT NULL DEFAULT now(),
        updated_at timestamptz NOT NULL DEFAULT now(),
        CONSTRAINT products_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS buyers (
        id uuid NOT NULL,
        name varchar(64) NOT NULL,
        email varchar(150) UNIQUE NOT NULL,
        created_at timestamptz NOT NULL DEFAULT now(),
        updated_at timestamptz NOT NULL DEFAULT now(),
        CONSTRAINT buyers_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS product_to_buyer (
    product_id uuid NOT NULL,
    buyer_id uuid NOT NULL,
    FOREIGN KEY (product_id) REFERENCES buyers(id), 
    FOREIGN KEY (buyer_id) REFERENCES products(id),
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    UNIQUE (product_id, buyer_id)
);

EOSQL
