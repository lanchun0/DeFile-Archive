CREATE TABLE `blockchain_simulation`.`asset_table` (
	`id` VARCHAR(256) NULL,
	`fileid` VARCHAR(256) NULL,
	`permission` TINYINT(8) NULL);

CREATE TABLE `blockchain_simulation`.`signature_table` (
  `fileid` varchar(256) DEFAULT NULL,
  `r` varchar(256) DEFAULT NULL,
  `s` varchar(256) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `blockchain_simulation`.`loginfo_table` (
  `behavior_publickey` varchar(255) NOT NULL,
  `timestamp` varchar(45) DEFAULT NULL,
  `publickey` varchar(256) DEFAULT NULL,
  `method` varchar(16) DEFAULT NULL,
  `fileid` varchar(256) DEFAULT NULL,
  `hashdigest` varchar(256) DEFAULT NULL,
  PRIMARY KEY (`behavior_publickey`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


CREATE TABLE `blockchain_simulation`.`behavior_table` (
  `publickey` VARCHAR(256) NOT NULL,
  `privatekey` VARCHAR(256) NULL,
  PRIMARY KEY (`publickey`));
  
  CREATE TABLE `blockchain_simulation`.`permissionlist_table` (
  `id` VARCHAR(256) NULL,
  `publickey` VARCHAR(256) NULL,
  `permission` VARCHAR(45) NULL);

  CREATE TABLE `blockchain_simulation`.`metedata_table` (
  `id` VARCHAR(256) NOT NULL,
  `filename` VARCHAR(45) NULL,
  `size` INT NULL,
  `timestamp` VARCHAR(45) NULL,
  PRIMARY KEY (`id`));

  CREATE TABLE `blockchain_simulation`.`data_table` (
  `id` VARCHAR(256) NOT NULL,
  `hashdigest` VARCHAR(256) NULL,
  `owner` VARCHAR(256) NULL,
  `permissionlevel` VARCHAR(45) NULL,
  PRIMARY KEY (`id`));