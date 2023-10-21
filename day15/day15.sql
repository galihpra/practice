USE tugas_day15;

-- nomor 1 A
INSERT INTO product_types (id_product_types,nama_product_types,created_at, updated_at)
VALUES
	(null,'Minuman',null,null),
	(null,'Makanan Ringan',null,null),
	(null,'Deterjen',null,null);

-- nomor 1 B
INSERT INTO products (id_product,product_type_id, operator_id,`code`,`name`,`status`,price,created_at, updated_at)
VALUES
	(NULL,'1',NULL,'prd123','Air Mineral', 1, 3500, NULL, NULL),
	(NULL,'1',NULL,'prd234','Soda', 2, 5500, NULL, NULL);

-- nomor 1 C
INSERT INTO products (id_product,product_type_id, operator_id,`code`,`name`,`status`,price,created_at, updated_at)
VALUES
	(NULL,'2',NULL,'prd212','Pop Corn', 2, 8700, NULL, NULL),
	(NULL,'2',NULL,'prd443','Keripik Kentang', 1, 9200, NULL, NULL),
	(NULL,'2',NULL,'prd514','Cokelat', 2, 7000, NULL, NULL);
	
-- nomor 1 D
INSERT INTO products (id_product,product_type_id, operator_id,`code`,`name`,`status`,price,created_at, updated_at)
VALUES
	(NULL,'3',NULL,'prd612','Deterjen Bubuk', 2, 9000, NULL, NULL),
	(NULL,'3',NULL,'prd532','Deterjen Cair', 1, 11100, NULL, NULL),
	(NULL,'3',NULL,'prd114','Sabun Colek', 1, 6400, NULL, NULL);
	
-- nomor 1 E
INSERT INTO product_description (id_product_description, `description`, created_at,updated_at)
VALUES
	(NULL, "Air mineral botol 600ml", NULL,NULL),
	(NULL, "Minuman berkarbonasi kemasan kaleng 400ml", NULL,NULL),
	(NULL, "Pop corn rasa original 65gram", NULL,NULL),
	(NULL, "Keripik kentang rasa rumput laut 89gram", NULL,NULL),
	(NULL, "Cokelat batang 40gram", NULL,NULL),
	(NULL, "Deterjen bubuk kemasan 150gram", NULL,NULL),
	(NULL, "Deterjen cair kemasan 800ml", NULL,NULL),
	(NULL, "Sabun colek 100gram gratis piring cantik", NULL,NULL);
	
-- nomor 1 F
INSERT INTO payment_methods (id_payment_methods, nama_payment_methods, status_payment, created_at, updated_at)
VALUES
	(NULL, 'Cash', 1, NULL, NULL),
	(NULL, 'Kredit', 2, NULL, NULL),
	(NULL, 'Transfer Bank', 3, NULL, NULL);
	
-- nomor 1 G
INSERT INTO users (id, nama, `status`, dob, gender, created_at, updated_at)
VALUES
	(NULL, 'Ahmad', 1, '2023-10-21','M',NULL,NULL),
	(NULL, 'Jeni', 2, '2023-10-20','F',NULL,NULL),
	(NULL, 'Anton', 1, '2023-10-21','M',NULL,NULL),
	(NULL, 'Lita', 1, '2023-10-19','F',NULL,NULL),
	(NULL, 'Maya', 2, '2023-10-21','F',NULL,NULL);
	
-- nomor 1 H
INSERT INTO transactions (id_transaction, user_id, payment_method_id, `status`, total_qty, total_price,created_at,updated_at)
VALUES
	(NULL, 1, 1, 'Lunas', 3, 23400,NULL,NULL),
	(NULL, 1, 1, 'Lunas', 2, 7000,NULL,NULL),
	(NULL, 1, 3, 'Lunas', 1, 11100,NULL,NULL),
	(NULL, 2, 3, 'Lunas', 4, 23000,NULL,NULL),
	(NULL, 2, 2, 'Kredit', 2, 12800,NULL,NULL),
	(NULL, 2, 2, 'Kredit', 1, 9000,NULL,NULL),
	(NULL, 3, 1, 'Lunas', 8, 42000,NULL,NULL),
	(NULL, 3, 1, 'Lunas', 5, 55500,NULL,NULL),
	(NULL, 3, 2, 'Kredit', 3, 23000,NULL,NULL),
	(NULL, 4, 3, 'Kredit', 2, 9000,NULL,NULL),
	(NULL, 4, 3, 'Lunas', 3, 15400,NULL,NULL),
	(NULL, 4, 3, 'Lunas', 1, 9200,NULL,NULL);
	
-- nomor 1 I
INSERT INTO transaction_details (transaction_id,product_id,`status`,qty,price,created_at,updated_at)
VALUES
	(1, 2,'Lunas',1,5500,NULL,NULL),
	(1, 4,'Lunas',1,9200,NULL,NULL),
	(1, 3,'Lunas',1,8700,NULL,NULL)
	;
INSERT INTO transaction_details (transaction_id,product_id,`status`,qty,price,created_at,updated_at)
VALUES
	(2, 1,'Lunas',1,3500,NULL,NULL),
	(2, 2,'Lunas',1,5500,NULL,NULL),
	(2, 5,'Lunas',1,7000,NULL,NULL)
	;
INSERT INTO transaction_details (transaction_id,product_id,`status`,qty,price,created_at,updated_at)
VALUES
	(3, 4,'Lunas',2,9200,NULL,NULL),
	(3, 5,'Lunas',1,7000,NULL,NULL),
	(3, 8,'Lunas',3,6400,NULL,NULL)
	;
INSERT INTO transaction_details (transaction_id,product_id,`status`,qty,price,created_at,updated_at)
VALUES
	(4, 3,'Lunas',2,8700,NULL,NULL),
	(4, 5,'Lunas',4,7000,NULL,NULL),
	(4, 6,'Lunas',1,9000,NULL,NULL)
	;
INSERT INTO transaction_details (transaction_id,product_id,`status`,qty,price,created_at,updated_at)
VALUES
	(5, 1,'Kredit',1,3500,NULL,NULL),
	(5, 4,'Kredit',2,9200,NULL,NULL),
	(5, 8,'Kredit',1,6400,NULL,NULL)
	;
INSERT INTO transaction_details (transaction_id,product_id,`status`,qty,price,created_at,updated_at)
VALUES
	(6, 2,'Kredit',2,5500,NULL,NULL),
	(6, 3,'Kredit',3,8700,NULL,NULL),
	(6, 7,'Kredit',1,11100,NULL,NULL)
	;
INSERT INTO transaction_details (transaction_id,product_id,`status`,qty,price,created_at,updated_at)
VALUES
	(7, 1,'Lunas',3,3500,NULL,NULL),
	(7, 2,'Lunas',2,5500,NULL,NULL),
	(7, 5,'Lunas',1,7000,NULL,NULL)
	;
INSERT INTO transaction_details (transaction_id,product_id,`status`,qty,price,created_at,updated_at)
VALUES
	(8, 1,'Lunas',1,3500,NULL,NULL),
	(8, 3,'Lunas',4,8700,NULL,NULL),
	(8, 6,'Lunas',1,9000,NULL,NULL)
	;
INSERT INTO transaction_details (transaction_id,product_id,`status`,qty,price,created_at,updated_at)
VALUES
	(9, 2,'Kredit',1,5500,NULL,NULL),
	(9, 5,'Kredit',1,7000,NULL,NULL),
	(9, 8,'Kredit',2,6400,NULL,NULL)
	;
INSERT INTO transaction_details (transaction_id,product_id,`status`,qty,price,created_at,updated_at)
VALUES
	(10, 1,'Kredit',3,3500,NULL,NULL),
	(10, 3,'Kredit',2,8700,NULL,NULL),
	(10, 7,'Kredit',2,11100,NULL,NULL)
	;
INSERT INTO transaction_details (transaction_id,product_id,`status`,qty,price,created_at,updated_at)
VALUES
	(11, 1,'Lunas',1,3500,NULL,NULL),
	(11, 3,'Lunas',1,8700,NULL,NULL),
	(11, 5,'Lunas',2,7000,NULL,NULL)
	;
INSERT INTO transaction_details (transaction_id,product_id,`status`,qty,price,created_at,updated_at)
VALUES
	(12, 2,'Lunas',1,5500,NULL,NULL),
	(12, 3,'Lunas',4,8700,NULL,NULL),
	(12, 4,'Lunas',3,9200,NULL,NULL)
	;
	
-- nomor 2 A
SELECT nama
FROM users
WHERE users.gender = 'M';

-- nomor 2 B
SELECT *
FROM products
WHERE products.id_product = 3;

-- nomor 2 C
SELECT *
FROM users
WHERE users.nama LIKE '%a%' 
AND users.created_at BETWEEN '2023-10-14' AND '2023-10-22';

-- nomor 2 D
SELECT COUNT(*) FROM users
WHERE gender = 'F';

-- nomor 2 E
SELECT *
FROM users
ORDER BY nama ASC;

-- nomor 2 F
SELECT *
FROM transaction_details
WHERE product_id = 3;

-- nomor 3 A
UPDATE products
SET `name` = 'product dummy'
WHERE id_product = 1;

-- nomor 3 B
UPDATE transaction_details
SET qty = 3
WHERE product_id = 1;

-- nomor 4 A
DELETE FROM products
WHERE id_product = 1;

-- nomor 4 B
DELETE FROM products
WHERE product_type_id = 1;