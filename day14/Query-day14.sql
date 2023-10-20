-- nomor1
CREATE DATABASE toko_online_alta;

USE toko_online_alta;

-- nomor2
CREATE TABLE pembeli(
	hp VARCHAR(13) PRIMARY KEY,
	nama VARCHAR(255) NOT NULL,
	alamat VARCHAR(255),
	tanggal_lahir DATE,
	gender SMALLINT CHECK(gender < 3),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP(),
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(),
	`status` BOOL DEFAULT true
);

CREATE TABLE produk(
	produk_id VARCHAR(50) PRIMARY KEY,
	nama_produk VARCHAR(255) NOT NULL,
	jenis_produk VARCHAR(255) NOT NULL,
	deskripsi TEXT,
	merk VARCHAR(50) NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP(),
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(),
	stok INTEGER(11) NOT NULL,
	harga INTEGER(11) NOT NULL
);

CREATE TABLE transaksi (
	nomor_invoice VARCHAR(20) PRIMARY KEY,
	customer VARCHAR(13),
	tanggal_transaksi DATE,
	qty INTEGER(5),
	total_transaksi INTEGER(15),
	FOREIGN KEY (customer) REFERENCES pembeli(hp)
);

CREATE TABLE pembayaran (
	nomor_invoice VARCHAR(20) PRIMARY KEY,
	payment_method VARCHAR(45) NOT NULL,
	status_pembayaran SMALLINT NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP(),
	FOREIGN KEY (nomor_invoice) REFERENCES transaksi(nomor_invoice)
);

CREATE TABLE detail_transaksi (
	nomor_invoice VARCHAR(20),
	produk_id VARCHAR(50),
	qty INTEGER(5) NOT NULL,
	sub_total INTEGER(11) NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP(),
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(),
	FOREIGN KEY (nomor_invoice) REFERENCES transaksi(nomor_invoice),
	FOREIGN KEY (produk_id) REFERENCES produk(produk_id),
	PRIMARY KEY (nomor_invoice,produk_id)
);

-- nomor3
CREATE TABLE kurir (
	hp VARCHAR(13) PRIMARY KEY,
	nama VARCHAR(255) NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP(),
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP()	
);

-- nomor4
ALTER TABLE kurir ADD COLUMN ongkos_dasar INTEGER(11);

-- nomor5
ALTER TABLE kurir RENAME TO shipping;

-- nomor6
DROP TABLE shipping;

-- nomor7
CREATE TABLE payment_method (
	id_method INTEGER(11) PRIMARY KEY,
	nama_method VARCHAR(255) NOT NULL
);
CREATE TABLE deskripsi (
	id_deskripsi INTEGER(11) PRIMARY KEY,
	id_method INTEGER(11),
	deskripsi VARCHAR(255) NOT NULL,
	FOREIGN KEY (id_method) REFERENCES payment_method(id_method)
);

ALTER TABLE pembeli DROP COLUMN alamat;

CREATE TABLE alamat (
	id_alamat VARCHAR(60) PRIMARY KEY,
	alamat VARCHAR(255) NOT NULL,
	hp VARCHAR(13),
	FOREIGN KEY (hp) REFERENCES pembeli(hp)
);


CREATE TABLE user_payment_method_detail (
	hp VARCHAR(13),
	id_method INTEGER(11),
	PRIMARY KEY (hp,id_method),
	FOREIGN KEY (hp) REFERENCES pembeli(hp),
	FOREIGN KEY (id_method) REFERENCES payment_method(id_method)
);