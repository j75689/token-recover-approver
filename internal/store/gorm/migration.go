package gorm

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var v202407161600 = &gormigrate.Migration{
	ID: "202311021600",
	Migrate: func(tx *gorm.DB) error {
		if err := tx.AutoMigrate(&Proof{}); err != nil {
			return err
		}
		if err := tx.AutoMigrate(&ChainState{}); err != nil {
			return err
		}
		if err := tx.AutoMigrate(&TokenRecoverEvent{}); err != nil {
			return err
		}
		return nil
	},
	Rollback: func(tx *gorm.DB) error {
		if err := tx.Migrator().DropTable(&Proof{}); err != nil {
			return err
		}
		if err := tx.Migrator().DropTable(&ChainState{}); err != nil {
			return err
		}
		if err := tx.Migrator().DropTable(&TokenRecoverEvent{}); err != nil {
			return err
		}
		return nil
	},
}

// Version is a migrate version of database
type Version struct {
	ID   int64
	Name string
}

// Migrations is a collection of storage migration patterns
var Migrations = []*gormigrate.Migration{
	v202407161600,
}
