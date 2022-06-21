package infrastructure

import (
	"log"
	"spser/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func openConnection() (*gorm.DB, error) {
	connectSQL := "host=" + dbHost +
		" user=" + dbUser +
		" dbname=" + dbName +
		" password=" + dbPassword +
		" sslmode=disable"
	db, err := gorm.Open(postgres.Open(connectSQL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
		// DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		ErrLog.Printf("Not connect to database: %+v\n", err)
		return nil, err
	}

	return db, nil
}

func CloseConnection(db *gorm.DB) {
	sqlDB, _ := db.DB()
	sqlDB.Close()
}

// InitDatabase open connection and migrate database
func InitDatabase(allowMigrate bool) error {
	var err error
	db, err = openConnection()
	if err != nil {
		return err
	}

	// err = db.SetupJoinTable(&model.Department{}, "Profiles", &model.DepartmentProfile{})
	// if err != nil {
	// 	return err
	// }

	// err = db.SetupJoinTable(&model.KPIYear{}, "Profiles", &model.KPIYearProfile{})
	// if err != nil {
	// 	return err
	// }

	if allowMigrate {
		log.Println("Migrating database...")

		db.AutoMigrate(
			&model.User{},
			&model.Customer{},
			&model.Staff{},
			&model.Call{},
			&model.Segment{},
		)
		log.Println("Done migrating database")
	}

	// set up extension
	err = db.Exec("CREATE EXTENSION IF NOT EXISTS pg_trgm").Error
	if err != nil {
		return err
	}

	return nil
}
