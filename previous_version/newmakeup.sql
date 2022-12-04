<<<<<<< Updated upstream:newmakeup.sql
/* 
	like I have so much of it I can like do a /\ \/ on individual products, 
	ID
	brands
	product type
	>sub type?
	color / shape / texture
	use (primary use) 
	rating 0-10
	pics > item, swatch, use */

CREATE TABLE cosmetics (
	id INTEGER PRIMARY KEY,
	brand TEXT,
	product_type TEXT,
	color TEXT,
	texture TEXT,
	use TEXT,
	rating INTEGER);

INSERT INTO cosmetics (brand, product_type, color, use, rating) VALUES ("NYX", "lippie", "milan - dusty rose", "matte liquid lip color", 4), ("Maybeline", "lippie", "playful purple", "hydrating color balm", 10);

SELECT * FROM cosmetics;

CREATE TABLE brands (
	id INTEGER PRIMARY KEY,
	brand_name TEXT);

INSERT INTO brands (brand_name) VALUES ("NYX"), ("Maybeline"), ("Milk Makeup"), ("Jeffree Star Cosmetics"), ("Kat Von D"), ("Sugar Pill"), ("Rainbow Honey nail lacquer"), ("wetnwild");

SELECT * FROM brands;

CREATE TABLE collections (
	id INTEGER PRIMARY KEY,
	brand_id INTEGER,
	collection_name TEXT);

INSERT INTO collections (brand_id, collection_name) VALUES (1, "soft matte lip cream"), (2, "baby lips"), (3, "holographic"), (4, "valure liquid lip color"), (5, "pastel goth"), (6, "little twin star"), (7, "mlp"), (8, "gothographic");

INSERT INTO collections (brand_id, collection_name) VALUES (8, "packman");

SELECT * FROM collections;

CREATE TABLE SubCollections (
	id INTEGER PRIMARY KEY,
	collection_id INTEGER,
	subcollection_name TEXT);

INSERt INTO SubCollections (collection_id, subcollection_name) VALUES (2, "color balm crayon");

SELECT * FROM SubCollections;

CREATE TABLE productType (
	id INTEGER PRIMARY KEY,
	type_id INTEGER,
	product_type TEXT);

INSERT INTO productType (type_id, product_type) VALUES (1, "lippie"), (2, "eye"), (3, "highlight"), (4, "blush"), (5, "body"), (6, "face"), (7, "conture"), (8, "glitter");

SELECT * FROM productType;


=======
/* 
	like I have so much of it I can like do a /\ \/ on individual products, 
	ID
	brands
	product type
	>sub type?
	color / shape / texture
	use (primary use) 
	rating 0-10
	pics > item, swatch, use */

CREATE TABLE cosmetics (
	id INTEGER PRIMARY KEY,
	brand TEXT,
	product_type TEXT,
	color TEXT,
	texture TEXT,
	use TEXT,
	rating INTEGER
	tag_id INTEGER);

INSERT INTO cosmetics (brand, product_type, color, use, rating) VALUES ("NYX", "lippie", "milan - dusty rose", "matte liquid lip color", 4), ("Maybeline", "lippie", "playful purple", "hydrating color balm", 10);

SELECT * FROM cosmetics;

CREATE TABLE brands (
	id INTEGER PRIMARY KEY,
	brand_name TEXT);

INSERT INTO brands (brand_name) VALUES ("NYX"), ("Maybeline"), ("Milk Makeup"), ("Jeffree Star Cosmetics"), ("Kat Von D"), ("Sugar Pill"), ("Rainbow Honey nail lacquer"), ("wetnwild");

SELECT * FROM brands;

CREATE TABLE collections (
	id INTEGER PRIMARY KEY,
	brand_id INTEGER,
	collection_name TEXT);

INSERT INTO collections (brand_id, collection_name) VALUES (1, "soft matte lip cream"), (2, "baby lips"), (3, "holographic"), (4, "valure liquid lip color"), (5, "pastel goth"), (6, "little twin star"), (7, "mlp"), (8, "gothographic");

INSERT INTO collections (brand_id, collection_name) VALUES (8, "packman");

SELECT * FROM collections;

CREATE TABLE SubCollections (
	id INTEGER PRIMARY KEY,
	collection_id INTEGER,
	subcollection_name TEXT);

INSERt INTO SubCollections (collection_id, subcollection_name) VALUES (2, "color balm crayon");

SELECT * FROM SubCollections;

CREATE TABLE productType (
	id INTEGER PRIMARY KEY,
	type_id INTEGER,
	product_type TEXT);

INSERT INTO productType (type_id, product_type) VALUES (1, "lippie"), (2, "eye"), (3, "highlight"), (4, "blush"), (5, "body"), (6, "face"), (7, "conture"), (8, "glitter");

SELECT * FROM productType;


>>>>>>> Stashed changes:previous_version/newmakeup.sql
