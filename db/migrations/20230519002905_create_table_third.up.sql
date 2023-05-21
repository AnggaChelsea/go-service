CREATE TABLE table_third
(
    id INT NOT NULL AUTO_INCREMENT,
    nama_lengkap VARCHAR(200) NOT NULL,
    username VARCHAR(200) NOT NULL,
    email VARCHAR(100) NOT NULL,
    isPerusahaan BOOL NOT NULL,
    password VARCHAR(100) NOT NULL,
    PRIMARY KEY (id)
)ENGINE = InnoDB;