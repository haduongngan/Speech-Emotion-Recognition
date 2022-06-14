package infrastructure

import (
	"log"
	"spser/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// OpenConnection Open session using db
func openConnection() (*gorm.DB, error) {
	connectSQL := "host=" + dbHost +
		" user=" + dbUser +
		" dbname=" + dbName +
		" password=" + dbPassword +
		" sslmode=disable"

	db, err := gorm.Open(postgres.Open(connectSQL), &gorm.Config{})
	if err != nil {
		ErrLog.Printf("Problem connecting to database: %+v\n", err)
		return nil, err
	}

	return db, nil
}

func InitDatabase(allowMigrate bool) error {
	var err error
	db, err = openConnection()
	if err != nil {
		return err
	}

	if allowMigrate {
		log.Println("Migrating database...")
		db.AutoMigrate(
			&model.User{},
			&model.Customer{},
			&model.Employee{},
			&model.Call{},
			&model.Segment{},
		)
	}
	return nil
}
