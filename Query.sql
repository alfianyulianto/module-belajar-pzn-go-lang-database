CREATE TABLE customer (
id VARCHAR(100) NOT NULL,
name VARCHAR(100) NOT NULL,
PRIMARY KEY (id)
)ENGINE = INNODB;

SELECT * FROM customer;

DELETE FROM customer;

ALTER TABLE customer
	ADD COLUMN email VARCHAR(100),
	ADD COLUMN balance INT DEFAULT 0,
	ADD COLUMN rating DOUBLE DEFAULT 0.0,
	ADD COLUMN created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	ADD COLUMN birth_date DATE,
	ADD COLUMN married BOOLEAN DEFAULT FALSE;
		
DESCRIBE customer;

INSERT INTO customer (id, name, email, balance, rating, birth_date, married) 
	VALUES ('alfian', 'Alfian', 'alfian@mail.com', 100000, 5.0, '2000-07-01', FALSE);
INSERT INTO customer (id, name, email, balance, rating, birth_date, married) 
	VALUES ('budi', 'Budi', 'budi@mail.com', 100000, 5.0, '1991-01-01', TRUE);
INSERT INTO customer (id, name, email, balance, rating, birth_date, married) 
	VALUES ('indah', 'Indah', 'indah@mail.com', 200000, 4.5, '1998-01-11', false);
	
	
## Insert Data Null
INSERT INTO customer (id, name, email, balance, rating, birth_date, married) 
	VALUES ('joko', 'Joko', NULL, 200000, 4.5, NULL, false);