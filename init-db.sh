#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL

CREATE TABLE IF NOT EXISTS products (
        id uuid NOT NULL PRIMARY KEY,
        title varchar(64) NOT NULL,
        image varchar(400) NOT NULL,
        price float NOT NULL DEFAULT 0,
        brand varchar(20) NOT NULL,
        review_score float NOT NULL DEFAULT 0, 
        created_at timestamptz NOT NULL DEFAULT now(),
        updated_at timestamptz NOT NULL DEFAULT now(),
        UNIQUE (title, brand) 
);

CREATE TABLE IF NOT EXISTS buyers (
        id uuid NOT NULL PRIMARY KEY,
        name varchar(64) NOT NULL,
        email varchar(150) UNIQUE NOT NULL,
        created_at timestamptz NOT NULL DEFAULT now(),
        updated_at timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS product_to_buyer (
    product_id uuid NOT NULL REFERENCES products(id) ON DELETE CASCADE,
    buyer_id uuid NOT NULL REFERENCES buyers(id) ON DELETE CASCADE,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    CONSTRAINT product_to_buyer_pkey PRIMARY KEY (product_id, buyer_id) 
);

INSERT INTO buyers (id, name, email) VALUES ('925fa490-7fa5-4fee-8035-3b691af02cb8', 'Fulano', 'fulano@gmail.com');
INSERT INTO buyers (id, name, email) VALUES ('e5b3cac0-66d5-4352-9e12-72e695ac7c35', 'Beltrano', 'beltrano@gmail.com');
INSERT INTO buyers (id, name, email) VALUES ('ae382773-e6ed-47eb-8113-dcb7af7d2d95', 'Sicrano', 'sicrano@gmail.com');

INSERT INTO products 
    (id, title, image, price, brand, review_score) 
    VALUES ('f4cacf02-4e71-42ad-9fd8-ace6d27a4c87', 'Galaxy S20', 'http://image/glaxy.jpg', 2000.00, 'Samsung', 4.0);

INSERT INTO products 
    (id, title, image, price, brand, review_score) 
    VALUES ('469fcf5c-dd46-4e53-9ca1-98524a0f3da8', 'Galaxy Note S21', 'http://image/glaxy.jpg', 3500.10, 'Samsung', 4.0);

INSERT INTO products 
    (id, title, image, price, brand, review_score) 
    VALUES ('f11adfcc-6de1-473e-badc-22853e5d2137', 'Smart TV Crystal', 'https://a-static.mlcdn.com.br/618x463/smart-tv-crystal-uhd-4k-led-50-samsung-50tu8000-wi-fi-bluetooth-hdr-3-hdmi-2-usb/magazineluiza/225605700/a24e306a3952101b716be6c2a35be53b.jpg', 1060.99, 'Samsung', 4.0);

INSERT INTO products 
    (id, title, image, price, brand, review_score) 
    VALUES ('8615f5ee-08ba-41af-b5ae-15ee381f64eb', 'Moto X', 'https://a-static.mlcdn.com.br/618x463/smartphone-motorola-moto-g10-64gb-branco-floral-4g-4gb-ram-tela-65-cam-quadrupla-selfie-8mp/magazineluiza/155626200/a5c10227a1af3489bf695b4e8ae67995.jpg', 100.98, 'Motorola', 4.0);

INSERT INTO products 
    (id, title, image, price, brand, review_score) 
    VALUES ('d494605f-e822-4c9f-9a77-80b56630907a', 'MIband 20', 'http://image/miband.jpg', 257.00, 'Xiaomi', 4.0);

INSERT INTO products 
    (id, title, image, price, brand, review_score) 
    VALUES ('3423eea8-3fab-4350-8654-a5f4c1a0c542', 'HyperX Cloud Core 2', 'http://image/miband.jpg', 490.90, 'HyperX', 4.0);

INSERT INTO product_to_buyer (buyer_id, product_id) VALUES ('925fa490-7fa5-4fee-8035-3b691af02cb8', 'f4cacf02-4e71-42ad-9fd8-ace6d27a4c87');
INSERT INTO product_to_buyer (buyer_id, product_id) VALUES ('925fa490-7fa5-4fee-8035-3b691af02cb8', '469fcf5c-dd46-4e53-9ca1-98524a0f3da8');
INSERT INTO product_to_buyer (buyer_id, product_id) VALUES ('e5b3cac0-66d5-4352-9e12-72e695ac7c35', '469fcf5c-dd46-4e53-9ca1-98524a0f3da8');
INSERT INTO product_to_buyer (buyer_id, product_id) VALUES ('e5b3cac0-66d5-4352-9e12-72e695ac7c35', '3423eea8-3fab-4350-8654-a5f4c1a0c542');

EOSQL
