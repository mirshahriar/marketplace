package queries

import "gorm.io/gorm"

func GetMigrationQueryV010() func(tx *gorm.DB) error {
	return func(tx *gorm.DB) error {
		cErr := tx.Exec(`
				CREATE TABLE products (
				  id bigint unsigned NOT NULL AUTO_INCREMENT,
				  name varchar(100) NOT NULL,
				  description varchar(200) NOT NULL,
				  price double NOT NULL,
				  created_at datetime(3) DEFAULT NULL,
				  updated_at datetime(3) DEFAULT NULL,
				  deleted_at datetime(3) DEFAULT NULL,
				  PRIMARY KEY (id),
				  KEY idx_products_deleted_at (deleted_at)
				) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb3;
			`).Error

		if cErr != nil {
			return cErr
		}

		cErr = tx.Exec(`
			CREATE TABLE users (
			  id bigint unsigned NOT NULL AUTO_INCREMENT,
			  username varchar(191) NOT NULL,
			  password longtext NOT NULL,
			  email varchar(191) DEFAULT NULL,
			  PRIMARY KEY (id),
			  UNIQUE KEY username (username),
			  UNIQUE KEY email (email)
			) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb3;
		`).Error

		if cErr != nil {
			return cErr
		}

		cErr = tx.Exec(`
			CREATE TABLE tokens (
			  user_id bigint unsigned NOT NULL,
			  token longtext NOT NULL,
			  KEY fk_tokens_user (user_id),
			  CONSTRAINT fk_tokens_user FOREIGN KEY (user_id) REFERENCES users (id)
			) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb3;
		`).Error

		if cErr != nil {
			return cErr
		}

		return nil
	}
}
