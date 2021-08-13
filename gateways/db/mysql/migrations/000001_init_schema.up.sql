CREATE TABLE IF NOT EXISTS `account` (
  `id` varchar(100) NOT NULL,
  `nome` varchar(45) DEFAULT NULL,
  `cpf` varchar(45) DEFAULT NULL,
  `secret` varchar(100) DEFAULT NULL,
  `balance` float DEFAULT NULL,
  `created_at` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE IF NOT EXISTS `transfer` (
  `id` varchar(45) NOT NULL,
  `account_origin_id` varchar(45) DEFAULT NULL,
  `account_destination_id` varchar(45) DEFAULT NULL,
  `amount` varchar(45) DEFAULT NULL,
  `created_at` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;