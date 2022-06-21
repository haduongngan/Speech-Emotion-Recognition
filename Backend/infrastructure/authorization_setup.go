package infrastructure

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// InitAuthorization init authorization
func InitAuthorization() error {
	var err error
	var adapter *gormadapter.Adapter

	adapter, err = gormadapter.NewAdapterByDB(db)
	if err != nil {
		InfoLog.Println("Not connect into adapter")
		return err
	}

	enforcer, err = casbin.NewEnforcer(keyMatchModel, adapter)
	if err != nil {
		InfoLog.Println("Can't create enforcer casbin")
	}

	enforcer.LoadPolicy()
	// allRole := enforcer.GetAllRoles()
	// log.Println(allRole)
	// allPolicy := enforcer.GetFilteredPolicy(0, "admin")
	// log.Println(allPolicy)
	return nil
}
