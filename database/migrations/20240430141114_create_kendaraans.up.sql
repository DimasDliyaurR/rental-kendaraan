CREATE TABLE kendaraans (
  id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  nama_kendaraan varchar(20) NOT NULL,
  plat varchar(20) NOT NULL,
  kapasitas integer(10) NOT NULL,
  created_at datetime(3) NOT NULL,
  updated_at datetime(3) NOT NULL,
  PRIMARY KEY (id),
  KEY idx_kendaraans_created_at (created_at),
  KEY idx_kendaraans_updated_at (updated_at)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
